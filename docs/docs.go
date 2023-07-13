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
        "/articles": {
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
                    "Articles"
                ],
                "summary": "Get random article.",
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
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "Create article.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "the body to create article",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.SaveArticleInput"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/articles/{id}": {
            "put": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "Update article.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "the body to update article",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.SaveArticleInput"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Article id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "Delete article.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Article id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/articles/{username}": {
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
                    "Articles"
                ],
                "summary": "Get user article.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit returning value",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Paging",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "user username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/articles/{username}/{id}": {
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
                    "Articles"
                ],
                "summary": "Get detail article.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "user username",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/articles/{username}/{id}/comment": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "Comment article.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Article id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to comment an article",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/article.CommentInput"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/articles/{username}/{id}/like": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "Likes article.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Article id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
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
        "/followers": {
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
                    "Followers"
                ],
                "summary": "Get follower user.",
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
        },
        "/following": {
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
                    "Following"
                ],
                "summary": "Get following user.",
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
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Following"
                ],
                "summary": "Follow a user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "the body to follow a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/following.FollowUserInput"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/following/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Following"
                ],
                "summary": "Delete a following user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Following id",
                        "name": "id",
                        "in": "path",
                        "required": true
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
        },
        "/profile/change-password": {
            "post": {
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
                "summary": "Change password user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "the body to change password",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/profile.ChangePasswordInput"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "article.CommentInput": {
            "type": "object",
            "required": [
                "comment"
            ],
            "properties": {
                "comment": {
                    "type": "string"
                }
            }
        },
        "article.SaveArticleInput": {
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
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
        },
        "following.FollowUserInput": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "profile.ChangePasswordInput": {
            "type": "object",
            "required": [
                "old_password",
                "password",
                "password_confirm"
            ],
            "properties": {
                "old_password": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "password_confirm": {
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
