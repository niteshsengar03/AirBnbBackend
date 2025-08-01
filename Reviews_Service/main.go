package main

import (
	"Reviews_Service/app"
	"Reviews_Service/config"
	"log"
)

func main() {
	config.Load()
	app := app.NewApplication(config.GetString("PORT", ":3001"))
	err:=app.Run()
	if err != nil {
		log.Fatalf("Application exited with error: %v", err)
	}
}
