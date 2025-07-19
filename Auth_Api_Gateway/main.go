package main

import (
	"Auth_Api_Gateway/app"
	"Auth_Api_Gateway/config"
	// dbConfig "Auth_Api_Gateway/config/db"
)

func main() {
	config.Load()
	cfg := app.NewConfig()
	app := app.NewApplication(*cfg)
	// dbConfig.SetupDB();
	app.Run()
}
