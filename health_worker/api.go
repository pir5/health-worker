package health_worker

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os" // Import this package
	"os/signal"
	"time"

	stdLog "log"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/facebookgo/pidfile"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pir5/health-worker/docs"
	"github.com/pir5/pir5-go/dnsapi/operations"
)

const (
	AllowDomainsKey = "AllowDomains"
)

var CmdAPI = &Command{
	Run:       runAPI,
	UsageLine: "api",
	Short:     "Start API Server",
	Long: `
Start API Server
	`,
}
var globalConfig *Config

// runAPI executes sub command and return exit code.
func runAPI(cmdFlags *GlobalFlags, args []string) error {
	conf, err := initCommand(cmdFlags)
	if err != nil {
		return err
	}
	globalConfig = conf
	logger := log.New("health-status")
	defer func() {
		if err := os.Remove(pidfile.GetPidfilePath()); err != nil {
			log.Fatalf("Error removing %s: %s", pidfile.GetPidfilePath(), err)
		}
	}()

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		conf.DB.UserName,
		conf.DB.Password,
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.DBName,
	))
	if err != nil {
		return err
	}
	e := echo.New()
	e.Logger = logger
	e.StdLogger = stdLog.New(e.Logger.Output(), e.Logger.Prefix()+": ", 0)

	e.GET("/status", status)

	if err := pidfile.Write(); err != nil {
		return err
	}

	defer func() {
		if err := os.Remove(pidfile.GetPidfilePath()); err != nil {
			e.Logger.Fatalf("Error removing %s: %s", pidfile.GetPidfilePath(), err)
		}
	}()

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}}` + "\n",
		Output: logger.Output(),
	}))

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Validator: func(key string, c echo.Context) (bool, error) {
			if globalConfig.IsTokenAuth() {
				for _, v := range globalConfig.Auth.Tokens {
					if key == v {
						return true, nil
					}
				}
			}
			return true, nil

		},
		Skipper: func(c echo.Context) bool {
			return !globalConfig.IsTokenAuth()
		},
	}))

	go func() {
		if err := e.Start(globalConfig.Listen); err != nil {
			e.Logger.Fatalf("shutting down the server: %s", err)
		}
	}()

	transport := httptransport.New(fmt.Sprintf("%s:%d", conf.PdnsAPI.Host, conf.PdnsAPI.Port), "v1", nil)
	pdnsAPI := operations.New(transport, strfmt.Default)

	v1 := e.Group("/v1")
	HealthCheckEndpoints(v1, db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false))
	RoutingPolicyEndpoints(v1, db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false), pdnsAPI)
	VironEndpoints(v1)

	docs.SwaggerInfo.Host = globalConfig.Listen

	if globalConfig.Endpoint != "" {
		u, err := url.Parse(globalConfig.Endpoint)
		if err != nil {
			return err
		}
		docs.SwaggerInfo.Schemes = []string{u.Scheme}
		docs.SwaggerInfo.Host = u.Host
		docs.SwaggerInfo.BasePath = u.Path
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello! HealthWorker!!1")
	})
	v1.GET("/swagger/*", echoSwagger.WrapHandler)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

func status(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

type HTTPError struct {
	echo.HTTPError
}
