package main

import (
	"Auth_Api_Gateway/app"
	"Auth_Api_Gateway/config"
)

func main() {
	config.Load()
	cfg := app.NewConfig()
	app := app.NewApplication(*cfg)
	app.Run()
}
