package handler

import (
	"traveland/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRountes() *gin.Engine {
	router := gin.New()
	router.Static("/storage","./storage")
	// router.StaticFS("/more_static", http.Dir("my_file_system"))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		place := api.Group("/place")
		{
			place.GET("/get-place/:id", h.getPlaceByID)

			queryParams := place.Group("/",h.placeQueryParametrs)
			{
				queryParams.GET("/get-all-place/:place-ind/:offset", h.getAllPlace)
			}
			// place.GET("/get-place-by-type/:type-id/:offset", h.getLocalByType)
			// place.GET("/get-house-by-type/:type-id/:offset", h.getHouseByType)

			place.GET("/get-local-types",h.getLocalTypes)
			place.GET("/get-house-types",h.getHouseTypes)


			place.GET("/get-count-of-place-favorites/:place-id",h.getCountOfPlaceFavorites)
			
			authPlace := place.Group("/",h.userIdentity,h.placeQueryParametrs)
			{
				authPlace.POST("/add-favorite-place/",h.addFavoritePlace)
				authPlace.GET("/get-all-user-favorite-places",h.getAllUserFavoritePlaces)
			}
			
			
		}
		review := api.Group("/review")
		{
			authReview := review.Group("/",h.userIdentity)
			{
				authReview.POST("/add-review/", h.addReview)
				authReview.DELETE("/delete-review/:id", h.delteReview)
			}
			review.GET("/get-all-reviews/:place-id/:guide-id/:offset",h.getAllReview)

			// review.PUT("/update-review", h.updateReview)		
			
		}
		user := api.Group("/user",h.userIdentity)
		{
			// user.POST("/add-user", h.addUser)
			// user.DELETE("/delete-user/:id", h.delteUser)

			// user.POST("/add-photo/",h.addPhoto)

			user.DELETE("/delete-user", h.deleteUser)
			user.POST("/update-user", h.updateUser)
			user.GET("/get-user", h.getUserByID)
			
			user.GET("/get-all-users", h.getAllUsers)


			user.GET("/get-users-by-role/:role-id/:offset", h.getUsersByRole)
		}
		tour := api.Group("/tour",h.userIdentity)
		{

			tour.POST("/add-user-tour/",h.addUserTour)
			tour.GET("/get-all-user-tours/:user-id/:offset",h.getAllUserTours)

			tour.DELETE("/delete-user-tour/:tour-id",h.deleteUserTour)

			tour.GET("/get-all-guide-tours/:offset",h.getAllGuideTours)

			tour.GET("/get-tour-info/:tour-id",h.getTourInfo)
			// tour.GET("/get-all-tours/",h.getAllTours)
			
		}
	}
	return router
}
