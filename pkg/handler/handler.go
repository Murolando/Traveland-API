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
			place.GET("/get-all-place/:place-ind", h.getAllPlace)
			place.GET("/get-place-by-type/:type-id", h.getPlaceByType)
		}
		user := api.Group("/user")
		{
			// user.POST("/add-user", h.addUser)
			// user.DELETE("/delete-user/:id", h.delteUser)
			// user.PUT("/update-user/:id", h.updateUser)

			user.GET("/get-user/:id", h.getUserByID)
			user.GET("/get-all-users/", h.getAllUsers)
			user.GET("/get-users-by-role/:role-id", h.getUsersByRole)
		}
	}
	return router
}
