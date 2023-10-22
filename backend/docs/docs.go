// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API manga",
            "email": "hanan@asyrawih.id"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/manga": {
            "get": {
                "description": "get all manga",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "manga"
                ],
                "summary": "List manga",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Manga"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create manga by accept body json",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "manga"
                ],
                "summary": "create manga",
                "parameters": [
                    {
                        "description": "manga requested info",
                        "name": "manga",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateMangaRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entity.Manga"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/manga/{id}": {
            "get": {
                "description": "get all manga",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "manga"
                ],
                "summary": "List manga",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Manga Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Manga"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Manga based on id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "manga"
                ],
                "summary": "Remove manga",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Manga Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Manga": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "release_date": {
                    "type": "string"
                },
                "sinopsis": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "total_chapter": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.CreateMangaRequest": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "sinopsis": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "total_chapter": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error_code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Hanan Test",
	Description:      "Manga service api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}