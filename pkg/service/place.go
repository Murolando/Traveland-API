package service

import (
	"traveland/ent"
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

func (s PlaceService) GetPlaceByID(id int) (interface{}, error) {
	return s.repo.GetPlaceByID(id)
}
func (s PlaceService) GetAllPlaces(placeInd int, offset int) (interface{}, error) {
	return s.repo.GetAllPlaces(placeInd, offset)
}
func (s PlaceService) GetLocalByType(placeType int, offset int) (*[]ent.Location, error){
	return s.repo.GetLocalByType(placeType,offset)
}
func (s PlaceService)  GetHouseByType(houseType int,offset int) (*[]ent.Housing, error){
	return s.repo.GetHouseByType(houseType,offset)
}
func (s PlaceService) GetLocalTypes() (*[]ent.LocalType,error){
	return s.repo.GetLocalTypes()
}
func (s PlaceService) GetHouseTypes() (*[]ent.HouseType,error){
	return s.repo.GetHouseTypes()
}
func (s PlaceService) AddFavoritePlace(userId int, placeId int) (bool, error){
	return s.repo.AddFavoritePlace(userId,placeId)
}
func (s PlaceService) GetAllUserFavoritePlaces(userId int) (*[]interface{}, error){
	return s.repo.GetAllUserFavoritePlaces(userId)
}
func (s PlaceService) GetCountOfPlaceFavorites(placeId int) (int, error){
	return s.repo.GetCountOfPlaceFavorites(placeId)
}