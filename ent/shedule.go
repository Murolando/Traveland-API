package ent

import "database/sql"

type Shedule struct {
	PlaceId      int    `json:"place-id"`
	WeekDay      string `json:"week-day"`
	StartWork    sql.NullString `json:"start-work"`
	EndWork      sql.NullString `json:"end-work"`
	StartTimeOut sql.NullString `json:"start-timeout"`
	EndTimeOut   sql.NullString `json:"end-timeout"`
}
