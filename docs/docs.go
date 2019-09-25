// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-09-25 19:25:49.120499 +0900 JST m=+0.064564557

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "This is PIR5 HealthCheck worker and API",
        "title": "health-worker",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "127.0.0.1:8080",
    "basePath": "/v1",
    "paths": {
        "/healthchecks": {
            "get": {
                "description": "get healthchecks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get healthchecks",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "HealthCheck ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.HealthCheck"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "create healthCheck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create healthCheck",
                "parameters": [
                    {
                        "description": "HealthCheck Object",
                        "name": "healthCheck",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HealthCheck"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    }
                }
            }
        },
        "/healthchecks/{id}": {
            "put": {
                "description": "update healthCheck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "update healthCheck",
                "parameters": [
                    {
                        "type": "string",
                        "description": "HealthCheck ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "HealthCheck Object",
                        "name": "healthCheck",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HealthCheck"
                        }
                    }
                ],
                "responses": {
                    "200": {},
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete healthCheck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete healthCheck",
                "parameters": [
                    {
                        "type": "string",
                        "description": "HealthCheck ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    }
                }
            }
        },
        "/routingpolicies": {
            "get": {
                "description": "get routingpolicies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get routingpolicies",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "RoutingPolicy ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.RoutingPolicy"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "create healthCheck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create healthCheck",
                "parameters": [
                    {
                        "description": "RoutingPolicy Object",
                        "name": "healthCheck",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.RoutingPolicy"
                        }
                    }
                ],
                "responses": {
                    "201": {},
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    }
                }
            }
        },
        "/routingpolicies/{id}": {
            "put": {
                "description": "update healthCheck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "update healthCheck",
                "parameters": [
                    {
                        "type": "string",
                        "description": "RoutingPolicy ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "RoutingPolicy Object",
                        "name": "healthCheck",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.RoutingPolicy"
                        }
                    }
                ],
                "responses": {
                    "200": {},
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete healthCheck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete healthCheck",
                "parameters": [
                    {
                        "type": "string",
                        "description": "RoutingPolicy ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/health_worker.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "health_worker.HTTPError": {
            "type": "object"
        },
        "model.HealthCheck": {
            "type": "object",
            "properties": {
                "checkInterval": {
                    "type": "integer"
                },
                "db": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "params": {
                    "type": "object",
                    "$ref": "#/definitions/model.HealthCheckParams"
                },
                "routingPolicies": {
                    "type": "object",
                    "$ref": "#/definitions/model.RoutingPolicies"
                },
                "threshould": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "model.HealthCheckParams": {
            "type": "object",
            "properties": {
                "addr": {
                    "type": "string"
                },
                "hostName": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "protocol": {
                    "type": "integer"
                },
                "searchWord": {
                    "type": "string"
                },
                "timeout": {
                    "type": "string"
                }
            }
        },
        "model.RoutingPolicies": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "db": {
                        "type": "string"
                    },
                    "healthCheckID": {
                        "type": "integer"
                    },
                    "id": {
                        "type": "integer"
                    },
                    "recordID": {
                        "type": "integer"
                    },
                    "type": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.RoutingPolicy": {
            "type": "object",
            "properties": {
                "db": {
                    "type": "string"
                },
                "healthCheckID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "recordID": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{ Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface {}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
