package config

import "os"

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// JWTSecret JWT密钥
func JWTSecret() string {
	return GetEnv("JWT_SECRET", "your-secret-key")
}
