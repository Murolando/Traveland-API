package ent

type Location struct{
	MinPrice 	string `json:"min-price"`
	Pushkin	 	bool 	`json:"pushkin"`
	PlaceInfo	Place  `json:"place-info"`
}