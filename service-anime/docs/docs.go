// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-05-26 15:43:31.364117379 +0300 MSK m=+0.048951473

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
            "required": [
                "accessToken"
            ],
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "anime": {
                    "type": "object",
                    "$ref": "#/definitions/handlers.anime"
                },
                "anime_id": {
                    "description": "API version",
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "handlers.anime": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "string"
                },
                "age": {
                    "type": "integer"
                },
                "category": {
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
                "episodes_count": {
                    "type": "integer"
                },
                "genres": {
                    "type": "string"
                },
                "iframe_url": {
                    "type": "string"
                },
                "imdb_rating": {
                    "type": "number"
                },
                "imdb_votes": {
                    "type": "integer"
                },
                "kinopoisk_id": {
                    "type": "integer"
                },
                "kinopoisk_rating": {
                    "type": "number"
                },
                "kinopoisk_votes": {
                    "type": "integer"
                },
                "poster": {
                    "type": "string"
                },
                "seasons_count": {
                    "type": "integer"
                },
                "studios": {
                    "type": "string"
                },
                "tagline": {
                    "type": "string"
                },
                "title_en": {
                    "type": "string"
                },
                "title_ru": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "trailer_iframe_url": {
                    "type": "string"
                },
                "trailer_token": {
                    "type": "string"
                },
                "translator": {
                    "type": "string"
                },
                "translator_id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "world_art_id": {
                    "type": "integer"
                },
                "year": {
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
