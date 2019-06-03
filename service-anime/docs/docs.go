// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-06-03 20:19:57.219574866 +0300 MSK m=+0.028497895

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Anims methods",
        "title": "Service Anime API",
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
        "/anime.create": {
            "post": {
                "description": "Create anime and return it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create anime",
                "operationId": "create-anime",
                "parameters": [
                    {
                        "description": "anime body",
                        "name": "anime",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handlers.anime"
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
        "/anime.get": {
            "get": {
                "description": "Return info about anime by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Return anime by id",
                "operationId": "get-anime",
                "parameters": [
                    {
                        "type": "string",
                        "description": "service version",
                        "name": "v",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "anime id",
                        "name": "anime",
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
        "/anime.remove": {
            "get": {
                "description": "Remove anime by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Remove anime by id",
                "operationId": "remove-anime",
                "parameters": [
                    {
                        "type": "string",
                        "description": "service version",
                        "name": "v",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "anime id",
                        "name": "anime_id",
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
                "anime": {
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
                "anime": {
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
            "properties": {
                "anime": {
                    "type": "object",
                    "$ref": "#/definitions/handlers.anime"
                }
            }
        },
        "handlers.anime": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "string"
                },
                "annotation": {
                    "type": "string"
                },
                "countries": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "directors": {
                    "type": "string"
                },
                "genres": {
                    "type": "string"
                },
                "kinopoisk_id": {
                    "type": "string"
                },
                "posters": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "studios": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "title_en": {
                    "type": "string"
                },
                "title_or": {
                    "type": "string"
                },
                "translators": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "world_art_id": {
                    "type": "string"
                },
                "year": {
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
