package helpers

import (
	"github.com/labstack/gommon/log"
	"os"
)

// GetEnv returns env value or return fallback value
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetEnvWithPanic returns env value or Panic
func GetEnvWithPanic(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Panic("not set " + key)
	return ""
}
