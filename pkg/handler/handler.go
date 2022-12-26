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
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		place := api.Group("/place")
		{
			// place.POST("/add-place/", h.addPlace)
			// place.DELETE("/delete-place/:id", h.deltePlace)
			// place.PUT("/update-place/:id", h.updatePlace)

			
			place.GET("/get-place/:id", h.getPlaceByID)
			place.GET("/get-all-place/:place-ind/:offset", h.getAllPlace)

			place.GET("/get-place-by-type/:type-id/:offset", h.getLocalByType)
			place.GET("/get-house-by-type/:type-id/:offset", h.getHouseByType)

			place.GET("/get-local-types",h.getLocalTypes)
			place.GET("/get-house-types",h.getHouseTypes)

			
		}
		review := api.Group("/review")
		{
			review.POST("/add-review/", h.addReview)
			review.DELETE("/delete-review/:id", h.delteReview)
			review.GET("/get-all-reviews/:place-id/:guide-id/:offset",h.getAllReview)

			// review.PUT("/update-review", h.updateReview)		
			
		}
		user := api.Group("/user")
		{
			// user.POST("/add-user", h.addUser)
			// user.DELETE("/delete-user/:id", h.delteUser)
			user.POST("/update-user/", h.updateUser)
			user.GET("/get-user/:id", h.getUserByID)
			user.GET("/get-all-users/", h.getAllUsers)
			user.GET("/get-users-by-role/:role-id/:offset", h.getUsersByRole)
		}
	}
	return router
}
