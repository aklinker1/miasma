// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Miasma",
    "version": "0.0.0"
  },
  "paths": {
    "/api/apps": {
      "get": {
        "summary": "List all the running apps",
        "operationId": "getApps",
        "parameters": [
          {
            "type": "boolean",
            "description": "Whether or not to show hidden apps",
            "name": "hidden",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/App"
              }
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      },
      "post": {
        "summary": "Create and start a new app",
        "operationId": "createApp",
        "parameters": [
          {
            "name": "app",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/AppInput"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      }
    },
    "/api/apps/{appName}": {
      "get": {
        "summary": "Get an app by name",
        "operationId": "getApp",
        "parameters": [
          {
            "$ref": "#/parameters/appName"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "summary": "Stop and delete an app",
        "operationId": "deleteApp",
        "parameters": [
          {
            "$ref": "#/parameters/appName"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      }
    },
    "/api/apps/{appName}/config": {
      "get": {
        "summary": "get an app's current config",
        "operationId": "getAppConfig",
        "parameters": [
          {
            "$ref": "#/parameters/appName"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/AppConfig"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      },
      "put": {
        "summary": "update an app's config",
        "operationId": "updateAppConfig",
        "parameters": [
          {
            "$ref": "#/parameters/appName"
          },
          {
            "name": "newAppConfig",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/AppConfig"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/AppConfig"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      }
    },
    "/api/apps/{appName}/env": {
      "get": {
        "summary": "get an app's environment variables",
        "operationId": "getAppEnv",
        "parameters": [
          {
            "$ref": "#/parameters/appName"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      },
      "put": {
        "summary": "update an app's env",
        "operationId": "updateAppEnv",
        "parameters": [
          {
            "$ref": "#/parameters/appName"
          },
          {
            "name": "newEnv",
            "in": "body",
            "schema": {
              "type": "object"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      }
    },
    "/api/apps/{appName}/start": {
      "post": {
        "summary": "start the app",
        "operationId": "startApp",
        "parameters": [
          {
            "$ref": "#/parameters/appName"
          }
        ],
        "responses": {
          "200": {
            "description": "Started"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      }
    },
    "/api/apps/{appName}/stop": {
      "post": {
        "summary": "stop the app",
        "operationId": "stopApp",
        "parameters": [
          {
            "$ref": "#/parameters/appName"
          }
        ],
        "responses": {
          "200": {
            "description": "Stopped"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      }
    },
    "/api/health": {
      "get": {
        "summary": "Standard health check endpoint that checks all the service's statuses",
        "operationId": "getHealthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "$ref": "#/definitions/Health"
            }
          }
        }
      }
    },
    "/api/plugins": {
      "get": {
        "summary": "List all available plugins and if they are installed",
        "operationId": "listPlugins",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Plugin"
              }
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      }
    },
    "/api/plugins/{pluginName}": {
      "get": {
        "summary": "Get a plugin",
        "operationId": "getPlugin",
        "parameters": [
          {
            "$ref": "#/parameters/pluginName"
          }
        ],
        "responses": {
          "200": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Plugin"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      },
      "post": {
        "summary": "Install (and start) a plugin",
        "operationId": "installPlugin",
        "parameters": [
          {
            "$ref": "#/parameters/pluginName"
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Plugin"
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      },
      "delete": {
        "summary": "Stop and uninstall a plugin",
        "operationId": "uninstallPlugin",
        "parameters": [
          {
            "$ref": "#/parameters/pluginName"
          }
        ],
        "responses": {
          "200": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Plugin"
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      }
    }
  },
  "definitions": {
    "App": {
      "type": "object",
      "required": [
        "name",
        "image",
        "running"
      ],
      "properties": {
        "hidden": {
          "description": "Whether or not the app is hidden during regular requests",
          "type": "boolean"
        },
        "image": {
          "description": "The image the app is based off of",
          "type": "string"
        },
        "name": {
          "description": "The apps name, used in the CLI with the ` + "`" + `-a|--app` + "`" + ` flag",
          "type": "string"
        },
        "running": {
          "type": "boolean"
        }
      }
    },
    "AppConfig": {
      "type": "object",
      "properties": {
        "networks": {
          "description": "A list of other apps that the service communicates with using their service name and docker's internal DNS. Services don't have to be two way; only the service that accesses the other needs the other network added",
          "type": "array",
          "uniqueItems": true,
          "items": {
            "type": "string"
          }
        },
        "placement": {
          "description": "The placement constraints specifying which nodes the app will be ran on. Any valid value for the [` + "`" + `--constraint` + "`" + ` flag](https://docs.docker.com/engine/swarm/services/#placement-constraints) is valid item in this list",
          "type": "array",
          "uniqueItems": true,
          "items": {
            "type": "string"
          }
        },
        "publishedPorts": {
          "description": "The ports that you access the app through in the swarm. This field can, and should be left empty. Miasma automatically manages assigning published ports between 3001-4999. If you need to specify a port, make sure it's outside that range or the port has not been taken. Plugins have set ports starting with 4000, so avoid 4000-4020 if you want to add a plugin at a later date. If these ports are ever cleared, the app will continue using the same ports it was published to before, so that the ports don't change unnecessarily. If you removed it to clear a port for another app/plugin, make sure to restart the app and a new, random port will be allocated for the app, freeing the old port",
          "type": "array",
          "uniqueItems": true,
          "items": {
            "type": "integer"
          }
        },
        "route": {
          "description": "When the Traefik plugin is installed, the route describes where the app can be accessed from.",
          "type": "object",
          "properties": {
            "host": {
              "description": "Describes the hostname the app is served at (\"test.domain.com\")",
              "type": "string",
              "x-nullable": true
            },
            "path": {
              "description": "The path at a given host the app can be reached from (\"/api\"). It should start with a \"/\"",
              "type": "string",
              "x-nullable": true
            },
            "traefikRule": {
              "description": "Instead of using ` + "`" + `host` + "`" + ` and/or ` + "`" + `path` + "`" + `, you can specify the exact rule Traefik will use to route to the app. See [Traefik's documentation]() for how to use this field. This field takes priority over ` + "`" + `host` + "`" + ` and ` + "`" + `path` + "`" + `",
              "type": "string",
              "x-nullable": true
            }
          },
          "x-nullable": true
        },
        "targetPorts": {
          "description": "The ports that the app is listening to inside the container. If no target ports are specified, then the container should respect the ` + "`" + `PORT` + "`" + ` env var.",
          "type": "array",
          "uniqueItems": true,
          "items": {
            "type": "integer"
          }
        }
      }
    },
    "AppInput": {
      "type": "object",
      "required": [
        "name",
        "image"
      ],
      "properties": {
        "hidden": {
          "description": "Whether or not the app is hidden during regular requests",
          "type": "boolean"
        },
        "image": {
          "description": "The image the app is based off of",
          "type": "string"
        },
        "name": {
          "description": "The apps name, used in the CLI with the ` + "`" + `-a|--app` + "`" + ` flag",
          "type": "string"
        }
      }
    },
    "Health": {
      "type": "object",
      "required": [
        "version",
        "dockerVersion",
        "swarm"
      ],
      "properties": {
        "dockerVersion": {
          "description": "The version of docker running on the host, or null if docker is not running",
          "type": "string"
        },
        "swarm": {
          "description": "The info about the docker swarm if the host running miasma is apart of one. If it is not apart of a swarm, it returns ` + "`" + `null` + "`" + `",
          "type": "object",
          "properties": {
            "createdAt": {
              "description": "UTC timestamps when the swarm was created",
              "type": "string"
            },
            "id": {
              "description": "The swarm's ID",
              "type": "string"
            },
            "joinCommand": {
              "description": "The command for a node to run to join the swarm",
              "type": "string"
            },
            "updatedAt": {
              "description": "UTC timestamps when the swarm was last updated",
              "type": "string"
            }
          }
        },
        "version": {
          "description": "Miasma's current version",
          "type": "string"
        }
      }
    },
    "Plugin": {
      "type": "object",
      "required": [
        "name",
        "installed"
      ],
      "properties": {
        "installCommand": {
          "description": "Command to run to install the plugin",
          "type": "string",
          "x-nullable": true
        },
        "installed": {
          "description": "Whether or not the plugin is installed",
          "type": "boolean"
        },
        "name": {
          "description": "The plugin's name. It can be used to install a plugin",
          "type": "string"
        },
        "uninstallCommand": {
          "description": "Command to run to uninstall the plugin",
          "type": "string",
          "x-nullable": true
        }
      }
    }
  },
  "parameters": {
    "appName": {
      "type": "string",
      "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
      "name": "appName",
      "in": "path",
      "required": true
    },
    "pluginName": {
      "type": "string",
      "name": "pluginName",
      "in": "path",
      "required": true
    }
  },
  "responses": {
    "unknown": {
      "description": "Unknown Error",
      "schema": {
        "type": "string"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Miasma",
    "version": "0.0.0"
  },
  "paths": {
    "/api/apps": {
      "get": {
        "summary": "List all the running apps",
        "operationId": "getApps",
        "parameters": [
          {
            "type": "boolean",
            "description": "Whether or not to show hidden apps",
            "name": "hidden",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/App"
              }
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "summary": "Create and start a new app",
        "operationId": "createApp",
        "parameters": [
          {
            "name": "app",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/AppInput"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/apps/{appName}": {
      "get": {
        "summary": "Get an app by name",
        "operationId": "getApp",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "summary": "Stop and delete an app",
        "operationId": "deleteApp",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/apps/{appName}/config": {
      "get": {
        "summary": "get an app's current config",
        "operationId": "getAppConfig",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/AppConfig"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "put": {
        "summary": "update an app's config",
        "operationId": "updateAppConfig",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          },
          {
            "name": "newAppConfig",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/AppConfig"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/AppConfig"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/apps/{appName}/env": {
      "get": {
        "summary": "get an app's environment variables",
        "operationId": "getAppEnv",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "put": {
        "summary": "update an app's env",
        "operationId": "updateAppEnv",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          },
          {
            "name": "newEnv",
            "in": "body",
            "schema": {
              "type": "object"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/apps/{appName}/start": {
      "post": {
        "summary": "start the app",
        "operationId": "startApp",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Started"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/apps/{appName}/stop": {
      "post": {
        "summary": "stop the app",
        "operationId": "stopApp",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Stopped"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/health": {
      "get": {
        "summary": "Standard health check endpoint that checks all the service's statuses",
        "operationId": "getHealthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "$ref": "#/definitions/Health"
            }
          }
        }
      }
    },
    "/api/plugins": {
      "get": {
        "summary": "List all available plugins and if they are installed",
        "operationId": "listPlugins",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Plugin"
              }
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/plugins/{pluginName}": {
      "get": {
        "summary": "Get a plugin",
        "operationId": "getPlugin",
        "parameters": [
          {
            "type": "string",
            "name": "pluginName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Plugin"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "summary": "Install (and start) a plugin",
        "operationId": "installPlugin",
        "parameters": [
          {
            "type": "string",
            "name": "pluginName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Plugin"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "summary": "Stop and uninstall a plugin",
        "operationId": "uninstallPlugin",
        "parameters": [
          {
            "type": "string",
            "name": "pluginName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Plugin"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "App": {
      "type": "object",
      "required": [
        "name",
        "image",
        "running"
      ],
      "properties": {
        "hidden": {
          "description": "Whether or not the app is hidden during regular requests",
          "type": "boolean"
        },
        "image": {
          "description": "The image the app is based off of",
          "type": "string"
        },
        "name": {
          "description": "The apps name, used in the CLI with the ` + "`" + `-a|--app` + "`" + ` flag",
          "type": "string"
        },
        "running": {
          "type": "boolean"
        }
      }
    },
    "AppConfig": {
      "type": "object",
      "properties": {
        "networks": {
          "description": "A list of other apps that the service communicates with using their service name and docker's internal DNS. Services don't have to be two way; only the service that accesses the other needs the other network added",
          "type": "array",
          "uniqueItems": true,
          "items": {
            "type": "string"
          }
        },
        "placement": {
          "description": "The placement constraints specifying which nodes the app will be ran on. Any valid value for the [` + "`" + `--constraint` + "`" + ` flag](https://docs.docker.com/engine/swarm/services/#placement-constraints) is valid item in this list",
          "type": "array",
          "uniqueItems": true,
          "items": {
            "type": "string"
          }
        },
        "publishedPorts": {
          "description": "The ports that you access the app through in the swarm. This field can, and should be left empty. Miasma automatically manages assigning published ports between 3001-4999. If you need to specify a port, make sure it's outside that range or the port has not been taken. Plugins have set ports starting with 4000, so avoid 4000-4020 if you want to add a plugin at a later date. If these ports are ever cleared, the app will continue using the same ports it was published to before, so that the ports don't change unnecessarily. If you removed it to clear a port for another app/plugin, make sure to restart the app and a new, random port will be allocated for the app, freeing the old port",
          "type": "array",
          "uniqueItems": true,
          "items": {
            "type": "integer"
          }
        },
        "route": {
          "description": "When the Traefik plugin is installed, the route describes where the app can be accessed from.",
          "type": "object",
          "properties": {
            "host": {
              "description": "Describes the hostname the app is served at (\"test.domain.com\")",
              "type": "string",
              "x-nullable": true
            },
            "path": {
              "description": "The path at a given host the app can be reached from (\"/api\"). It should start with a \"/\"",
              "type": "string",
              "x-nullable": true
            },
            "traefikRule": {
              "description": "Instead of using ` + "`" + `host` + "`" + ` and/or ` + "`" + `path` + "`" + `, you can specify the exact rule Traefik will use to route to the app. See [Traefik's documentation]() for how to use this field. This field takes priority over ` + "`" + `host` + "`" + ` and ` + "`" + `path` + "`" + `",
              "type": "string",
              "x-nullable": true
            }
          },
          "x-nullable": true
        },
        "targetPorts": {
          "description": "The ports that the app is listening to inside the container. If no target ports are specified, then the container should respect the ` + "`" + `PORT` + "`" + ` env var.",
          "type": "array",
          "uniqueItems": true,
          "items": {
            "type": "integer"
          }
        }
      }
    },
    "AppConfigRoute": {
      "description": "When the Traefik plugin is installed, the route describes where the app can be accessed from.",
      "type": "object",
      "properties": {
        "host": {
          "description": "Describes the hostname the app is served at (\"test.domain.com\")",
          "type": "string",
          "x-nullable": true
        },
        "path": {
          "description": "The path at a given host the app can be reached from (\"/api\"). It should start with a \"/\"",
          "type": "string",
          "x-nullable": true
        },
        "traefikRule": {
          "description": "Instead of using ` + "`" + `host` + "`" + ` and/or ` + "`" + `path` + "`" + `, you can specify the exact rule Traefik will use to route to the app. See [Traefik's documentation]() for how to use this field. This field takes priority over ` + "`" + `host` + "`" + ` and ` + "`" + `path` + "`" + `",
          "type": "string",
          "x-nullable": true
        }
      },
      "x-nullable": true
    },
    "AppInput": {
      "type": "object",
      "required": [
        "name",
        "image"
      ],
      "properties": {
        "hidden": {
          "description": "Whether or not the app is hidden during regular requests",
          "type": "boolean"
        },
        "image": {
          "description": "The image the app is based off of",
          "type": "string"
        },
        "name": {
          "description": "The apps name, used in the CLI with the ` + "`" + `-a|--app` + "`" + ` flag",
          "type": "string"
        }
      }
    },
    "Health": {
      "type": "object",
      "required": [
        "version",
        "dockerVersion",
        "swarm"
      ],
      "properties": {
        "dockerVersion": {
          "description": "The version of docker running on the host, or null if docker is not running",
          "type": "string"
        },
        "swarm": {
          "description": "The info about the docker swarm if the host running miasma is apart of one. If it is not apart of a swarm, it returns ` + "`" + `null` + "`" + `",
          "type": "object",
          "properties": {
            "createdAt": {
              "description": "UTC timestamps when the swarm was created",
              "type": "string"
            },
            "id": {
              "description": "The swarm's ID",
              "type": "string"
            },
            "joinCommand": {
              "description": "The command for a node to run to join the swarm",
              "type": "string"
            },
            "updatedAt": {
              "description": "UTC timestamps when the swarm was last updated",
              "type": "string"
            }
          }
        },
        "version": {
          "description": "Miasma's current version",
          "type": "string"
        }
      }
    },
    "HealthSwarm": {
      "description": "The info about the docker swarm if the host running miasma is apart of one. If it is not apart of a swarm, it returns ` + "`" + `null` + "`" + `",
      "type": "object",
      "properties": {
        "createdAt": {
          "description": "UTC timestamps when the swarm was created",
          "type": "string"
        },
        "id": {
          "description": "The swarm's ID",
          "type": "string"
        },
        "joinCommand": {
          "description": "The command for a node to run to join the swarm",
          "type": "string"
        },
        "updatedAt": {
          "description": "UTC timestamps when the swarm was last updated",
          "type": "string"
        }
      }
    },
    "Plugin": {
      "type": "object",
      "required": [
        "name",
        "installed"
      ],
      "properties": {
        "installCommand": {
          "description": "Command to run to install the plugin",
          "type": "string",
          "x-nullable": true
        },
        "installed": {
          "description": "Whether or not the plugin is installed",
          "type": "boolean"
        },
        "name": {
          "description": "The plugin's name. It can be used to install a plugin",
          "type": "string"
        },
        "uninstallCommand": {
          "description": "Command to run to uninstall the plugin",
          "type": "string",
          "x-nullable": true
        }
      }
    }
  },
  "parameters": {
    "appName": {
      "type": "string",
      "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
      "name": "appName",
      "in": "path",
      "required": true
    },
    "pluginName": {
      "type": "string",
      "name": "pluginName",
      "in": "path",
      "required": true
    }
  },
  "responses": {
    "unknown": {
      "description": "Unknown Error",
      "schema": {
        "type": "string"
      }
    }
  }
}`))
}
