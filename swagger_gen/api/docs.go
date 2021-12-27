// Package api GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package api

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Alexander Sharov",
            "url": "http://github.com/kvendingoldo",
            "email": "kvendingoldo@gmail.com"
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
        "/distance": {
            "get": {
                "description": "TODO",
                "tags": [
                    "location"
                ],
                "summary": "Get user distance by range",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/locations.NewDistanceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/location": {
            "put": {
                "description": "Update location on the system",
                "tags": [
                    "location"
                ],
                "summary": "Update location",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/locations.NewLocationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get all users on the system",
                "tags": [
                    "location"
                ],
                "summary": "Get location",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/locations.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/locations.MessageResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "locations.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "locations.NewDistanceRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "800"
                },
                "range": {
                    "type": "integer",
                    "minimum": 0,
                    "example": 1
                },
                "username": {
                    "type": "string",
                    "example": "Bill"
                }
            }
        },
        "locations.NewLocationRequest": {
            "type": "object",
            "required": [
                "coordinates"
            ],
            "properties": {
                "coordinates": {
                    "type": "string",
                    "example": "39.12355, 27.64538"
                },
                "uid": {
                    "type": "integer",
                    "example": 800
                },
                "username": {
                    "type": "string",
                    "example": "Bill"
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
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "GU location service",
	Description: "Documentation's GU location service",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
