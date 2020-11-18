package utils

import (
	"log"
	"os"
)

// GetEnvKey : return Environment value
func GetEnvKey(env string) string {
	key, ok := os.LookupEnv(env)
	if !ok {
		log.Println("App : " + env + " environment variable required but not set")
		os.Exit(1)
	}
	return key
}
