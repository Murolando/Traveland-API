package ent

import "database/sql"

type Review struct {
	ReviewId   int    `json:"rewiew-id"`
	UserId     int    `json:"user-id"`
	Rating     int    `json:"rating"`
	ReviewText string `json:"review-text"`
	ReviewTime string `json:"review-time"`
	PlaceId    int    `json:"place-id"`
	GuideId    int    `json:"guide-id"`
}

type ReviewResponce struct {
	ReviewId   int    `json:"rewiew-id"`
	UserName   string `json:"user-name"`
	Rating     int    `json:"rating"`
	ReviewText string `json:"review-text"`
	ReviewTime string `json:"review-time"`
}
type AllReviewResponce struct {
	Reviews        []ReviewResponce `json:"reviews"`
	MeanRating     sql.NullFloat64   `json:"mean-rating"`
	// NumberOfRating sql.NullInt32     `json:"number-of-rating"`
}
