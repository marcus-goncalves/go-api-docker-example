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
        "/support/ping": {
            "get": {
                "description": "Check if API is up returning \"OK\"",
                "tags": [
                    "Support"
                ],
                "responses": {}
            }
        },
        "/user/add": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "Create a new User",
                "parameters": [
                    {
                        "description": "Add user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.ResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.ResponseBody"
                        }
                    }
                }
            }
        },
        "/user/del/{id}": {
            "delete": {
                "description": "Delete an user by its id",
                "tags": [
                    "User"
                ],
                "summary": "Delete an user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID User",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.ResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.ResponseBody"
                        }
                    }
                }
            }
        },
        "/user/find": {
            "get": {
                "description": "Return a user by its id, name or age as query params",
                "tags": [
                    "User"
                ],
                "summary": "Return a user by id, name or age",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search by id",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search by name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search by age",
                        "name": "age",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.ResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.ResponseBody"
                        }
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "description": "Return a list of user objects",
                "tags": [
                    "User"
                ],
                "summary": "Return all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.ResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.ResponseBody"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "routes.ResponseBody": {
            "type": "object",
            "properties": {
                "body": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "routes.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
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
