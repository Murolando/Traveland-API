package ent

type Guide struct{
	UserId   	   	int	  	`json:"user_id"`
	Name 	   	string 	`json:"name"`
	SecondName 	string 	`json:"second_name"`
	Number 		string  `json:"number"`
	Mail        string	`json:"mail"`
	Password	string 	`json:"password"`
	RegisterTime string `json:"register_time"`
	
	Achieve     struct{
		AchieveId  	int	  	`json:"achieve_id"`
		Name 	   	string 	`json:"name"`
		Description	string 	`json:"description"`
	}`json:"achieve"`

	Role     struct{
		RoleId  	int	  	`json:"role_id"`
		Name 	   	string 	`json:"name"`
	}`json:"role"`

}