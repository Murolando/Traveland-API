package service

import "traveland/pkg/repository"

type Authorization interface {
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

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}