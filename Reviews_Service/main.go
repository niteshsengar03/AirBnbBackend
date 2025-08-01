package main

import (
	"Reviews_Service/app"
	"Reviews_Service/config"
)

func main() {
	config.Load()
	app := app.NewApplication(config.GetString("PORT", ":3001"))
	app.Run()
}
