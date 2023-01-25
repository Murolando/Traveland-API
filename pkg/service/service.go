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
	CreateUser(user ent.User) (map[string]interface{}, error)
	SignIn(mail string,password string)(int,error)
	GenerateToken(id int) (string,error)
	ParseToken(token string) (int,error)
}

type Place interface {
	GetPlaceByID(id int) (interface{}, error)
	GetAllPlaces(placeInd int,offset int) (interface{}, error)

	GetLocalByType(placeType int,offset int) (*[]ent.Location, error)
	GetHouseByType(houseType int,offset int) (*[]ent.Housing, error)

	GetLocalTypes() (*[]ent.LocalType,error)
	GetHouseTypes() (*[]ent.HouseType,error)

	AddFavoritePlace(userId int, placeId int) (bool, error)
	GetAllUserFavoritePlaces(userId int) (*[]interface{}, error)
	GetCountOfPlaceFavorites(placeId int) (int, error)
}
type User interface{
	GetUserByID(id int) (ent.User,error)
	GetAllUsers() ([]ent.User,error)
	GetUsersByRole(role_id int,offset int) ([]ent.User,error)
	UpdateUserInfo(user ent.User)(bool,error)
	AddPhoto(userId int,photo []byte,imgExt string) (bool,error)
}
type Tour interface{
	AddUserTour(newTour ent.AddPoints)(int,error)
	GetUserTours(userId int,offset int) (*[]ent.Tour, error)
	DeleteTour(tourId int)(bool,error)
	GetAllGuideTours(offset int)(*[]ent.Tour,error)
	GetTourInfo(tourId int)(*ent.Tour,error)
}

type Service struct {
	Authorization
	Place
	User
	Review
	Tour
}


func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		User: NewUserCRUDService(repo.User),
		Place: NewPlaceService(repo.Place),
		Review: NewReviewService(repo.Review),
		Tour: NewTourService(repo.Tour),
	}
}
