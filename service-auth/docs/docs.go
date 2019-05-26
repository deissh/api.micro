// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-05-26 15:38:44.176774549 +0300 MSK m=+0.026998621

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Auth, create tokens, and refresh old",
        "title": "Service Auth API",
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
        "/token.check": {
            "get": {
                "description": "Check access_token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Deactivate old token and create new",
                "operationId": "refresh-token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "service version",
                        "name": "v",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "access_token",
                        "name": "access_token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.CheckResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.ResponseData"
                        }
                    }
                }
            }
        },
        "/token.create": {
            "get": {
                "description": "Generate new access_token and refresh_token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create new token",
                "operationId": "create-token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "service version",
                        "name": "v",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "user email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "permissions, to check on authorization and request if necessary (Example: email,notif)",
                        "name": "scope",
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
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.ResponseData"
                        }
                    }
                }
            }
        },
        "/token.refresh": {
            "get": {
                "description": "Generate new access_token and refresh_token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Deactivate old token and create new",
                "operationId": "refresh-token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "service version",
                        "name": "v",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "refresh_token",
                        "name": "refresh_token",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.RefreshResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.ResponseData"
                        }
                    }
                }
            }
        },
        "/token.remove": {
            "get": {
                "description": "Remove access_token or revome all tokens from this user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Remove token or all",
                "operationId": "remove-token",
                "parameters": [
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
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "remove all tokens",
                        "name": "all",
                        "in": "query"
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
        "handlers.CheckResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "object",
                    "$ref": "#/definitions/handlers.shortToken"
                },
                "v": {
                    "description": "API version",
                    "type": "string"
                }
            }
        },
        "handlers.CreateResponse": {
            "type": "object",
            "properties": {
                "token": {
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
        "handlers.RefreshResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
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
        "handlers.health": {
            "type": "object",
            "properties": {
                "alive": {
                    "type": "boolean"
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
        },
        "handlers.shortToken": {
            "type": "object",
            "properties": {
                "permissions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "role": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
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
