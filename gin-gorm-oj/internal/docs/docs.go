// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "/admin/problem-create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "问题创建",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "ProblemBasic",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.ProblemBasic"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/problem-modify": {
            "put": {
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "问题修改",
                "parameters": [
                    {
                        "type": "string",
                        "description": "authorization",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "ProblemBasic",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/define.ProblemBasic"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/problem-detail": {
            "get": {
                "tags": [
                    "公共方法"
                ],
                "summary": "问题详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "problem identity",
                        "name": "identity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/problem-list": {
            "get": {
                "tags": [
                    "公共方法"
                ],
                "summary": "问题列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "category_identity",
                        "name": "category_identity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "define.ProblemBasic": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "问题内容",
                    "type": "string"
                },
                "identity": {
                    "description": "问题表的唯一标识",
                    "type": "string"
                },
                "max_mem": {
                    "description": "最大运行内存",
                    "type": "integer"
                },
                "max_runtime": {
                    "description": "最大运行时长",
                    "type": "integer"
                },
                "problem_categories": {
                    "description": "关联问题分类表",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "test_cases": {
                    "description": "关联测试用例表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/define.TestCase"
                    }
                },
                "title": {
                    "description": "问题标题",
                    "type": "string"
                }
            }
        },
        "define.TestCase": {
            "type": "object",
            "properties": {
                "input": {
                    "description": "输入",
                    "type": "string"
                },
                "output": {
                    "description": "输出",
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
