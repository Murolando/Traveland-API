package repository

import "traveland/ent"

type Authorization interface {
	CreateUser(ent.User)(int, error)
	LoginUser(ent.User)(bool,error)
}

type Place interface {
	
}

type Guide interface {
}

type Repository struct {
	Authorization
	Place
	Guide
}

func NewRepository() *Repository {
	return &Repository{}
}
