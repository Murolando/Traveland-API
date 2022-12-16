package ent

import "database/sql"

//  place -
type Place struct{
	PlaceId   	int	  			`json:"place_id"`
	Name 	   	string 			`json:"name"`
	Description	string 			`json:"description"`
	Adress 		string  		`json:"adress"`
	Latitude	sql.NullFloat64	`json:"latitude"`
	Longitude   sql.NullFloat64	`json:"longitude"`
	Number 		sql.NullString  `json:"number"`
	PlaceType   []int 			`json:"place-type"`
}

