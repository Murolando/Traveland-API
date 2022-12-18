package ent

import "database/sql"

type Location struct {
	MinPrice  sql.NullString `json:"min-price"`
	Pushkin   bool           `json:"pushkin"`
	PlaceInfo Place          `json:"place-info"`
}
