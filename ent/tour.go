package ent

import "database/sql"

type Tour struct {
	TourId      int            `json:"tour-id"`
	UserId      int            `json:"user-id"`
	Name        sql.NullString `json:"name"`
	Price		sql.NullInt32  `json:"price"`
	Description sql.NullString `json:"description"`
	Points      []Point        `json:"points"`
}

type Point struct {
	PlaceId   int  `json:"place-id"`
	StartTour bool `json:"start"`
	EndTour   bool `json:"end"`
}

type AddPoints struct{
	UserId   int  `json:"user-id"`
	Points   []int  `json:"points"`
}

