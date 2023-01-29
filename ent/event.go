package ent

import "database/sql"

type Event struct {
	Price          sql.NullString `json:"price"`
	Pushkin        bool           `json:"pushkin"`
	PlaceInfo      Place          `json:"place-info"`
	EventDay       sql.NullString `json:"event-day"`
	EventStartTime sql.NullString `json:"event-start-time"`
	EventEndTime   sql.NullString `json:"event-end-time"`
}

type EventCard struct {
	Price          sql.NullString `json:"price"`
	EventDay       sql.NullString `json:"event-day"`
	EventStartTime sql.NullString `json:"event-start-time"`
	EventEndTime   sql.NullString `json:"event-end-time"`
	Pushkin        bool           `json:"pushkin"`
	PlaceCardInfo  PlaceCard      `json:"place-card"`
}
