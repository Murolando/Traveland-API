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

func (s TourService) AddUserTour(newTour ent.AddPoints) (bool, error) {
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
