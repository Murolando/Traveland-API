package ent

import "database/sql"

type Event struct {
	Price     sql.NullString `json:"price"`
	Pushkin   bool           `json:"pushkin"`
	PlaceInfo Place          `json:"place-info"`
}
