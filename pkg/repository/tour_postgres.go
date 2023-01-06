package repository

import (
	"fmt"
	"traveland/ent"

	"github.com/jmoiron/sqlx"
)

type TourBD struct {
	db *sqlx.DB
}

func NewTourBD(db *sqlx.DB) *TourBD {
	return &TourBD{
		db: db,
	}
}

func (r TourBD) AddUserTour(fullTour ent.Tour) (bool, error) {
	// add tour
	var tourId int
	query := fmt.Sprintf("INSERT INTO \"%s\" (user_id) values ($1) RETURNING id", tourTable)
	row := r.db.QueryRow(query,fullTour.UserId)
	if err := row.Scan(&tourId); err != nil {
		return false, err
	}
	// add tour_points
	for _, value := range fullTour.Points {
		var id int
		query := fmt.Sprintf("INSERT INTO \"%s\" (tour_id, place_id,start_tour,end_tour) values ($1,$2,$3,$4) RETURNING id", tourPlaceTable)
		row := r.db.QueryRow(query,tourId,value.PlaceId,value.StartTour,value.EndTour)
		if err := row.Scan(&id); err != nil {
			return false, err
		}
	}
	return true, nil
}
