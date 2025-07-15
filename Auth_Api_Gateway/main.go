package main

import (
	"Auth_Api_Gateway/app"
	"Auth_Api_Gateway/config"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	config.Load()
	cfg := app.Config{
		Addr: config.GetString("PORT", ":3002"),
	}
	app := app.Application{
		Config: cfg,
	}

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	app.Run(r)
}
