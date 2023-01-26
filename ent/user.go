package ent

type User struct {
	UserId       int    `json:"user-id" db:"id"`
	Name         string `json:"name" db:"name"`
	LastName     string `json:"last-name" db:"last_name"`
	Number       string `json:"number" db:"numbers"`
	Sex          bool   `json:"sex" db:"sex"`
	Mail         string `json:"mail" db:"email"`
	Password     string `json:"password" db:"password_hash"`
	RegisterTime string `json:"register-time" db:"registration_datetime"`
	Role_id      int    `json:"role-id" db:"role_id"`
	// Session     Session `json:"session"`
	// Photo       []string  `json:"image"`
}
type Session struct {
	RefreshToken string `json:"refres-token"`
	ExpiredAt    int64	`json:"expired-at"`
}
