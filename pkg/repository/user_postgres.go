package repository

import (
	"fmt"
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
	query := fmt.Sprintf("SELECT id,name,last_name,role_id,email,numbers,sex,image_src FROM \"%s\" WHERE id = $1", userTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&user.UserId, &user.Name,&user.LastName,&user.Role_id, &user.Mail,&user.Number, &user.Sex,&user.Image_src); err != nil {
		return ent.User{}, err
	}
	// pht,err := r.getPhoto(id)
	// if err != nil{
	// 	return ent.User{}, err
	// }
	// user.Photo = append(user.Photo, pht)
	return user, nil
}
func (r UserBD) GetAllGuides() ([]ent.User, error) {
	users := make([]ent.User, 0)
	query := fmt.Sprintf(`SELECT id,name,
	last_name,
	sex,
	image_src
	FROM "%s"
	WHERE role_id = 2`, userTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user ent.User

		if err := rows.Scan(&user.UserId, &user.Name,&user.LastName,&user.Role_id, &user.Mail,&user.Number, &user.Sex, &user.Image_src); err != nil {
			return nil, err
		}
		// pht,err := r.getPhoto(user.UserId)
		// if err != nil{
		// 	return nil, err
		// }
		// user.Photo = append(user.Photo, pht)
		users = append(users, user)
	}
	return users, nil
}

func (r UserBD) UpdateUserInfo(user ent.User) (bool, error) {
	query := fmt.Sprintf(`UPDATE "%s" SET name = $1,password_hash = $2,sex = $3 WHERE id = $4`, userTable)
	_, err := r.db.Exec(query, user.Name, user.Password, user.Sex, user.UserId)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r UserBD) DeleteUser(userId int) (bool, error) {
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`, userTable)
	res, err := r.db.Exec(query,userId)
	if err!=nil{
		return false,err
	}
	count, err := res.RowsAffected()
    if err != nil {
        return false,err
    }
	if count != 0{
		return true,nil
	}
	return false,nil
}
