package config

import (
	"os"
)

var (
	apiUrl           = "postgres://root:postgres@postgres:5432/shoppingdb?sslmode=disable"
	port             = "8181"
	databaseHost     = "localhost"
	databasePort     = "5432"
	databaseName     = "shoppingdb"
	databaseUser     = "root"
	databasePassword = "secret"
	databaseSSLMode  = "disable"
	databaseVendor   = "postgres"
)

type Config struct {
	Port string
	// Database
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseSSLMode  string
	DatabaseVendor   string
}

func NewConfig() *Config {
	return &Config{
		// App mode

		Port: getEnvStr("LISTEN_PORT", port),

		// Database
		DatabaseHost:     getEnvStr("DB_HOST", databaseHost),
		DatabasePort:     getEnvStr("DATABASE_PORT", databasePort),
		DatabaseName:     getEnvStr("DB_NAME", databaseName),
		DatabaseUser:     getEnvStr("DB_USERNAME", databaseUser),
		DatabasePassword: getEnvStr("DB_PASSWORD", databasePassword),
		DatabaseSSLMode:  getEnvStr("DB_SSL_MODE", databaseSSLMode),
		DatabaseVendor:   getEnvStr("DATABASE_VENDOR", databaseVendor),
	}
}
func getEnvStr(variable string, defaultValue string) string {
	if env, ok := os.LookupEnv(variable); ok {
		return env
	}
	return defaultValue
}
