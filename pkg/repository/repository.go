package repository

import (
	"traveland/ent"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user ent.User) (int, error)
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
