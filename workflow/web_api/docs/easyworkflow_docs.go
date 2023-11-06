// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplateeasyworkflow = `{
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
        "/def/get": {
            "get": {
                "description": "返回的是Node数组，流程是由N个Node组成的",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "流程定义"
                ],
                "summary": "获取流程定义",
                "parameters": [
                    {
                        "type": "string",
                        "example": "1",
                        "description": "流程ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "流程定义",
                        "schema": {
                            "$ref": "#/definitions/model.Process"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/def/list": {
            "get": {
                "description": "引擎可能被多个系统、组件等使用，source表示从哪个来源创建的流程",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "流程定义"
                ],
                "summary": "获取特定source下所有流程",
                "parameters": [
                    {
                        "type": "string",
                        "example": "办公系统",
                        "description": "来源",
                        "name": "source",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/database.ProcDef"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/def/save": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "流程定义"
                ],
                "summary": "流程定义保存/升级",
                "parameters": [
                    {
                        "type": "string",
                        "example": "json字符串",
                        "description": "流程定义资源(json)",
                        "name": "Resource",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "0001",
                        "description": "创建者ID",
                        "name": "CreateUserID",
                        "in": "formData",
                        "required": true
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
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/inst/revoke": {
            "post": {
                "description": "注意，Force 是否强制撤销，若为false,则只有流程回到发起人这里才能撤销",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "流程实例"
                ],
                "summary": "撤销流程",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "流程实例ID",
                        "name": "InstanceID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"U001\"",
                        "description": "撤销发起用户ID",
                        "name": "RevokeUserID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "是否强制撤销",
                        "name": "Force",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/inst/start": {
            "post": {
                "description": "注意，VariablesJson格式是key-value对象集合:[{\"Key\":\"starter\",\"Value\":\"U0001\"}]",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "流程实例"
                ],
                "summary": "开始流程",
                "parameters": [
                    {
                        "type": "string",
                        "example": "1",
                        "description": "流程ID",
                        "name": "ProcessID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"订单001\"",
                        "description": "业务ID",
                        "name": "BusinessID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"家中有事请假三天,请领导批准\"",
                        "description": "评论意见",
                        "name": "Comment",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "example": "[{\"Key\":\"starter\",\"Value\":\"U0001\"}]",
                        "description": "变量(Json)",
                        "name": "VariablesJson",
                        "in": "formData"
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
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/inst/start/by": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "流程实例"
                ],
                "summary": "获取起始人为特定用户的流程实例",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"U001\"",
                        "description": "用户ID",
                        "name": "userid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"请假\"",
                        "description": "指定流程名称，非必填",
                        "name": "procname",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 0,
                        "description": "分页用,开始index",
                        "name": "idx",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "example": 0,
                        "description": "分页用,最大返回行数",
                        "name": "rows",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "流程实例列表",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Instance"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/inst/task_history": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "流程实例"
                ],
                "summary": "流程实例中任务执行记录",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "流程实例ID",
                        "name": "instid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "任务列表",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Task"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/action": {
            "get": {
                "description": "前端无法提前知道当前任务可以做哪些操作，此方法目的是解决这个困扰",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务"
                ],
                "summary": "当前任务可以执行哪些操作",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "任务ID",
                        "name": "taskid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "可执行任务",
                        "schema": {
                            "$ref": "#/definitions/model.TaskAction"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/finished": {
            "get": {
                "description": "返回的是任务数组",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务"
                ],
                "summary": "获取已办任务",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"U001\"",
                        "description": "用户ID",
                        "name": "userid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"请假\"",
                        "description": "指定流程名称，非必填",
                        "name": "procname",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "忽略由我开启流程,而生成处理人是我自己的任务",
                        "name": "ignorestartbyme",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "example": true,
                        "description": "是否按照任务完成时间升序排列",
                        "name": "asc",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "example": 0,
                        "description": "分页用,开始index",
                        "name": "idx",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "example": 0,
                        "description": "分页用,最大返回行数",
                        "name": "rows",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Task"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/info": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务"
                ],
                "summary": "任务信息",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "任务ID",
                        "name": "taskid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "任务信息",
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/pass": {
            "post": {
                "description": "任务通过后根据流程定义，进入下一个节点进行处理",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务"
                ],
                "summary": "任务通过",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "任务ID",
                        "name": "TaskID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"同意请假\"",
                        "description": "评论意见",
                        "name": "Comment",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "example": "\"{\"User\":\"001\"}\"",
                        "description": "变量(Json)",
                        "name": "VariablesJson",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/pass/directly": {
            "post": {
                "description": "此功能只有在非会签节点时才能使用",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务"
                ],
                "summary": "任务通过后流程直接返回到上一个驳回我的节点",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "任务ID",
                        "name": "TaskID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"同意请假\"",
                        "description": "评论意见",
                        "name": "Comment",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "example": "\"{\"User\":\"001\"}\"",
                        "description": "变量(Json)",
                        "name": "VariablesJson",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/reject": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务"
                ],
                "summary": "任务驳回",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "任务ID",
                        "name": "TaskID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"不同意\"",
                        "description": "评论意见",
                        "name": "Comment",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "example": "\"{\"User\":\"001\"}\"",
                        "description": "变量(Json)",
                        "name": "VariablesJson",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/reject/free": {
            "post": {
                "description": "驳回到上游任意一个节点",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务"
                ],
                "summary": "自由任务驳回",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "任务ID",
                        "name": "TaskID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"不同意\"",
                        "description": "评论意见",
                        "name": "Comment",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "example": "\"{\"User\":\"001\"}\"",
                        "description": "变量(Json)",
                        "name": "VariablesJson",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "example": "\"流程开始节点\"",
                        "description": "驳回到哪个节点",
                        "name": "RejectToNodeID",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/todo": {
            "get": {
                "description": "返回的是任务数组",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务"
                ],
                "summary": "获取待办任务",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"U001\"",
                        "description": "用户ID",
                        "name": "userid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"请假\"",
                        "description": "指定流程名称，非必填",
                        "name": "procname",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "example": true,
                        "description": "是否按照任务生成时间升序排列",
                        "name": "asc",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "example": 0,
                        "description": "分页用,开始index",
                        "name": "idx",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "example": 0,
                        "description": "分页用,最大返回行数",
                        "name": "rows",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Task"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/transfer": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务"
                ],
                "summary": "将任务转交给他人处理",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "任务ID",
                        "name": "TaskID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"U001,U002,U003\"",
                        "description": "用户,多个用户使用逗号分隔",
                        "name": "Users",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/task/upstream": {
            "get": {
                "description": "此功能为自由驳回使用",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务"
                ],
                "summary": "获取本任务所在节点的所有上游节点",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "任务ID",
                        "name": "taskid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Node"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "database.ProcDef": {
            "type": "object",
            "properties": {
                "creatTime": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "resource": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "model.Condition": {
            "type": "object",
            "properties": {
                "expression": {
                    "description": "条件表达式",
                    "type": "string"
                },
                "nodeID": {
                    "description": "满足条件后转跳到哪个节点",
                    "type": "string"
                }
            }
        },
        "model.HybridGateway": {
            "type": "object",
            "properties": {
                "conditions": {
                    "description": "条件判断节点",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Condition"
                    }
                },
                "inevitableNodes": {
                    "description": "必然执行的节点",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "waitForAllPrevNode": {
                    "description": "0:等于包含网关，只要上级节点有一个完成，就可以往下走   1:等于并行网关，必须要上级节点全部完成才能往下走",
                    "type": "integer"
                }
            }
        },
        "model.Instance": {
            "type": "object",
            "properties": {
                "businessID": {
                    "description": "业务ID",
                    "type": "string"
                },
                "createTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "currentNodeID": {
                    "description": "当前进行节点ID",
                    "type": "string"
                },
                "procID": {
                    "description": "流程ID",
                    "type": "integer"
                },
                "procInstID": {
                    "description": "流程实例ID",
                    "type": "integer"
                },
                "procName": {
                    "description": "流程名称",
                    "type": "string"
                },
                "procVersion": {
                    "description": "流程版本号",
                    "type": "integer"
                },
                "starter": {
                    "description": "流程发起人用户ID",
                    "type": "string"
                },
                "status": {
                    "description": "0:未完成(审批中) 1:已完成(通过) 2:撤销",
                    "type": "integer"
                }
            }
        },
        "model.Node": {
            "type": "object",
            "properties": {
                "gwconfig": {
                    "description": "网关。只有在节点类型为GateWay的情况下此字段才会有值",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.HybridGateway"
                        }
                    ]
                },
                "isCosigned": {
                    "description": "是否会签  只有任务节点才会用到，会签的情况下需要所有任务通过才能进行下一节点，只要有一人反对，则整个节点驳回",
                    "type": "integer"
                },
                "nodeEndEvents": {
                    "description": "节点结束时触发的事件",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "nodeID": {
                    "description": "节点名称",
                    "type": "string"
                },
                "nodeName": {
                    "description": "节点名字",
                    "type": "string"
                },
                "nodeStartEvents": {
                    "description": "节点开始时触发的事件",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "nodeType": {
                    "description": "节点类型 0:开始节点 1:任务节点,指的是需要人完成的节点 2:网关 3:结束节点",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeType"
                        }
                    ]
                },
                "prevNodeIDs": {
                    "description": "上级节点(不管是任务节点还是结束节点，因为分支的存在，所以它的上级节点可能都会有多个)",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "roles": {
                    "description": "节点处理角色数组。注意，因为系统无法预先知道角色中存在多少用户，所以必须用StartEvents解析角色，将角色中的用户加到UserIDs中",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "taskFinishEvents": {
                    "description": "任务完成(通过、驳回)时触发的事件。节点中可能产生N个任务，任务完成事件，会在每个任务完成时触发",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "userIDs": {
                    "description": "节点处理人数组",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.NodeType": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3
            ],
            "x-enum-comments": {
                "EndNode": "结束节点,结束节点不需要人参与，到了此节点，则流程实例完成",
                "GateWayNode": "参考activiti的网关.此项目中使用混合网关,等于activiti中排他、并行网关、包含网关的混合体",
                "RootNode": "开始(根)节点",
                "TaskNode": "任务节点,指的是需要人完成的节点"
            },
            "x-enum-varnames": [
                "RootNode",
                "TaskNode",
                "GateWayNode",
                "EndNode"
            ]
        },
        "model.Process": {
            "type": "object",
            "properties": {
                "nodes": {
                    "description": "节点",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Node"
                    }
                },
                "processName": {
                    "description": "流程名",
                    "type": "string"
                },
                "revokeEvents": {
                    "description": "流程撤销事件.在流程实例撤销时触发",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "source": {
                    "description": "来源(引擎可能被多个系统、组件等使用，这里记下从哪个来源创建的流程",
                    "type": "string"
                }
            }
        },
        "model.Task": {
            "type": "object",
            "properties": {
                "batchCode": {
                    "description": "批次码.节点会被驳回，一个节点可能产生多批task,用此码做分别",
                    "type": "string"
                },
                "businessID": {
                    "description": "业务ID",
                    "type": "string"
                },
                "comment": {
                    "description": "评论意见",
                    "type": "string"
                },
                "createTime": {
                    "description": "任务创建时间",
                    "type": "string"
                },
                "finishedTime": {
                    "description": "处理任务时间",
                    "type": "string"
                },
                "isCosigned": {
                    "description": "0:任意一人通过即可 1:会签",
                    "type": "integer"
                },
                "isFinished": {
                    "description": "0:任务未完成 1:处理完成",
                    "type": "integer"
                },
                "nodeID": {
                    "description": "节点ID",
                    "type": "string"
                },
                "nodeName": {
                    "description": "节点名称",
                    "type": "string"
                },
                "prevNodeID": {
                    "description": "上一节点ID",
                    "type": "string"
                },
                "procID": {
                    "description": "流程ID",
                    "type": "integer"
                },
                "procInstCreateTime": {
                    "description": "流程实例创建时间",
                    "type": "string"
                },
                "procInstID": {
                    "description": "流程实例ID",
                    "type": "integer"
                },
                "procName": {
                    "description": "流程名称",
                    "type": "string"
                },
                "starter": {
                    "description": "流程起始人",
                    "type": "string"
                },
                "status": {
                    "description": "任务状态:0:初始 1:通过 2:驳回",
                    "type": "integer"
                },
                "taskID": {
                    "description": "任务ID",
                    "type": "integer"
                },
                "userID": {
                    "description": "分配用户ID",
                    "type": "string"
                }
            }
        },
        "model.TaskAction": {
            "type": "object",
            "properties": {
                "canDirectlyToWhoRejectedMe": {
                    "description": "任务可以执行“直接提交到上一个驳回我的节点”",
                    "type": "boolean"
                },
                "canFreeRejectToUpstreamNode": {
                    "description": "任务可以执行“自由驳回”",
                    "type": "boolean"
                },
                "canPass": {
                    "description": "任务可以执行“通过”",
                    "type": "boolean"
                },
                "canReject": {
                    "description": "任务可以执行“驳回”",
                    "type": "boolean"
                },
                "canRevoke": {
                    "description": "任务可以执行\"撤销\"",
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfoeasyworkflow holds exported Swagger Info so clients can modify it
var SwaggerInfoeasyworkflow = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "easyworkflow",
	SwaggerTemplate:  docTemplateeasyworkflow,
}

func init() {
	swag.Register(SwaggerInfoeasyworkflow.InstanceName(), SwaggerInfoeasyworkflow)
}
