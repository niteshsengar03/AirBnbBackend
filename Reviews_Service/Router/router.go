package router

import (
	"Reviews_Service/controller"
)

type Router interface {
}
type ReviewRouter struct {
	ReviewController *controller.ReviewController
}

func NewReviewRouter(controller *controller.ReviewController) Router {
	return &ReviewRouter{
		ReviewController: controller,
	}
}

func SetupRouter (){
	
}
