package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_"github.com/lib/pq"
)
const (
	userTable = "user"
	roleTable = "role"
	favoritePlaceTable = "favorite_place"
	placeTable = "place"
	houseTypeTable ="house_type"
	reviewTable = "review"
	typeTable = "type"
	placeTypeTable = "place_type"
	achieveTable = "achieve"
	userAchiveTable= "user_achieve"
	weekTable = "week"
	dayyTable = "dayy"
	tourTable = "tour"
	tourPlaceTable = "tour_place"
)

const (
	limit int = 20
)

type Config struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func NewConfig(host string, port int, user string, password string, dbname string) *Config {
	return &Config{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}
}

// pq - driver ,sqlx - interfase  , pgx - driver+interfase
/*Драйвер это реализация интерфейсов для sql(sqlx - просто расширение sql)*/
func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host = %s port = %d user = %s dbname = %s password = %s  sslmode = disable",
		cfg.host, cfg.port, cfg.user, cfg.dbname, cfg.password)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
