package config

import (
	"log"
	"os"
)

func MustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("environment variable %s not set", key)
	}
	return val
}
