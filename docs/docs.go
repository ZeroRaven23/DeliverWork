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
        "/auth": {
            "put": {
                "description": "Autenticar o usuario atraves do email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Autenticar User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email do usuario",
                        "name": "Email",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Token do usuario",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Result\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/loggout": {
            "put": {
                "description": "Altera o modo de login para false",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "Desconectar o usuario",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email do usuario",
                        "name": "Email",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Result\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "put": {
                "description": "Altera o modo de login para true",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "Conectar o usuario",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email do usuario",
                        "name": "Email",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Senha do usuario",
                        "name": "Password",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Result\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Criar um usuario a partir das novas infos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Criar um novo usuario",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email do usuario",
                        "name": "Email",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Nome do usuario",
                        "name": "Name",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Senha do usuario",
                        "name": "Password",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Result\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/token": {
            "post": {
                "description": "Criar o token e ser enviado ao usuario",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Criar token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email do usuario",
                        "name": "Email",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Result\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Buscar um usuario",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Buscar um usuario",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email do usuario",
                        "name": "Email",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Result\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
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
