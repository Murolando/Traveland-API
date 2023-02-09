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
	gin.SetMode(gin.ReleaseMode)
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

			queryParams := place.Group("/",h.placeQueryParams)
			{
				queryParams.GET("/get-all-place/:place-ind", h.getAllPlace)
			}
			// place.GET("/get-place-by-type/:type-id/:offset", h.getLocalByType)
			// place.GET("/get-house-by-type/:type-id/:offset", h.getHouseByType)

			place.GET("/get-local-types",h.getLocalTypes)
			place.GET("/get-house-types",h.getHouseTypes)


			place.GET("/get-count-of-place-favorites/:place-id",h.getCountOfPlaceFavorites)
			place.GET("/get-all-places-by-search/:search-string",h.searchQueryParams,h.getAllPlacesBySearch)
			place.GET("/get-banner-places/:banner-id",h.getBannerPlaces)
			authPlace := place.Group("/",h.userIdentity)
			{
				authPlace.POST("/add-favorite-place/",h.addFavoritePlace)
				authPlace.GET("/get-all-user-favorite-places",h.getAllUserFavoritePlaces)
			}
			
			
		}
		review := api.Group("/review")
		{
			authReview := review.Group("/",h.userIdentity)
			{
				authReview.POST("/add-review/",h.addReview)
				authReview.DELETE("/delete-review/:id", h.delteReview)
			}

			review.GET("/get-all-reviews",h.reviewQueryParams,h.getAllReview)

			// review.PUT("/update-review", h.updateReview)		
			
		}
		user := api.Group("/user")
		{
			// user.POST("/add-user", h.addUser)
			// user.DELETE("/delete-user/:id", h.delteUser)

			// user.POST("/add-photo/",h.addPhoto)
			userIden:= user.Group("/",h.userIdentity)
			{
				userIden.DELETE("/delete-user", h.deleteUser)
				userIden.POST("/update-user", h.updateUser)
				userIden.GET("/get-user", h.getUserByID)
			}
			
			
			user.GET("/get-all-guides", h.getAllGuides)


			// user.GET("/get-users-by-role/:role-id/:offset", h.getUsersByRole)
		}
		tour := api.Group("/tour")
		{
			tourAuth := tour.Group("/",h.userIdentity)
			{
				tourAuth.POST("/add-user-tour",h.addUserTour)
				tourAuth.GET("/get-all-user-tours",h.tourQueryParams,
					h.getAllUserTours)

				tourAuth.DELETE("/delete-user-tour/:tour-id",h.deleteUserTour)
			}


			tour.GET("/get-all-guide-tours",h.tourQueryParams,h.getAllGuideTours)
			tour.GET("/get-tour-info/:tour-id",h.getTourInfo)
			// tour.GET("/get-all-tours/",h.getAllTours)
			
		}
	}
	return router
}
