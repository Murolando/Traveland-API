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
	numId, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "userCtx not found")
		return
	}
	id := numId.(int)
	input.UserId = id

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
	numId, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "userCtx not found")
		return
	}
	userId := numId.(int)
	params,ok :=c.Keys["tourQueryParams"].(*ent.TourQueryParams)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "reviewQueryParams not found")
		return
	}

	tours,err := h.service.Tour.GetUserTours(userId,params)
	if err!=nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "tours", tours)
}
func (h *Handler) deleteUserTour(c *gin.Context) {
	numId, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "userCtx not found")
		return
	}
	userId := numId.(int)
	tourId, err := strconv.Atoi(c.Param("tour-id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := h.service.Tour.DeleteTour(tourId, userId)
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
	params,ok :=c.Keys["tourQueryParams"].(*ent.TourQueryParams)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "reviewQueryParams not found")
		return
	}

	tours,err := h.service.Tour.GetAllGuideTours(params)
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
