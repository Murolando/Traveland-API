package service

import (
	"traveland/ent"
	"traveland/pkg/repository"
)

type Authorization interface {
	CreateUser(ent.User) (int, error)
	GenerateToken(mail string, password string) (string,error)
	ParseToken(token string) (int,error)
}

type Place interface {
	GetPlaceByID(id int) (interface{}, error)
	GetAllPlaces(placeInd int) (interface{}, error)
}
type User interface{
	GetUserByID(id int) (ent.User,error)
	GetAllUsers() ([]ent.User,error)
	GetUsersByRole(role_id int) ([]ent.User,error)
}

type Service struct {
	Authorization
	Place
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		User: NewUserCRUDService(repo.User),
		Place: NewPlaceService(repo.Place),
	}
}
