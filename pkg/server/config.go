package server

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port           string
	TimeoutS	   int
	WebDir         string
	LogLevel       string
}

func LoadConfig() *Config {
	config := &Config{
		Port:           getEnvWithDefault("TODO_PORT", "7540"),
		TimeoutS:       getEnvIntWithDefault("TODO_TIMEOUT", 30),
		WebDir:         getEnvWithDefault("TODO_WEB_DIR", "./web"),
		LogLevel:       getEnvWithDefault("TODO_LOG_LEVEL", "info"),
	}
	
	return config
}

func getEnvWithDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func getEnvIntWithDefault(key string, defaultValue int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func (c *Config) GetTimeout() time.Duration {
	return time.Duration(c.TimeoutS) * time.Second
}