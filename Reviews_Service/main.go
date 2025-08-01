package main

import (
	"Reviews_Service/app"
	"Reviews_Service/config"
)

func main(){
	config.Load()
	cfg := app.NewConfig(config.GetString("PORT",":3001"))
	app := app.NewApplication(*cfg)
	app.Run()
}