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

func (r TourBD) AddUserTour(fullTour ent.Tour) (int, error) {
	// check tour limits
	var count int = 0
	query := fmt.Sprintf("SELECT COUNT(id) FROM \"%s\" WHERE user_id = $1", tourTable)
	row := r.db.QueryRow(query,fullTour.UserId)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	if (count == 20){
		return -1,nil
	}
	// add tour
	var tourId int
	query = fmt.Sprintf("INSERT INTO \"%s\" (user_id) values ($1) RETURNING id", tourTable)
	row = r.db.QueryRow(query,fullTour.UserId)
	if err := row.Scan(&tourId); err != nil {
		return 0, err
	}
	// add tour_points
	for _, value := range fullTour.Points {
		var id int
		query := fmt.Sprintf("INSERT INTO \"%s\" (tour_id, place_id,start_tour,end_tour) values ($1,$2,$3,$4) RETURNING id", tourPlaceTable)
		row := r.db.QueryRow(query,tourId,value.PlaceId,value.StartTour,value.EndTour)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
	}
	return tourId, nil
}

func (r TourBD) GetUserTours(userId int,offset int) (*[]ent.Tour, error) {
	// get tours 
	var tours []ent.Tour = make([]ent.Tour, 0)
	query := fmt.Sprintf("SELECT (id) FROM \"%s\" WHERE user_id = $1 LIMIT $2 OFFSET $3", tourTable)
	rows, err := r.db.Query(query, userId, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		var tour ent.Tour
		if err := rows.Scan(&tour.TourId); err != nil {
			return nil, err
		}
		tour.UserId = userId

		query := fmt.Sprintf("SELECT place_id,start_tour,end_tour FROM \"%s\" WHERE tour_id = $1 LIMIT $2 OFFSET $3", tourPlaceTable)
		
		pointRows, err := r.db.Query(query, tour.TourId, limit, offset)
		if err != nil {
			return nil, err
		}
		var points []ent.Point = make([]ent.Point, 0)
		for pointRows.Next(){
			var point ent.Point
			if err := pointRows.Scan(&point.PlaceId,&point.StartTour,&point.EndTour); err != nil {
				return nil, err
			}
			points = append(points,point)
		}
		tour.Points = points

		tours = append(tours, tour)
	}
	return  &tours,nil
}

func (r TourBD) DeleteTour(tourId int)(bool,error){
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`,tourTable)
	res1,err := r.db.Exec(query,tourId)
	if err!=nil{
		return false,err
	}
    count, err := res1.RowsAffected()
    if err != nil {
        return false,err
    }
	if count != 0{
		return true,nil
	}
	return false,nil
}

func (r TourBD) GetAllGuideTours(offset int)(*[]ent.Tour,error){
	var tours []ent.Tour = make([]ent.Tour, 0)
	query := fmt.Sprintf(`SELECT id,name,description,user_id FROM %s WHERE name IS NOT NULL LIMIT $1 OFFSET $2`, tourTable)
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		var tour ent.Tour
		if err := rows.Scan(&tour.TourId,&tour.Name,&tour.Description,&tour.UserId); err != nil {
			return nil, err
		}
		query := fmt.Sprintf("SELECT place_id,start_tour,end_tour FROM \"%s\" WHERE tour_id = $1", tourPlaceTable)
		
		pointRows, err := r.db.Query(query, tour.TourId)
		if err != nil {
			return nil, err
		}
		var points []ent.Point = make([]ent.Point, 0)
		for pointRows.Next(){
			var point ent.Point
			if err := pointRows.Scan(&point.PlaceId,&point.StartTour,&point.EndTour); err != nil {
				return nil, err
			}
			points = append(points,point)
		}
		tour.Points = points

		tours = append(tours, tour)
	}
	return  &tours,nil
}

func (r TourBD) GetTourInfo(tourId int)(*ent.Tour,error){
	var tour ent.Tour
	query := fmt.Sprintf(`SELECT name,description,user_id FROM %s WHERE id = $1`, tourTable)
	row := r.db.QueryRow(query,tourId)
	if err := row.Scan(&tour.Name,&tour.Description,&tour.UserId); err != nil {
		return nil, err
	}
	tour.TourId = tourId
	query = fmt.Sprintf("SELECT place_id,start_tour,end_tour FROM \"%s\" WHERE tour_id = $1", tourPlaceTable)
	pointRows, err := r.db.Query(query,tourId)
	if err != nil {
		return nil, err
	}
	var points []ent.Point = make([]ent.Point, 0)
	for pointRows.Next(){
		var point ent.Point
		if err := pointRows.Scan(&point.PlaceId,&point.StartTour,&point.EndTour); err != nil {
			return nil, err
		}
		points = append(points,point)
	}
	return  &tour,nil
}