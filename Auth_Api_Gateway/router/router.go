package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger)
	// chiRouter.Use(middlewares.RateLimitter)
	// chiRouter.HandleFunc("/fakestoreservice/*",utils.ProxyToService("https://fakestoreapi.in","/fakestoreservice"))
	UserRouter.Register(chiRouter)
	return chiRouter
}
