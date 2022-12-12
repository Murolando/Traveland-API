package repository

import (
	"traveland/ent"

	"github.com/jmoiron/sqlx"
)

type PlaceCRUDPostgres struct {
	db *sqlx.DB
}

func NewPlaceCRUDPostgres(db *sqlx.DB) *PlaceCRUDPostgres {
	return &PlaceCRUDPostgres{
		db: db,
	}
}

func (p PlaceCRUDPostgres) CreatePlace(place ent.Place) (int, error) {
	return 1,nil
}
