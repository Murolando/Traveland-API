package ent

type Shedule struct {
	PlaceId int    `json:"place-id"`
	WeekDay string `json:"week-day"`
	StartWork string `json:"start-work"`
	EndWork string `json:"end-work"`
	StartTimeOut string `json:"start-timeout"`
	EndTimeOut string `json:"end-timeout"`
}
