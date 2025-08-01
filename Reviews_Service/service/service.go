package service

import "Reviews_Service/db"

type ReviewService interface {

}

type ReviewServiceImp struct {
	ReviewRepository db.ReviewRepository 
}

func NewReviewService(_rp db.ReviewRepository) ReviewService{
	return &ReviewServiceImp{
		ReviewRepository :_rp,
	}
}

