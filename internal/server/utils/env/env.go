package env

import "os"

var IS_PROD = os.Getenv("MODE") == "production"

var PLUGIN_POSTGRES_DATA_VOLUME = os.Getenv("PLUGIN_POSTGRES_DATA_VOLUME")
var PLUGIN_POSTGRES_PASSWORD = os.Getenv("PLUGIN_POSTGRES_PASSWORD")
var PLUGIN_POSTGRES_USER = os.Getenv("PLUGIN_POSTGRES_USER")

var PLUGIN_MONGO_DATA_VOLUME = os.Getenv("PLUGIN_POSTGRES_DATA_VOLUME")

func init() {
	// Postgres
	if PLUGIN_POSTGRES_DATA_VOLUME == "" {
		PLUGIN_POSTGRES_DATA_VOLUME = "/data/postgres"
	}
	if PLUGIN_POSTGRES_PASSWORD == "" {
		PLUGIN_POSTGRES_PASSWORD = "password"
	}
	if PLUGIN_POSTGRES_USER == "" {
		PLUGIN_POSTGRES_USER = "postgres"
	}

	// Mongo
	if PLUGIN_MONGO_DATA_VOLUME == "" {
		PLUGIN_MONGO_DATA_VOLUME = "/data/mongo"
	}
}
