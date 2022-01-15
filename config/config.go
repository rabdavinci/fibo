package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	LogLevel string

	HTTP_HOST string
	HTTP_PORT string

	GRPC_HOST string
	GRPC_PORT string

	REDIS_HOST string
	REDIS_PORT string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.HTTP_HOST = cast.ToString(getOrReturnDefault("HTTP_HOST", "localhost"))
	c.HTTP_PORT = cast.ToString(getOrReturnDefault("HTTP_PORT", 9090))

	c.GRPC_HOST = cast.ToString(getOrReturnDefault("GRPC_HOST", "localhost"))
	c.GRPC_PORT = cast.ToString(getOrReturnDefault("GRPC_PORT", 8080))

	c.REDIS_HOST = cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost"))
	c.REDIS_PORT = cast.ToString(getOrReturnDefault("REDIS_PORT", 6379))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
