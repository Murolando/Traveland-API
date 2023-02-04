package handler

import (
	"net/http"
	"strconv"
	"traveland/ent"

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
	newResponse(c, "place", place)
}

func (h *Handler) getAllPlace(c *gin.Context) {
	// place category
	id, err := strconv.Atoi(c.Param("place-ind"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// query params
	params,ok :=c.Keys["placeQueryParams"].(*ent.PlaceQueryParams)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "placeQueryParams not found")
		return
	}
	// Костылек)
	if (id == 1 && params.SortBy == "min_price"){
		params.SortBy = "house_price"
	}
	places, err := h.service.GetAllPlaces(id, params)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "places", places)
}
func (h *Handler) getAllPlacesBySearch(c *gin.Context){
	searchStr := c.Param("search-string")
	params,ok := c.Keys["placeQueryParams"].(*ent.PlaceQueryParams)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "placeQueryParams not found")
		return
	}
	params.SearchStr = searchStr
	places, err := h.service.GetAllPlacesBySearch(params)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "places", places)
}
func (h *Handler) getBannerPlaces(c *gin.Context){
	bannerId, err := strconv.Atoi(c.Param("banner-id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	places,err := h.service.GetBannerPlaces(bannerId)
	if err!=nil{
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	newResponse(c, "places", places)
}

func (h *Handler) getLocalTypes(c *gin.Context) {
	localTypes, err := h.service.GetLocalTypes()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "local-types", localTypes)
}
func (h *Handler) getHouseTypes(c *gin.Context) {
	houseTypes, err := h.service.GetHouseTypes()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "house-types", houseTypes)
}

func (h *Handler) addFavoritePlace(c *gin.Context) {
	var input ent.FavoritePlace
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
	result, err := h.service.AddFavoritePlace(input.UserId, input.PlaceId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResponse(c, "", result)
}

func (h *Handler) getAllUserFavoritePlaces(c *gin.Context) {
	numId, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "userCtx not found")
		return
	}
	id := numId.(int)
	favPlaces, err := h.service.GetAllUserFavoritePlaces(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "favorite-places", favPlaces)
}

func (h *Handler) getCountOfPlaceFavorites(c *gin.Context) {
	placeId, err := strconv.Atoi(c.Param("place-id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	counts, err := h.service.GetCountOfPlaceFavorites(placeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "counts", counts)
}
