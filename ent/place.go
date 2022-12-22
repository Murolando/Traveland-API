package ent

import "database/sql"

//  place -
type Place struct{
	PlaceId   		int	  			`json:"place-id"`
	Name 	   		string 			`json:"name"`
	Description		string 			`json:"description"`
	Adress 			string  		`json:"adress"`
	Latitude		sql.NullFloat64	`json:"latitude"`
	Longitude   	sql.NullFloat64	`json:"longitude"`
	Number 			sql.NullString  `json:"number"`
	NumberOfRating	sql.NullFloat64 `json:"number-of-rating"`
	MeanRating      sql.NullFloat64 `json:"mean-rating"`
}

