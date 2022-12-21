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
	typeId, err := strconv.Atoi(c.Param("type-id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	places, err := h.service.GetLocalByType(typeId, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c,"locals",places)

}

func (h *Handler) getHouseByType(c *gin.Context) {
	typeId, err := strconv.Atoi(c.Param("type-id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	places, err := h.service.GetHouseByType(typeId, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c,"houses",places)
}
func (h *Handler) getLocalTypes(c *gin.Context) {
	localTypes, err := h.service.GetLocalTypes()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c,"local-types",localTypes)
}

func (h *Handler) getHouseTypes(c *gin.Context) {
	houseTypes, err := h.service.GetHouseTypes()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c,"house-types",houseTypes)
}
