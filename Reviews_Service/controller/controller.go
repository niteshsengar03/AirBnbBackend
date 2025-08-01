package controller

import (
	"Reviews_Service/service"
	"encoding/json"
	"net/http"
)

type ReviewController struct {
	ReviewService service.ReviewService
}

func NewReviewController (service service.ReviewService) *ReviewController{
		return &ReviewController{
			ReviewService:  service,
		}
}
func(u *ReviewController) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set header
	w.WriteHeader(200)
	response := map[string]string{"message": "pong"} // Or use a struct
	json.NewEncoder(w).Encode(response)              // Encode to JSON
}
