package ent

type Review struct{
	ReviewId 	int 	`json:"rewiew-id"`
	UserId   	int		`json:"user-id"`
	Rating		int		`json:"rating"`
	ReviewText 	string	`json:"rewiew-text"`
	ReviewTime	string	`json:"rewiew-time"`
	PlaceId		int		`json:"place-id"`
	GuideId		int		`json:"guide-id"`
}