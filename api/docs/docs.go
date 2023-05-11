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
        "/dataset": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dataset"
                ],
                "summary": "Create a new dataset",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dataset.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Dataset"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            }
        },
        "/dataset/{datasetName}": {
            "delete": {
                "tags": [
                    "Dataset"
                ],
                "summary": "Remove a dataset",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dataset name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            }
        },
        "/dataset/{name}/piece": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dataset"
                ],
                "summary": "Register a CAR file piece with a dataset",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dataset name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dataset.AddPieceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Car"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            }
        },
        "/dataset/{name}/source": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dataset"
                ],
                "summary": "Add a source to a dataset",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dataset name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dataset.AddSourceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Source"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            }
        },
        "/dataset/{name}/source/{sourcePath}": {
            "delete": {
                "tags": [
                    "Dataset"
                ],
                "summary": "Remove a source from a dataset",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dataset name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Source path",
                        "name": "sourcePath",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            }
        },
        "/dataset/{name}/sources": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dataset"
                ],
                "summary": "List all sources for a dataset",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dataset name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Source"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            }
        },
        "/datasets": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dataset"
                ],
                "summary": "List all datasets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Dataset"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            }
        },
        "/deal/send_manual": {
            "post": {
                "description": "Send a manual deal proposal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deal"
                ],
                "summary": "Send a manual deal proposal",
                "parameters": [
                    {
                        "description": "Proposal",
                        "name": "proposal",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/deal.Proposal"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/init": {
            "post": {
                "tags": [
                    "Global"
                ],
                "summary": "Initialize the database",
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            }
        },
        "/wallet": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "Import a private key",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/wallet.ImportRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            }
        },
        "/wallets": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "List all imported wallets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Wallet"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dataset.AddPieceRequest": {
            "type": "object",
            "properties": {
                "filePath": {
                    "type": "string"
                },
                "fileSize": {
                    "type": "integer"
                },
                "pieceCID": {
                    "type": "string"
                },
                "pieceSize": {
                    "type": "string"
                },
                "rootCID": {
                    "type": "string"
                }
            }
        },
        "dataset.AddSourceRequest": {
            "type": "object",
            "properties": {
                "httpheaders": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "pushOnly": {
                    "type": "boolean"
                },
                "s3AccessKeyID": {
                    "type": "string"
                },
                "s3Endpoint": {
                    "type": "string"
                },
                "s3Region": {
                    "type": "string"
                },
                "s3SecretAccessKey": {
                    "type": "string"
                },
                "scanInterval": {
                    "$ref": "#/definitions/time.Duration"
                },
                "sourcePath": {
                    "type": "string"
                }
            }
        },
        "dataset.CreateRequest": {
            "type": "object",
            "properties": {
                "maxSize": {
                    "type": "string"
                },
                "minSize": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "outputDirs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "pieceSize": {
                    "type": "string"
                },
                "recipients": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "script": {
                    "type": "string"
                }
            }
        },
        "deal.Proposal": {
            "type": "object",
            "properties": {
                "clientAddress": {
                    "type": "string"
                },
                "durationDays": {
                    "type": "number"
                },
                "fileSize": {
                    "type": "integer"
                },
                "httpHeaders": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ipni": {
                    "type": "boolean"
                },
                "keepUnsealed": {
                    "type": "boolean"
                },
                "label": {
                    "type": "string"
                },
                "pieceCID": {
                    "type": "string"
                },
                "pieceSize": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "providerID": {
                    "type": "string"
                },
                "startDelayDays": {
                    "type": "number"
                },
                "urlTemplate": {
                    "type": "string"
                },
                "verified": {
                    "type": "boolean"
                }
            }
        },
        "handler.HTTPError": {
            "type": "object",
            "properties": {
                "err": {
                    "type": "string"
                }
            }
        },
        "model.Car": {
            "type": "object",
            "properties": {
                "chunkID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "datasetID": {
                    "type": "integer"
                },
                "filePath": {
                    "type": "string"
                },
                "fileSize": {
                    "type": "integer"
                },
                "header": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "omitempty": {
                    "$ref": "#/definitions/model.Chunk"
                },
                "pieceCID": {
                    "type": "string"
                },
                "pieceSize": {
                    "type": "integer"
                },
                "rootCID": {
                    "type": "string"
                }
            }
        },
        "model.Chunk": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Item"
                    }
                },
                "packingState": {
                    "$ref": "#/definitions/model.WorkState"
                },
                "packingWorker": {
                    "$ref": "#/definitions/model.Worker"
                },
                "packingWorkerID": {
                    "type": "string"
                },
                "source": {
                    "$ref": "#/definitions/model.Source"
                },
                "sourceID": {
                    "type": "integer"
                }
            }
        },
        "model.Dataset": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "encryptionRecipients": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "encryptionScript": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "maxSize": {
                    "type": "integer"
                },
                "minSize": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "outputDirs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "pieceSize": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.Directory": {
            "type": "object",
            "properties": {
                "cid": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "omitempty": {
                    "$ref": "#/definitions/model.Directory"
                },
                "parentID": {
                    "type": "integer"
                }
            }
        },
        "model.Item": {
            "type": "object",
            "properties": {
                "chunkID": {
                    "type": "integer"
                },
                "cid": {
                    "type": "string"
                },
                "directoryID": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastModified": {
                    "type": "string"
                },
                "length": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "omitempty": {
                    "$ref": "#/definitions/model.Directory"
                },
                "path": {
                    "type": "string"
                },
                "scannedAt": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "sourceID": {
                    "type": "integer"
                },
                "type": {
                    "$ref": "#/definitions/model.ItemType"
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "model.ItemType": {
            "type": "string",
            "enum": [
                "file",
                "url",
                "s3object"
            ],
            "x-enum-varnames": [
                "File",
                "URL",
                "S3Object"
            ]
        },
        "model.Metadata": {
            "type": "object",
            "additionalProperties": true
        },
        "model.Source": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "datasetID": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastScannedTimestamp": {
                    "type": "integer"
                },
                "metadata": {
                    "$ref": "#/definitions/model.Metadata"
                },
                "path": {
                    "type": "string"
                },
                "pushOnly": {
                    "type": "boolean"
                },
                "rootDirectoryId": {
                    "type": "integer"
                },
                "scanIntervalSeconds": {
                    "type": "integer"
                },
                "scanningState": {
                    "$ref": "#/definitions/model.WorkState"
                },
                "scanningWorkerId": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/model.SourceType"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.SourceType": {
            "type": "string",
            "enum": [
                "dir",
                "website",
                "s3path",
                "upload"
            ],
            "x-enum-varnames": [
                "Dir",
                "Website",
                "S3Path",
                "Upload"
            ]
        },
        "model.Wallet": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "privateKey": {
                    "type": "string"
                }
            }
        },
        "model.WorkState": {
            "type": "string",
            "enum": [
                "created",
                "ready",
                "processing",
                "complete",
                "error"
            ],
            "x-enum-varnames": [
                "Created",
                "Ready",
                "Processing",
                "Complete",
                "Error"
            ]
        },
        "model.WorkType": {
            "type": "string",
            "enum": [
                "scan",
                "deal_making",
                "packing"
            ],
            "x-enum-varnames": [
                "Scan",
                "DealMaking",
                "Packing"
            ]
        },
        "model.Worker": {
            "type": "object",
            "properties": {
                "hostname": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastHeartbeat": {
                    "type": "string"
                },
                "workType": {
                    "$ref": "#/definitions/model.WorkType"
                },
                "workingOn": {
                    "type": "string"
                }
            }
        },
        "time.Duration": {
            "type": "integer",
            "enum": [
                -9223372036854775808,
                9223372036854775807,
                1,
                1000,
                1000000,
                1000000000,
                60000000000,
                3600000000000
            ],
            "x-enum-varnames": [
                "minDuration",
                "maxDuration",
                "Nanosecond",
                "Microsecond",
                "Millisecond",
                "Second",
                "Minute",
                "Hour"
            ]
        },
        "wallet.ImportRequest": {
            "type": "object",
            "properties": {
                "privateKey": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "beta",
	Host:             "localhost:9090",
	BasePath:         "/admin/api",
	Schemes:          []string{},
	Title:            "Singularity API",
	Description:      "This is the API for Singularity, a tool for large-scale clients with PB-scale data onboarding to Filecoin network.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
