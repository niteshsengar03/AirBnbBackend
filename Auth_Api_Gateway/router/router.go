package router

import (
	"Auth_Api_Gateway/controller"
	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	R := chi.NewRouter()
	R.Get("/", controller.Ping)
	return R
}
