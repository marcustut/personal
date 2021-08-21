package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/marcustut/personal/internal/app"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env: %s\n", err)
	}
}

func main() {
	app := app.New()
	app.Run(8080)
}
