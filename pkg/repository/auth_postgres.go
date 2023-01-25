package repository

import (
	"fmt"
	"io"
	"strconv"
	"os"
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

func (r AuthPostgres) CreateUser(user ent.User,realPass string) (int, error) {
	var id int
	t := time.Now()
	user.RegisterTime = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	if user.Number != "" {
		query := fmt.Sprintf("INSERT INTO \"%s\" (name, last_name, password_hash, numbers, role_id, sex,registration_datetime) values ($1 , $2 ,$3 ,$4, $5, $6, $7) RETURNING id", userTable)
		row := r.db.QueryRow(query, user.Name,user.LastName, user.Password, user.Number, user.Role_id, user.Sex, user.RegisterTime)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
	} else {
		query := fmt.Sprintf("INSERT INTO \"%s\" (name,password_hash, role_id, email, sex, registration_datetime) values ($1 , $2 ,$3 ,$4, $5, $6) RETURNING id", userTable)
		row := r.db.QueryRow(query, user.Name, user.Password, user.Role_id, user.Mail, user.Sex, user.RegisterTime)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}

	}
	// err := r.setDefault(id)
	// if err != nil {
	// 	return 0, err
	// }
	return id, nil
}

func (r AuthPostgres) GetUserByMailAndPassword(mail string, password string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM \"%s\" WHERE email = $1", userTable)
	row := r.db.QueryRow(query, mail)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	query = fmt.Sprintf("SELECT id FROM \"%s\" WHERE email = $1 AND password_hash = $2", userTable)
	row = r.db.QueryRow(query, mail, password)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}


func (r AuthPostgres) setDefault(id int) error {
	// fullFileName := fmt.Sprintf("%s.%s", "default", "png")
	// открываем папку для сохранения картинки
	path := fmt.Sprintf("./storage/user/%s", strconv.Itoa(id))

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	// fileOnDisk, err := os.Create(fmt.Sprintf("%s/%s", path, fullFileName))
	// if err != nil {
	// 	return err
	// }
	// defer fileOnDisk.Close()
	// путь до исходной картинки
	stokPlace := fmt.Sprintf("./storage/stok/change-user-account-image-windows-11.png")
	source, err := os.Open(stokPlace)
	if err != nil {
		return err
	}

	// создаем нвоый файл
	new, err := os.Create(fmt.Sprintf("%s/%s.%s",path,strconv.Itoa(id),"png"))
	if err != nil {
		return err
	}

	defer source.Close()
	defer new.Close()
	// fmt.Println(2)
	_, err = io.Copy(new, source)
	return nil
}
