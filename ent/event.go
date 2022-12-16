package ent

type Event struct{
	Price 		string 	`json:"price"`
	Pushkin	 	bool 	`json:"pushkin"`
	PlaceInfo   Place   `json:"place-info"`
}

