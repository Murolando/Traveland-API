package repository

import (
	"traveland/ent"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user ent.User) (int, error)
	GetUserByMailAndPassword(mail string , password string)(int, error)
	
}

type Place interface {
	GetPlaceByID(id int) (interface{}, error)
	GetAllPlaces(placeInd int) (interface{}, error)
}
type User interface{
	GetUserByID(id int) (ent.User, error)
	GetAllUsers()([]ent.User,error)
	GetUsersByRole(role_id int) ([]ent.User,error)

}
type Repository struct {
	Authorization
	Place
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User: NewUserBD(db),
		Place: NewPlaceBD(db),
	}
}
