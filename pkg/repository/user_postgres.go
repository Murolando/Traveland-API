package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"traveland/ent"

	"github.com/jmoiron/sqlx"
	// "github.com/lib/pq"
)

type UserBD struct {
	db *sqlx.DB
}

func NewUserBD(db *sqlx.DB) *UserBD {
	return &UserBD{
		db: db,
	}
}

func (r UserBD) GetUserByID(id int) (ent.User, error) {
	var user ent.User
	query := fmt.Sprintf("SELECT id,name,last_name,password_hash,role_id,email,sex,registration_datetime FROM \"%s\" WHERE id = $1", userTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&user.UserId, &user.Name, &user.LastName, &user.Password, &user.Role_id, &user.Mail, &user.Sex, &user.RegisterTime); err != nil {
		return ent.User{}, err
	}
	pht,err := r.getPhoto(id)
	if err != nil{
		return ent.User{}, err
	}
	user.Photo = append(user.Photo, pht)
	return user, nil
}
func (r UserBD) GetAllUsers() ([]ent.User, error) {
	users := make([]ent.User, 0)
	query := fmt.Sprintf("SELECT id,name,last_name,password_hash,role_id,email,sex,registration_datetime FROM \"%s\"", userTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user ent.User

		if err := rows.Scan(&user.UserId, &user.Name, &user.LastName, &user.Password, &user.Role_id, &user.Mail, &user.Sex, &user.RegisterTime); err != nil {
			return nil, err
		}
		pht,err := r.getPhoto(user.UserId)
		if err != nil{
			return nil, err
		}
		user.Photo = append(user.Photo, pht)
		users = append(users, user)
	}
	return users, nil
}
func (r UserBD) GetUsersByRole(role_id int, offset int) ([]ent.User, error) {
	users := make([]ent.User, 0)
	query := fmt.Sprintf("SELECT id,name,last_name,password_hash,role_id,email,sex,registration_datetime FROM \"%s\" WHERE role_id = $1 OFFSET $2 LIMIT $3", userTable)
	rows, err := r.db.Query(query, role_id, offset, limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user ent.User

		if err := rows.Scan(&user.UserId, &user.Name, &user.LastName, &user.Password, &user.Role_id, &user.Mail, &user.Sex, &user.RegisterTime); err != nil {
			return nil, err
		}
		pht,err := r.getPhoto(user.UserId)
		if err != nil{
			return nil, err
		}
		user.Photo = append(user.Photo, pht)
		users = append(users, user)
	}
	return users, nil
}

func (r UserBD) UpdateUserInfo(user ent.User) (bool, error) {
	query := fmt.Sprintf(`UPDATE "%s" SET name = $1,last_name = $2,password_hash = $3,email = $4,sex = $5 WHERE id = $6`, userTable)
	_, err := r.db.Exec(query, user.Name, user.LastName, user.Password, user.Mail, user.Sex, user.UserId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r UserBD) getPhoto(userId int) (string, error) {
	var photo string
	err := filepath.Walk(fmt.Sprintf("./storage/user/%d", userId),
		// err := filepath.Walk("./storage/place",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				photo = path
				return nil
			}
			return nil
		})
	if err != nil {
		return "", err
	}
	return photo, nil
}
