package repository

import (
	"traveland/ent"

	"github.com/jmoiron/sqlx"
)
type Review interface{
	AddReview(review ent.Review) (int, error) 
	DeleteReview(id int)(bool,error)
	GetAllReview(placeId int,guideId int, offset int)([]ent.Review,error)
	UpdateReview(reviewId int,rating int, reviewText string) (bool,error)
}
type Authorization interface {
	CreateUser(user ent.User) (int, error)
	GetUserByMailAndPassword(mail string , password string)(int, error)
	
}

type Place interface {
	GetPlaceByID(id int) (interface{}, error)
	GetAllPlaces(placeInd int,offset int) (interface{}, error)

	GetLocalByType(placeType int,offset int) (*[]ent.Location, error)
	GetHouseByType(houseType int,offset int) (*[]ent.Housing, error)

	GetLocalTypes() (*[]ent.LocalType,error)
	GetHouseTypes() (*[]ent.HouseType,error)
}
type User interface{
	GetUserByID(id int) (ent.User, error)
	GetAllUsers()([]ent.User,error)
	GetUsersByRole(role_id int) ([]ent.User,error)
	UpdateUserInfo(user ent.User)(bool,error)

}
type Repository struct {
	Authorization
	Place
	User
	Review
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User: NewUserBD(db),
		Place: NewPlaceBD(db),
		Review: NewReviewBD(db),
	}
}
