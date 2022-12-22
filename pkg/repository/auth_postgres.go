package repository

import (
	"fmt"
	"time"
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
	t := time.Now()
		user.RegisterTime = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	if user.Number != ""{
		query := fmt.Sprintf("INSERT INTO \"%s\" (name, last_name, password_hash, numbers, role_id, sex,registration_datetime) values ($1 , $2 ,$3 ,$4, $5, $6, $7) RETURNING id", userTable)
		row := r.db.QueryRow(query,user.Name, user.LastName, user.Password, user.Number,user.Role_id, user.Sex, user.RegisterTime)
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

func (r AuthPostgres)  GetUserByMailAndPassword(mail string , password string)(int, error){
	var id int
	query := fmt.Sprintf("SELECT id FROM \"%s\" WHERE email = $1",userTable)
	row := r.db.QueryRow(query,mail)
	if err := row.Scan(&id);err!=nil{
		return 0, err
	}
	query = fmt.Sprintf("SELECT id FROM \"%s\" WHERE email = $1 AND password_hash = $2",userTable)
	row = r.db.QueryRow(query,mail,password)
	if err := row.Scan(&id);err!=nil{
		return -1, err
	}
	return id,nil
}
