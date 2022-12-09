package main

import (
	"log"
	"traveland"
	"traveland/pkg/handler"
	"traveland/pkg/repository"
	"traveland/pkg/service"

	_"github.com/lib/pq"
)

func main() {
	
	dbConfig := repository.NewConfig("localhost",5432,"postgres","123","app_bd")
	db,err := repository.NewPostgresDB(dbConfig)
	if err != nil{
		log.Fatal("Problems with db connect ",err)
	}
    repository := repository.NewRepository(db)
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

