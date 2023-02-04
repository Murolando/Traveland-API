package repository

import (
	"traveland/ent"

	"github.com/jmoiron/sqlx"
)
type Review interface{
	AddReview(review ent.Review) (int, error) 
	DeleteReview(id int,userId int)(bool,error)
	GetAllReview(params *ent.ReviewQueryParams)([]ent.Review,error)
	UpdateReview(reviewId int,rating int, reviewText string) (bool,error)
}
type Authorization interface {
	CreateUser(user ent.User) (int, error)
	GetUserByMailAndPassword(mail string , password string)(int, error)
	
}

type Place interface {
	GetPlaceByID(id int) (interface{}, error)
	GetAllPlaces(placeInd int,params *ent.PlaceQueryParams) (interface{}, error)
	GetAllPlacesBySearch(params *ent.PlaceQueryParams)(interface{},error)
	GetBannerPlaces(bannerId int)(*[]ent.Banner,error)

	GetLocalTypes() (*[]ent.LocalType,error)
	GetHouseTypes() (*[]ent.HouseType,error)

	AddFavoritePlace(userId int, placeId int) (bool, error)
	GetAllUserFavoritePlaces(userId int) (*[]interface{}, error)
	GetCountOfPlaceFavorites(placeId int) (int, error)
}
type User interface{
	GetUserByID(id int) (ent.User, error)
	GetAllGuides()([]ent.User,error)
	GetUsersByRole(role_id int,offset int) ([]ent.User,error)
	UpdateUserInfo(user ent.User)(bool,error)
	AddPhoto(userId int,photo []byte,imgExt string) (bool, error)
	DeleteUser(userId int) (bool, error)
}
type Tour interface{
	AddUserTour(fullTour ent.Tour) (int, error)
	GetUserTours(userId int,params *ent.TourQueryParams)(*[]ent.Tour, error)
	DeleteTour(tourId int,userId int)(bool,error)
	GetAllGuideTours(params *ent.TourQueryParams)(*[]ent.Tour,error)
	GetTourInfo(tourId int)(*ent.Tour,error)
}
type Repository struct {
	Authorization
	Place
	User
	Review
	Tour
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User: NewUserBD(db),
		Place: NewPlaceBD(db),
		Review: NewReviewBD(db),
		Tour: NewTourBD(db),
	}
}
