package ent

type User struct{
	UserId   	int	  	`json:"user-id" db:"id"`
	Name 	   	string 	`json:"name" db:"name"`
	LastName 	string 	`json:"last-name" db:"last_name"`
	Number 		string  `json:"number" db:"numbers"`
	Sex 		bool    `json:"sex" db:"sex"`
	Mail        string	`json:"mail" db:"email"`
	Password	string 	`json:"password" db:"password_hash"`
	RegisterTime string `json:"register-time" db:"registration_datetime"`
	Role_id 	int 	`json:"role-id" db:"role_id"`
	// Photo       []string  `json:"image"`
}