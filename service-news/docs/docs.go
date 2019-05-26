// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-05-26 15:43:31.476337498 +0300 MSK m=+0.024890964

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Users methods",
        "title": "Service Users API",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {},
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/",
    "paths": {
        "/_/health": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show health service",
                "operationId": "get-service-health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.HealthResponse"
                        }
                    }
                }
            }
        },
        "/_/ping": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Ping service",
                "operationId": "ping-service",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.PingResponse"
                        }
                    }
                }
            }
        },
        "/news.create": {
            "post": {
                "description": "Create news and return it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create news",
                "operationId": "create-news",
                "parameters": [
                    {
                        "description": "news body",
                        "name": "news",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.news"
                        }
                    },
                    {
                        "type": "string",
                        "description": "service version",
                        "name": "v",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "user access_token",
                        "name": "access_token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.CreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.ResponseData"
                        }
                    }
                }
            }
        },
        "/news.get": {
            "get": {
                "description": "Return info about news by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Return news by id",
                "operationId": "get-news",
                "parameters": [
                    {
                        "type": "string",
                        "description": "service version",
                        "name": "v",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "news id",
                        "name": "news",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.GetResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.ResponseData"
                        }
                    }
                }
            }
        },
        "/news.remove": {
            "get": {
                "description": "Remove news by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Remove news by id",
                "operationId": "remove-news",
                "parameters": [
                    {
                        "type": "string",
                        "description": "service version",
                        "name": "v",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "news id",
                        "name": "news_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.RemoveResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.ResponseData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.CreateResponse": {
            "type": "object",
            "properties": {
                "news": {
                    "type": "string"
                },
                "v": {
                    "description": "API version",
                    "type": "string"
                }
            }
        },
        "handlers.GetResponse": {
            "type": "object",
            "properties": {
                "news": {
                    "type": "string"
                },
                "v": {
                    "description": "API version",
                    "type": "string"
                }
            }
        },
        "handlers.HealthResponse": {
            "type": "object",
            "properties": {
                "Health": {
                    "type": "object",
                    "$ref": "#/definitions/handlers.health"
                },
                "v": {
                    "description": "API version",
                    "type": "string"
                }
            }
        },
        "handlers.PingResponse": {
            "type": "object",
            "properties": {
                "ping": {
                    "type": "object",
                    "$ref": "#/definitions/handlers.ping"
                },
                "v": {
                    "description": "API version",
                    "type": "string"
                }
            }
        },
        "handlers.RemoveResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "v": {
                    "description": "API version",
                    "type": "string"
                }
            }
        },
        "handlers.ResponseData": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "handlers.UpdateRequest": {
            "type": "object",
            "required": [
                "accessToken"
            ],
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "news": {
                    "type": "object",
                    "$ref": "#/definitions/handlers.news"
                },
                "version": {
                    "description": "API version",
                    "type": "string"
                }
            }
        },
        "handlers.health": {
            "type": "object",
            "properties": {
                "alive": {
                    "type": "boolean"
                }
            }
        },
        "handlers.news": {
            "type": "object",
            "required": [
                "annotation",
                "body",
                "preview",
                "title"
            ],
            "properties": {
                "annotation": {
                    "type": "string"
                },
                "background": {
                    "type": "string"
                },
                "body": {
                    "type": "string"
                },
                "preview": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "types": {
                    "type": "string"
                }
            }
        },
        "handlers.ping": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "service": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
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
