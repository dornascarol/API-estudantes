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
        "/estudantes": {
            "get": {
                "description": "Recuperar detalhes dos estudantes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "estudantes"
                ],
                "summary": "Obter uma lista de estudantes",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Registro",
                        "name": "registrar",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "objeto"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "post": {
                "description": "Criar estudante",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "estudantes"
                ],
                "summary": "Criar estudante",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "objeto"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/estudantes/{id}": {
            "get": {
                "description": "Recuperar detalhes do estudante usando ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "estudantes"
                ],
                "summary": "Obter estudante por ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "objeto"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "description": "Atualizar detalhes do estudante",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "estudantes"
                ],
                "summary": "Atualizar estudante",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "objeto"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "Deletar detalhes do estudante",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "estudantes"
                ],
                "summary": "Deletar estudante",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "objeto"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
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
