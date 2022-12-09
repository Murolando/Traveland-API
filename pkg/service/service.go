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

type Service struct {
	Authorization
	Place
	Guide
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
