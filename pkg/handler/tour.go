package handler

import (
	"net/http"
	"strconv"
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
	if (result == -1){
		newErrorResponse(c, http.StatusInternalServerError, "exceeded the route limit")
		return
	}
	newResponse(c, "tour-id", result)
}
func (h *Handler) getAllUserTours(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	tours,err := h.service.Tour.GetUserTours(userId,offset)
	if err!=nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "tours", tours)
}
func (h *Handler) deleteUserTour(c *gin.Context) {
	tourId, err := strconv.Atoi(c.Param("tour-id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := h.service.Tour.DeleteTour(tourId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if result == true{
		newResponse(c, "", result)
	}else{
		newErrorResponse(c, http.StatusInternalServerError,"records are missing")
		return
	}
}
func (h *Handler) getAllGuideTours(c *gin.Context) {
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	tours,err := h.service.Tour.GetAllGuideTours(offset)
	if err!=nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "tours", tours)
}
func (h *Handler) getTourInfo(c *gin.Context) {
	tourId, err := strconv.Atoi(c.Param("tour-id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	tourInfo,err := h.service.Tour.GetTourInfo(tourId)
	if err!=nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "tour-info", tourInfo)
}
