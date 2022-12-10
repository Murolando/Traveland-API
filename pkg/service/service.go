package service

import (
	"traveland/ent"
	"traveland/pkg/repository"
)

type Authorization interface {
	CreateUser(ent.User) (int, error)
}

type Place interface {
}

type Guide interface {
}

type User interface{
	GetUserByID(id int) (ent.User,error)
	GetAllUsers() ([]ent.User,error)
	GetUsersByRole(role_id int) ([]ent.User,error)
}

type Service struct {
	Authorization
	Place
	Guide
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		User: NewUserCRUDService(repo.User),
	}
}
