package ent

type User struct{
	UserId   	int	  	`json:"user-id"`
	Name 	   	string 	`json:"name"`
	LastName 	string 	`json:"last-name"`
	Number 		string  `json:"number"`
	Sex 		bool    `json:"sex"`
	Mail        string	`json:"mail"`
	Password	string 	`json:"password"`
	RegisterTime string `json:"register-time"`
	Role_id 	int 	`json:"role-id"`
}