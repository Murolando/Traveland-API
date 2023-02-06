package main

import (
	"log"
	"os"
	"traveland"
	"traveland/pkg/handler"
	"traveland/pkg/repository"
	"traveland/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error init config", err)
	}
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env variables:", err)
	}
	// dbConfig := repository.NewConfig("database", 5432, "postgres", "123", "postgres")
	dbConfig := repository.NewConfig(
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.username"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.dbname"),
	)
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatal("Problems with db connect ", err)
	}
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	srv := new(traveland.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRountes()); err != nil {
		log.Fatal(err)
	}

}

// migrate -path ./db/migrations -database 'postgres://postgres:123@localhost:5437/postgres?sslmode=disable' up

// migrate -path ./db/migrations -database 'postgres://postgres:123@localhost:5432/app_bd?sslmode=disable' up
/*const (
	host     = "localhost"
	port     = 5437
	user     = "postgres"
	password = "123"
	dbname   = "api_db"
)*/

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
