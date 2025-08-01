package router

import (
	"Reviews_Service/controller"
	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}
type ReviewRouter struct {
	ReviewController *controller.ReviewController
}

func NewReviewRouter(controller *controller.ReviewController) Router {
	return &ReviewRouter{
		ReviewController: controller,
	}
}

func SetupRouter(ReviewRouter Router) *chi.Mux{
	chiRouter := chi.NewRouter()
	ReviewRouter.Register(chiRouter)
	return  chiRouter
}	


func(u *ReviewRouter) Register(r chi.Router){
	r.Get("/ping",u.ReviewController.Ping)
}