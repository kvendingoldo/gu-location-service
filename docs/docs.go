// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/distance": {
            "get": {
                "description": "Returns distance traveled by a person within some date/time range (in days).",
                "tags": [
                    "location"
                ],
                "summary": "Returns distance traveled by a person within some date/time range (in days).",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of user",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "time range",
                        "name": "range",
                        "in": "query"
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
                "description": "Update current user location by the username/uid.",
                "tags": [
                    "location"
                ],
                "summary": "Update current user location by the username/uid.",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.NewLocationRequest"
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
                "description": "Search for users in some location within the provided radius (with pagination).",
                "tags": [
                    "location"
                ],
                "summary": "Search users in some location within the provided radius.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Center latitude",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Center longitude",
                        "name": "lon",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "radius",
                        "name": "radius",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "radius (m|km|mi|ft)",
                        "name": "units",
                        "in": "query"
                    }
                ],
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
                            "$ref": "#/definitions/v1.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.MessageResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "v1.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.NewLocationRequest": {
            "type": "object",
            "required": [
                "lat",
                "lon"
            ],
            "properties": {
                "lat": {
                    "type": "number",
                    "example": 39.12355
                },
                "lon": {
                    "type": "number",
                    "example": 27.64538
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
