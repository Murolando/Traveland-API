package service

import (
	"traveland/ent"
	"traveland/pkg/repository"
)

type PlaceCRUDService struct {
	repo repository.Place
}

func NewPlaceCRUDService(repo repository.Place) *PlaceCRUDService {
	return &PlaceCRUDService{
		repo: repo,
	}
}

func (s PlaceCRUDService) CreatePlace(place ent.Place) (int, error) {
	return s.repo.CreatePlace(place)
}
