package env

import (
	"os"
	"strconv"
)

func GetInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.Atoi(value); err == nil {
			return result
		}
		return fallback
	}
	return fallback
}

func Get(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if result, err := strconv.ParseBool(value); err == nil {
			return result
		}
	}
	return fallback
}
