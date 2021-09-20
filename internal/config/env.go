package config

import (
	"log"
	"os"
	"strconv"
)

func getEnv(key string) string {
	value, found := os.LookupEnv(key)

	if found {
		return value
	} else {
		log.Fatalf("%s is not found in the environment", key)
		return ""
	}
}

func DatabaseUrl() string {
	return getEnv("DATABASE_URL")
}

func Port() uint {
	env := getEnv("PORT")
	port, err := strconv.ParseUint(env, 10, 32)

	if err != nil {
		log.Fatalf("%s is not a valid port number", env)
	}

	return uint(port)
}
