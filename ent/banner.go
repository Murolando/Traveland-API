package ent

import "database/sql"

type Banner struct {
	PlaceId     int            `json:"place-id"`
	// Name        string         `json:"name"`
	// OrderNumber int            `json:"order"`
	Image_src   sql.NullString `json:"image"`
	
}
