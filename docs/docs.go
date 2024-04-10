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
        "/api/v1/admin/devices": {
            "get": {
                "description": "Show all devices",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices"
                ],
                "summary": "Show Devices",
                "operationId": "devices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.DeviceCountResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/domains": {
            "get": {
                "description": "Show all domains",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "domains"
                ],
                "summary": "Show Domains",
                "operationId": "domains",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.DomainCountResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/profiles": {
            "get": {
                "description": "Show all profiles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profiles"
                ],
                "summary": "Show Profiles",
                "operationId": "profiles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.ProfileCountResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.CIRAConfig": {
            "type": "object",
            "properties": {
                "authMethod": {
                    "type": "integer",
                    "example": 2
                },
                "commonName": {
                    "type": "string",
                    "example": "example.com"
                },
                "configName": {
                    "type": "string",
                    "example": "My CIRA Config"
                },
                "mpsPort": {
                    "type": "integer",
                    "example": 443
                },
                "mpsRootCertificate": {
                    "type": "string",
                    "example": "-----BEGIN CERTIFICATE-----\n..."
                },
                "mpsServerAddress": {
                    "type": "string",
                    "example": "https://example.com"
                },
                "password": {
                    "type": "string",
                    "example": "my_password"
                },
                "proxyDetails": {
                    "type": "string",
                    "example": "http://example.com"
                },
                "regeneratePassword": {
                    "type": "boolean",
                    "example": true
                },
                "serverAddressFormat": {
                    "type": "integer",
                    "example": 201
                },
                "tenantId": {
                    "type": "string",
                    "example": "abc123"
                },
                "username": {
                    "type": "string",
                    "example": "my_username"
                },
                "version": {
                    "type": "string",
                    "example": "1.0.0"
                }
            }
        },
        "entity.CertCreationResult": {
            "type": "object",
            "properties": {
                "cert": {
                    "type": "string"
                },
                "certBin": {
                    "type": "string"
                },
                "checked": {
                    "type": "boolean",
                    "example": true
                },
                "h:": {
                    "type": "string"
                },
                "key": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "pem": {
                    "type": "string"
                },
                "privateKey": {
                    "type": "string"
                },
                "privateKeyBin": {
                    "type": "string"
                }
            }
        },
        "entity.Device": {
            "type": "object",
            "properties": {
                "allowSelfSigned": {
                    "type": "boolean"
                },
                "connectionStatus": {
                    "type": "boolean"
                },
                "deviceInfo": {
                    "$ref": "#/definitions/entity.DeviceInfo"
                },
                "dnsSuffix": {
                    "type": "string"
                },
                "friendlyName": {
                    "type": "string"
                },
                "guid": {
                    "type": "string"
                },
                "hostname": {
                    "type": "string"
                },
                "lastConnected": {
                    "type": "string"
                },
                "lastDisconnected": {
                    "type": "string"
                },
                "lastSeen": {
                    "type": "string"
                },
                "mpsInstance": {
                    "type": "string"
                },
                "mpsusername": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "tenantId": {
                    "type": "string"
                },
                "useTLS": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.DeviceInfo": {
            "type": "object",
            "properties": {
                "currentMode": {
                    "type": "string"
                },
                "features": {
                    "type": "string"
                },
                "fwBuild": {
                    "type": "string"
                },
                "fwSku": {
                    "type": "string"
                },
                "fwVersion": {
                    "type": "string"
                },
                "ipAddress": {
                    "type": "string"
                },
                "lastUpdated": {
                    "type": "string"
                }
            }
        },
        "entity.Domain": {
            "type": "object",
            "properties": {
                "domainSuffix": {
                    "type": "string",
                    "example": "example.com"
                },
                "profileName": {
                    "type": "string",
                    "example": "My Profile"
                },
                "provisioningCert": {
                    "type": "string",
                    "example": "-----BEGIN CERTIFICATE-----\n..."
                },
                "provisioningCertPassword": {
                    "type": "string",
                    "example": "my_password"
                },
                "provisioningCertStorageFormat": {
                    "type": "string",
                    "example": "PKCS12"
                },
                "tenantId": {
                    "type": "string",
                    "example": "abc123"
                },
                "version": {
                    "type": "string",
                    "example": "1.0.0"
                }
            }
        },
        "entity.IEEE8021xConfig": {
            "type": "object",
            "properties": {
                "activeInS0": {
                    "type": "boolean",
                    "example": true
                },
                "authenticationProtocol": {
                    "type": "integer",
                    "example": 1
                },
                "domain": {
                    "type": "string",
                    "example": "example.com"
                },
                "password": {
                    "type": "string",
                    "example": "my_password"
                },
                "profileName": {
                    "type": "string",
                    "example": "My Profile"
                },
                "pxeTimeout": {
                    "type": "integer",
                    "example": 60
                },
                "roamingIdentity": {
                    "type": "string",
                    "example": "my_roaming_identity"
                },
                "serverName": {
                    "type": "string",
                    "example": "example.com"
                },
                "tenantId": {
                    "type": "string",
                    "example": "abc123"
                },
                "username": {
                    "type": "string",
                    "example": "my_username"
                },
                "version": {
                    "type": "string",
                    "example": "1.0.0"
                },
                "wiredInterface": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "entity.Profile": {
            "type": "object",
            "properties": {
                "activation": {
                    "type": "string",
                    "example": "activate"
                },
                "amtPassword": {
                    "type": "string",
                    "example": "my_password"
                },
                "ciraConfigName": {
                    "type": "string",
                    "example": "My CIRA Config"
                },
                "ciraConfigObject": {
                    "$ref": "#/definitions/entity.CIRAConfig"
                },
                "dhcpEnabled": {
                    "type": "boolean",
                    "example": true
                },
                "generateRandomMEBxPassword": {
                    "type": "boolean",
                    "example": true
                },
                "generateRandomPassword": {
                    "type": "boolean",
                    "example": true
                },
                "iderEnabled": {
                    "type": "boolean",
                    "example": true
                },
                "ieee8021xProfileName": {
                    "type": "string",
                    "example": "My Profile"
                },
                "ieee8021xProfileObject": {
                    "$ref": "#/definitions/entity.IEEE8021xConfig"
                },
                "ipSyncEnabled": {
                    "type": "boolean",
                    "example": true
                },
                "kvmEnabled": {
                    "type": "boolean",
                    "example": true
                },
                "localWifiSyncEnabled": {
                    "type": "boolean",
                    "example": true
                },
                "mebxPassword": {
                    "type": "string",
                    "example": "my_password"
                },
                "profileName": {
                    "type": "string",
                    "example": "My Profile"
                },
                "solEnabled": {
                    "type": "boolean",
                    "example": true
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "tag1",
                        "tag2"
                    ]
                },
                "tenantId": {
                    "type": "string",
                    "example": "abc123"
                },
                "tlsCerts": {
                    "$ref": "#/definitions/entity.TLSCerts"
                },
                "tlsMode": {
                    "type": "integer",
                    "example": 1
                },
                "tlsSigningAuthority": {
                    "type": "string",
                    "example": "SelfSigned"
                },
                "userConsent": {
                    "type": "string",
                    "example": "All"
                },
                "version": {
                    "type": "string",
                    "example": "1.0.0"
                },
                "wifiConfigs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.ProfileWifiConfigs"
                    }
                }
            }
        },
        "entity.ProfileWifiConfigs": {
            "type": "object",
            "properties": {
                "priority": {
                    "type": "integer",
                    "example": 1
                },
                "profileName": {
                    "type": "string",
                    "example": "My Profile"
                },
                "tenantId": {
                    "type": "string",
                    "example": "abc123"
                }
            }
        },
        "entity.TLSCerts": {
            "type": "object",
            "properties": {
                "issuedCertificate": {
                    "$ref": "#/definitions/entity.CertCreationResult"
                },
                "rootCertificate": {
                    "$ref": "#/definitions/entity.CertCreationResult"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "v1.DeviceCountResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Device"
                    }
                },
                "totalAccount": {
                    "type": "integer"
                }
            }
        },
        "v1.DomainCountResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Domain"
                    }
                },
                "totalAccount": {
                    "type": "integer"
                }
            }
        },
        "v1.ProfileCountResponse": {
            "type": "object",
            "properties": {
                "profile": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Profile"
                    }
                },
                "totalAccount": {
                    "type": "integer"
                }
            }
        },
        "v1.response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "message"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Go Clean Template API",
	Description:      "Using a translation service as an example",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
