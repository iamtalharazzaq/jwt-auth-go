package config

import (
	"os"
	"strconv"
	"strings"
)

// Config holds all environment-based configuration
var JWTSecret = getEnv("JWT_SECRET", "secret_key")
var TokenExpiryMinutes = getEnvAsInt("TOKEN_EXPIRY_MINUTES", 5)
var Users = parseUsers(mustGetEnv("USERS"))

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvAsInt(name string, defaultVal int) int {
	if valStr := os.Getenv(name); valStr != "" {
		if val, err := strconv.Atoi(valStr); err == nil {
			return val
		}
	}
	return defaultVal
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Environment variable required but not set: " + key)
	}
	return value
}

// parseUsers parses a comma-separated list of user:password pairs
func parseUsers(s string) map[string]string {
	users := make(map[string]string)
	if strings.TrimSpace(s) == "" {
		panic("USERS environment variable is empty or only whitespace")
	}
	parts := strings.Split(s, ",")
	for _, part := range parts {
		kv := strings.SplitN(strings.TrimSpace(part), ":", 2)
		if len(kv) != 2 || kv[0] == "" || kv[1] == "" {
			panic("Invalid USERS entry: '" + part + "'. Expected format username:password")
		}
		users[kv[0]] = kv[1]
	}
	if len(users) == 0 {
		panic("No valid users found in USERS environment variable")
	}
	return users
}
