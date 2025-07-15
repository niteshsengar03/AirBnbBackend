package main

import (
	"Auth_Api_Gateway/app"
	"Auth_Api_Gateway/config"
)

func main() {
	config.Load()
	cfg := app.Config{
		Addr: config.GetString("PORT", ":3002"),
	}
	app := app.Application{
		Config: cfg,
	}	
	app.Run()
}
