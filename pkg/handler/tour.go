package handler

import (
	"net/http"
	"traveland/ent"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addUserTour(c *gin.Context) {
	var input ent.AddPoints
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := h.service.Tour.AddUserTour(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "", result)
}
func (h *Handler) getAllUserTours(c *gin.Context) {

}
func (h *Handler) deleteUserTour(c *gin.Context) {

}
func (h *Handler) getAllGuideTours(c *gin.Context) {

}
func (h *Handler) getTourInfo(c *gin.Context) {

}
func (h *Handler) getAllTours(c *gin.Context) {

}
