package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/marcustut/personal/internal/app"
	"github.com/marcustut/personal/internal/config"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env: %s\n", err)
	}
}

func main() {
	a := app.New()

	ap := app.AppParams{Port: config.Port(), DbUrl: config.DatabaseUrl()}

	a.Run(ap)
}
