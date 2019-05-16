package helpers

import (
	"github.com/labstack/gommon/log"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvWithPanic(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Panic("not set " + key)
	return ""
}
