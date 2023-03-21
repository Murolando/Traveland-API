package ent

import (
	"database/sql"
)

type Location struct {
	MinPrice  sql.NullString `json:"min-price"`
	Pushkin   bool           `json:"pushkin"`
	PlaceInfo Place          `json:"place-info"`
	Shedule   []Shedule      `json:"shedule"`
	Types     []int       `json:"type-id"`
}

type LocationCard struct {
	Pushkin       bool           `json:"pushkin"`
	MinPrice      sql.NullString `json:"min-price"`
	PlaceCardInfo PlaceCard      `json:"place-card"`
	Shedule       []Shedule      `json:"shedule"`
}
