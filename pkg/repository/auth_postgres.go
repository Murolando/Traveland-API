package repository

import (
	"fmt"
	"traveland/ent"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r AuthPostgres) CreateUser(user ent.User) (int, error) {
	var id int
	if user.Number != ""{
		query := fmt.Sprintf("INSERT INTO \"%s\" (name, last_name, password_hash, numbers, role_id, sex,registration_datetime) values ($1 , $2 ,$3 ,$4, $5, $6, $7) RETURNING id", userTable)
		row := r.db.QueryRow(query,user.Name, user.LastName, user.Password, user.Number,user.Role_id, user.Sex, user.RegisterTime)
		row.Scan(&id)
		if err := 	row.Scan(&id);err!=nil{
			return 0,err
		}
	} else {
		query := fmt.Sprintf("INSERT INTO \"%s\" (name,last_name, password_hash, role_id, email, sex, registration_datetime) values ($1 , $2 ,$3 ,$4, $5, $6,$7) RETURNING id", userTable)
		row := r.db.QueryRow(query,user.Name, user.LastName, user.Password, user.Role_id, user.Mail,user.Sex, user.RegisterTime)
		if err := 	row.Scan(&id);err!=nil{
			return 0,err
		}
	}
	return id,nil
}
