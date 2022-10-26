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
        "/api/v1/mix/transform/save/list": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "三合一"
                ],
                "summary": "查询所有",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/mix/transform/save/reboot": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "三合一"
                ],
                "summary": "重启所有",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/mix/transform/save/start": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "三合一"
                ],
                "summary": "开始",
                "parameters": [
                    {
                        "description": " ",
                        "name": "startMix3Req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.StartMix3Req"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/mix/transform/save/stop": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "三合一"
                ],
                "summary": "停止",
                "parameters": [
                    {
                        "description": " ",
                        "name": "stopReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.StopReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/mix/transform/save/stopAll": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "三合一"
                ],
                "summary": "停止所有",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/other/transform/save/list": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公区"
                ],
                "summary": "查询所有",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/other/transform/save/reboot": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公区"
                ],
                "summary": "重启所有",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/other/transform/save/start": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公区"
                ],
                "summary": "开始",
                "parameters": [
                    {
                        "description": " ",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.StartReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/other/transform/save/stop": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公区"
                ],
                "summary": "停止",
                "parameters": [
                    {
                        "description": " ",
                        "name": "stopReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.StopReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/other/transform/save/stopAll": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公区"
                ],
                "summary": "停止所有",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/single/transform/save/list": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "房间单画面"
                ],
                "summary": "查询所有",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/single/transform/save/reboot": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "房间单画面"
                ],
                "summary": "重启所有",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/single/transform/save/start": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "房间单画面"
                ],
                "summary": "开始",
                "parameters": [
                    {
                        "description": " ",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.StartReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/single/transform/save/stop": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "房间单画面"
                ],
                "summary": "停止",
                "parameters": [
                    {
                        "description": " ",
                        "name": "stopReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.StopReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/single/transform/save/stopAll": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "房间单画面"
                ],
                "summary": "停止所有",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.StartMix3Req": {
            "type": "object",
            "required": [
                "rtspUrlLeft",
                "rtspUrlMiddle",
                "rtspUrlRight"
            ],
            "properties": {
                "roomName": {
                    "type": "string",
                    "example": "1"
                },
                "rtspUrlLeft": {
                    "type": "string",
                    "example": "rtsp://admin:cebon61332433@192.168.99.215:554/cam/realmonitor?channel=1\u0026subtype=1"
                },
                "rtspUrlMiddle": {
                    "type": "string",
                    "example": "rtsp://admin:cebon61332433@192.168.99.215:554/cam/realmonitor?channel=1\u0026subtype=0"
                },
                "rtspUrlRight": {
                    "type": "string",
                    "example": "rtsp://admin:cebon61332433@192.168.99.215:554/cam/realmonitor?channel=1\u0026subtype=1"
                },
                "temperature": {
                    "type": "string",
                    "example": ""
                }
            }
        },
        "dto.StartReq": {
            "type": "object",
            "required": [
                "rtspUrl"
            ],
            "properties": {
                "rtspUrl": {
                    "type": "string",
                    "example": "rtsp://admin:cebon61332433@192.168.99.215:554/cam/realmonitor?channel=1\u0026subtype=0"
                }
            }
        },
        "dto.StopReq": {
            "type": "object",
            "required": [
                "rtmpUrl",
                "taskId"
            ],
            "properties": {
                "rtmpUrl": {
                    "type": "string"
                },
                "taskId": {
                    "type": "integer"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "响应码",
                    "type": "integer"
                },
                "data": {
                    "description": "响应数据"
                },
                "msg": {
                    "description": "响应消息",
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