package ent

//  places and houses
type Place struct{
	PlaceId   	int	  	`json:"place_id"`
	Name 	   	string 	`json:"name"`
	Description	string 	`json:"description"`
	Adress 		string  `json:"adress"`
	X	        int		`json:"x"`
	Y	        int		`json:"Y"`
	Type		struct{
		TypeId 	int 	`json:"type_id"`
		Name 	string 	`json:"name"`
	}					`json:"type"`
	
	Number 		string  `json:"number"`
	Pushkin		bool 	`json:"pushkin"`
	RegisterTime string `json:"register_time"`

	HousePrice  string  `json:"house_price"`
	CountRoom   int     `json:"count_room"`
	Square      float32 `json:"square"`
}