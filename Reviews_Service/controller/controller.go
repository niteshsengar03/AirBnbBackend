package controller

import (
	"Reviews_Service/service"
)

type ReviewController struct {
	ReviewService service.ReviewService
}

func NewReviewController (service service.ReviewService) *ReviewController{
		return &ReviewController{
			ReviewService:  service,
		}
}