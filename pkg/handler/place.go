package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getPlaceByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	place, err := h.service.GetPlaceByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c,"place",place)
}

func (h *Handler) getAllPlace(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("place-ind"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	offset, err  := strconv.Atoi(c.Param("offset"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	places, err := h.service.GetAllPlaces(id, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c,"places",places)
}
func (h *Handler) getLocalByType(c *gin.Context) {
	
}

func (h *Handler) getHouseByType(c *gin.Context) {
	
}
func (h *Handler) getPlaceTypes(c *gin.Context) {
	
}

func (h *Handler) getHouseTypes(c *gin.Context) {
	
}
