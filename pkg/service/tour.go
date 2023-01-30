package service

import (
	"traveland/ent"
	"traveland/pkg/repository"
)

type TourService struct {
	repo repository.Tour
}

func NewTourService(repo repository.Tour) *TourService {
	return &TourService{
		repo: repo,
	}
}

func (s TourService) AddUserTour(newTour ent.AddPoints) (int, error) {
	var fullTour ent.Tour
	fullTour.UserId = newTour.UserId
	for index, value := range newTour.Points {
		var newPoint ent.Point
		if index == 0 {
			newPoint.StartTour = true
		}
		if index == (len(newTour.Points) - 1) {
			newPoint.EndTour = true
		}
		newPoint.PlaceId = value

		fullTour.Points = append(fullTour.Points, newPoint)
	}
	return s.repo.AddUserTour(fullTour)
	// return true,nil
}

func (s TourService) GetUserTours(userId int,params *ent.TourQueryParams) (*[]ent.Tour, error){
	return s.repo.GetUserTours(userId,params)
}
func (s TourService) DeleteTour(tourId int,userId int)(bool,error){
	return s.repo.DeleteTour(tourId,userId)
}
func (s TourService) GetAllGuideTours(params *ent.TourQueryParams)(*[]ent.Tour,error){
	return s.repo.GetAllGuideTours(params)
}
func (s TourService) GetTourInfo(tourId int)(*ent.Tour,error){
	return s.repo.GetTourInfo(tourId)
}