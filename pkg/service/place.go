package service

import (
	"traveland/pkg/repository"
)

type PlaceService struct {
	repo repository.Place
}

func NewPlaceService(repo repository.Place) *PlaceService {
	return &PlaceService{
		repo: repo,
	}
}

func (s PlaceService) GetPlaceByID(id int) (interface{}, error){
	return s.repo.GetPlaceByID(id)
}
func (s PlaceService) GetAllPlaces(placeInd int) (interface{}, error){
	return s.repo.GetAllPlaces(placeInd)
}
