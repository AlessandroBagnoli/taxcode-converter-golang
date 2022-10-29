// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/v1/taxcode:calculate-person-data": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Calculate data of a person starting from his tax code.",
                "parameters": [
                    {
                        "description": "CalculatePersonDataRequest",
                        "name": "CalculatePersonDataRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.CalculatePersonDataRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.CalculatePersonDataResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/problemdetails.ProblemDetails"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/problemdetails.ProblemDetails"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/problemdetails.ProblemDetails"
                        }
                    }
                }
            }
        },
        "/api/v1/taxcode:calculate-tax-code": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Calculate tax code starting from the data of a person.",
                "parameters": [
                    {
                        "description": "CalculateTaxCodeRequest",
                        "name": "CalculateTaxCodeRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.CalculateTaxCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.CalculateTaxCodeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/problemdetails.ProblemDetails"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/problemdetails.ProblemDetails"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/problemdetails.ProblemDetails"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "problemdetails.ProblemDetails": {
            "type": "object",
            "properties": {
                "detail": {
                    "description": "Detail is a human-readable explanation specific to this\noccurrence of the problem.\nIf present, it ought to focus on helping the client\ncorrect the problem, rather than giving debugging information.",
                    "type": "string"
                },
                "instance": {
                    "description": "Instance is a URI reference that identifies the specific\noccurrence of the problem.  It may or may not yield further\ninformation if dereferenced.",
                    "type": "string"
                },
                "status": {
                    "description": "Status is the HTTP status code ([RFC7231], Section 6)\ngenerated by the origin server for this occurrence of the problem.",
                    "type": "integer"
                },
                "title": {
                    "description": "Title is a short, human-readable summary of the problem\ntype.  It SHOULD NOT change from occurrence to occurrence of the\nproblem, except for purposes of localization (e.g., using\nproactive content negotiation; see [RFC7231], Section 3.4).",
                    "type": "string"
                },
                "type": {
                    "description": "Type is a URI reference [RFC3986] that identifies the\nproblem type. This specification encourages that, when\ndereferenced, it provide human-readable documentation for the\nproblem type (e.g., using HTML [W3C.REC-html5-20141028]).  When\nthis member is not present, its value is assumed to be\n\"about:blank\".",
                    "type": "string"
                }
            }
        },
        "service.CalculatePersonDataRequest": {
            "type": "object",
            "required": [
                "taxCode"
            ],
            "properties": {
                "taxCode": {
                    "type": "string"
                }
            }
        },
        "service.CalculatePersonDataResponse": {
            "type": "object",
            "properties": {
                "birthPlace": {
                    "type": "string"
                },
                "dateOfBirth": {
                    "type": "string",
                    "format": "date"
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "UNKNOWN",
                        "MALE",
                        "FEMALE"
                    ]
                },
                "name": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "taxCode": {
                    "type": "string"
                }
            }
        },
        "service.CalculateTaxCodeRequest": {
            "type": "object",
            "required": [
                "birthPlace",
                "dateOfBirth",
                "gender",
                "name",
                "province",
                "surname"
            ],
            "properties": {
                "birthPlace": {
                    "type": "string"
                },
                "dateOfBirth": {
                    "type": "string",
                    "format": "date"
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "UNKNOWN",
                        "MALE",
                        "FEMALE"
                    ]
                },
                "name": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "service.CalculateTaxCodeResponse": {
            "type": "object",
            "properties": {
                "taxCode": {
                    "type": "string"
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
	Title:            "taxcode-converter",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
