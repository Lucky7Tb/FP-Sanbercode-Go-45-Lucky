// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/auth/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login.",
                "parameters": [
                    {
                        "description": "the body to login",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginInput"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/auth/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register.",
                "parameters": [
                    {
                        "description": "the body to register",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterInput"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/profile": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "auth.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "user_name"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "auth.RegisterInput": {
            "type": "object",
            "required": [
                "full_name",
                "password",
                "password_confirm",
                "user_name"
            ],
            "properties": {
                "full_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "password_confirm": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
