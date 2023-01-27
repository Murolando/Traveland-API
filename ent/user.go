package ent

import "database/sql"

type User struct {
	UserId       int            `json:"user-id" db:"id"`
	Name         string         `json:"name" db:"name"`
	LastName     sql.NullString `json:"last-name" db:"last_name"`
	Number       sql.NullString `json:"number" db:"numbers"`
	Sex          bool           `json:"sex" db:"sex"`
	Mail         string         `json:"mail" db:"email"`
	Password     string         `json:"password" db:"password_hash"`
	RegisterTime string         `json:"register-time" db:"registration_datetime"`
	Role_id      int            `json:"role-id" db:"role_id"`
	Image_src    sql.NullString `json:"image"`
	// Session     Session `json:"session"`
}
type Session struct {
	RefreshToken string `json:"refres-token"`
	ExpiredAt    int64  `json:"expired-at"`
}
