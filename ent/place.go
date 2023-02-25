package ent

import "database/sql"

//  place -
type Place struct {
	PlaceId        int             `json:"place-id"`
	Name           string          `json:"name"`
	Description    sql.NullString  `json:"description"`
	Adress         sql.NullString  `json:"adress"`
	Latitude       sql.NullFloat64 `json:"latitude"`
	Longitude      sql.NullFloat64 `json:"longitude"`
	Number         sql.NullString  `json:"number"`
	Mail           sql.NullString  `json:"mail"`
	Url            sql.NullString  `json:"url"`
	NumberOfRating sql.NullInt32   `json:"number-of-rating"`
	MeanRating     sql.NullFloat64 `json:"mean-rating"`
	Photos         []string        `json:"photos"`
	// Shedule        Shedule         `json:"shedule"`
}

type FavoritePlace struct {
	PlaceId int `json:"place-id"`
	UserId  int `json:"user-id"`
}

type PlaceCard struct {
	PlaceId        int             `json:"place-id"`
	Name           string          `json:"name"`
	Latitude       sql.NullFloat64 `json:"latitude"`
	Longitude      sql.NullFloat64 `json:"longitude"`
	Adress         string          `json:"adress"`
	NumberOfRating sql.NullInt32   `json:"number-of-rating"`
	MeanRating     sql.NullFloat64 `json:"mean-rating"`
	Photo          sql.NullString  `json:"photo"`
}
