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
type User interface{
	GetUserByID(id int) (ent.User, error)
	GetAllUsers()([]ent.User,error)
	GetUsersByRole(role_id int) ([]ent.User,error)

}
type Repository struct {
	Authorization
	Place
	Guide
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User: NewUserCRUD(db),
	}
}
