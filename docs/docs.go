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
        "contact": {
            "name": "Meruya Technology",
            "url": "https://blog.meruyatechnology.com",
            "email": "support@meruyatechnology.com"
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
        "/access/refresh": {
            "post": {
                "security": [
                    {
                        "ClientSecret": []
                    }
                ],
                "description": "Refresh Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "operationId": "auth-refresh-token",
                "parameters": [
                    {
                        "description": "Request payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.LoginResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BadRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/base.UnauthorizedError"
                        }
                    },
                    "403": {
                        "description": "Forbiden",
                        "schema": {
                            "$ref": "#/definitions/base.ForbidenError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/base.NotFoundError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/base.InternalServerError"
                        }
                    }
                }
            }
        },
        "/client/check": {
            "post": {
                "description": "Check a client for authorization",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "operationId": "client-check",
                "parameters": [
                    {
                        "description": "Request payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CheckClientRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.SuccessCreatedResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Client"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BadRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/base.UnauthorizedError"
                        }
                    },
                    "403": {
                        "description": "Forbiden",
                        "schema": {
                            "$ref": "#/definitions/base.ForbidenError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/base.NotFoundError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/base.InternalServerError"
                        }
                    }
                }
            }
        },
        "/client/create": {
            "post": {
                "description": "Create a client for authorization",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "operationId": "client-create",
                "parameters": [
                    {
                        "description": "Request payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreateClientRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.SuccessCreatedResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Client"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BadRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/base.UnauthorizedError"
                        }
                    },
                    "403": {
                        "description": "Forbiden",
                        "schema": {
                            "$ref": "#/definitions/base.ForbidenError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/base.NotFoundError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/base.InternalServerError"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "security": [
                    {
                        "ClientSecret": []
                    }
                ],
                "description": "User login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "operationId": "auth-login",
                "parameters": [
                    {
                        "description": "Request payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/base.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/responses.LoginResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BadRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/base.UnauthorizedError"
                        }
                    },
                    "403": {
                        "description": "Forbiden",
                        "schema": {
                            "$ref": "#/definitions/base.ForbidenError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/base.NotFoundError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/base.InternalServerError"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "security": [
                    {
                        "ClientSecret": []
                    }
                ],
                "description": "Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "operationId": "auth-register",
                "parameters": [
                    {
                        "description": "Request payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/base.SuccessCreatedResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/base.BadRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/base.UnauthorizedError"
                        }
                    },
                    "403": {
                        "description": "Forbiden",
                        "schema": {
                            "$ref": "#/definitions/base.ForbidenError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/base.NotFoundError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/base.InternalServerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "base.BadRequestError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string",
                    "example": "Bad Request"
                }
            }
        },
        "base.ForbidenError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string",
                    "example": "Forbiden Access"
                }
            }
        },
        "base.InternalServerError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string",
                    "example": "Internal Server Error"
                }
            }
        },
        "base.NotFoundError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string",
                    "example": "Not Found"
                }
            }
        },
        "base.SuccessCreatedResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {},
                "message": {
                    "type": "string",
                    "example": "Success"
                }
            }
        },
        "base.SuccessResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {},
                "message": {
                    "type": "string",
                    "example": "Success"
                }
            }
        },
        "base.UnauthorizedError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string",
                    "example": "Unauthorized"
                }
            }
        },
        "entities.Client": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Mobile Client"
                },
                "secret": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTEwLTIzVDEwOjU0OjUwLjQ4OTg0MiswNzowMCJ9.vl8apKYm9UQbj1qaG-BB2eStTEYy1ZJpPoVyuoXDr1k"
                }
            }
        },
        "requests.CheckClientRequest": {
            "type": "object",
            "properties": {
                "secret": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTEwLTIzVDIyOjUxOjEyLjM3NTYwOCswNzowMCJ9.TSIszUF7qDyI3EZM1NtNKVD0zZkwlbwvXfoO5uWhd9k"
                }
            }
        },
        "requests.CreateClientRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Mobile client"
                },
                "secret": {
                    "type": "string",
                    "example": "6wTqKFJ1c3QTJ3dkQ8fsKg=="
                }
            }
        },
        "requests.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "dwikurnianto.mulyadien@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "requests.RefreshTokenRequest": {
            "type": "object",
            "properties": {
                "grantType": {
                    "type": "string",
                    "example": "Client credentials"
                },
                "refreshToken": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIxLTEyLTA0IDA5OjExOjEzLjYyMTQ5OSArMDcwMCBXSUIgbT0rMTA1NjA3LjQ5NzcxNTUwMiJ9.2Cs3nDjqCuHTH_TruMGFmbjIjxIg-HJetJbFbrTBBK0"
                },
                "scope": {
                    "type": "string",
                    "example": "*"
                }
            }
        },
        "requests.RegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "dwikurnianto.mulyadien@gmail.com"
                },
                "firstName": {
                    "type": "string",
                    "example": "Dwi Kurnianto"
                },
                "lastName": {
                    "type": "string",
                    "example": "Mulyadien"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "phone": {
                    "type": "string",
                    "example": "082115100216"
                }
            }
        },
        "responses.LoginResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIxLTEyLTA0IDA5OjExOjEzLjYyMTQ5OSArMDcwMCBXSUIgbT0rMTA1NjA3LjQ5NzcxNTUwMiJ9.2Cs3nDjqCuHTH_TruMGFmbjIjxIg-HJetJbFbrTBBK0"
                },
                "expiresIn": {
                    "type": "integer",
                    "example": 123456
                },
                "refreshToken": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIxLTEyLTEwIDA5OjExOjEzLjYzOTcxMyArMDcwMCBXSUIgbT0rNjI0MDA3LjUxNTkyOTc5MyJ9.G7-lMOIsjJ2Y8zfpiTKME4K1XYvSlyniS3zBMYyOL5k"
                },
                "tokenType": {
                    "type": "string",
                    "example": "Bearer"
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
        "ClientSecret": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	BasePath:    "/api/auth",
	Schemes:     []string{},
	Title:       "Oauth2 API Documentation",
	Description: "Go boiler plate with Oauth2 implementation, documented with Swagger",
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
	swag.Register(swag.Name, &s{})
}
