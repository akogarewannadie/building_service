// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/buildings": {
            "get": {
                "description": "Get list of buildings with optional filters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Get list of buildings",
                "parameters": [
                    {
                        "type": "string",
                        "description": "City",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Year",
                        "name": "year",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of floors",
                        "name": "floors",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.GetBuildingsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new building to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Create a new building",
                "parameters": [
                    {
                        "description": "Building",
                        "name": "building",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Building"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateBuildingResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.CreateBuildingResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "controllers.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "controllers.GetBuildingsResponse": {
            "type": "object",
            "properties": {
                "buildings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Building"
                    }
                }
            }
        },
        "models.Building": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "floors": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Building Service API",
	Description:      "API for managing buildings",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
