package service

import (
	"traveland/ent"
	"traveland/pkg/repository"
)

type ReviewService struct {
	repo repository.Review
}

func NewReviewService(repo repository.Review)*ReviewService{
	return &ReviewService{
		repo:repo,
	}
}
func (s *ReviewService) AddReview(review ent.Review) (int,error){
	return s.repo.AddReview(review)
}
func (s *ReviewService) DeleteReview(id int,userId int)(bool,error){
	return s.repo.DeleteReview(id,userId)
}
func (s *ReviewService) GetAllReview(params *ent.ReviewQueryParams)(*ent.AllReviewResponce,error){
	return s.repo.GetAllReview(params)
}
func (s *ReviewService) UpdateReview(reviewId int,rating int, reviewText string) (bool,error){
	return s.repo.UpdateReview(reviewId,rating,reviewText)
}