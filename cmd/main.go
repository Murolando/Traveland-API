package main

import (
	"log"
	"traveland"
	"traveland/pkg/handler"
	"traveland/pkg/repository"
	"traveland/pkg/service"
)

func main() {
    repository := repository.NewRepository()
    service := service.NewService(repository)
    handler := handler.NewHandler(service)

	srv := new(traveland.Server)
	if err := srv.Run("8000",handler.InitRountes()); err != nil {
		log.Fatal(err)
	}

}


// migrate -path ./db/migrations -database 'postgres://postgres:123@localhost:5432/app_bd?sslmode=disable' up
/*const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "api_db"
)*/
// pq - driver ,sqlx - interfase  , pgx - driver+interfase
