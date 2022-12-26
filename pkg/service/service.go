package service

import (
	"traveland/ent"
	"traveland/pkg/repository"
)
type Review interface{
	AddReview(review ent.Review) (int,error)
	DeleteReview(id int)(bool,error)
	GetAllReview(placeId int,guideId int, offset int)([]ent.Review,error)
	UpdateReview(reviewId int,rating int, reviewText string) (bool,error)
}
type Authorization interface {
	CreateUser(ent.User) (int, error)
	GenerateToken(mail string, password string) (string, error, int)
	ParseToken(token string) (int,error)
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
	GetUserByID(id int) (ent.User,error)
	GetAllUsers() ([]ent.User,error)
	GetUsersByRole(role_id int,offset int) ([]ent.User,error)
	UpdateUserInfo(user ent.User)(bool,error)
}

type Service struct {
	Authorization
	Place
	User
	Review
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		User: NewUserCRUDService(repo.User),
		Place: NewPlaceService(repo.Place),
		Review: NewReviewService(repo.Review),
	}
}
