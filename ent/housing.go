package ent

// import "database/sql"

type Housing struct {
	HousePrice  string          `json:"house-price"`
	// CountRoom   sql.NullInt32   `json:"count-room"`
	// Square      sql.NullFloat64 `json:"square"`
	HouseTypeId int             `json:"house-type"`
	PlaceInfo   Place           `json:"place-info"`
	Shedule     []Shedule       `json:"shedule"`
}

type HouseCard struct {
	HousePrice    string    `json:"house-price"`
	HouseTypeId   int       `json:"house-type"`
	PlaceCardInfo PlaceCard `json:"place-card"`
	Shedule       []Shedule   `json:"shedule"`
}
