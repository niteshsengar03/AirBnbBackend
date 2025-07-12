package main

import (
	"Auth_Api_Gateway/app"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := app.Config{
		Addr: os.Getenv("PORT"),
	}
	app := app.Application{
		Config: cfg,
	}
	app.Run()
}
