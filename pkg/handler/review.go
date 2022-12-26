package handler

import (
	"net/http"
	"strconv"
	"traveland/ent"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addReview(c *gin.Context) {
	var input ent.Review
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Review.AddReview(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "id", id)
}
func (h *Handler) delteReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	t,err := h.service.Review.DeleteReview(id)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if t == true{
		newResponse(c, "", true)
	}else{
		newErrorResponse(c, http.StatusInternalServerError,"records are missing")
		return
	}
	
}
// func (h *Handler) updateReview(c *gin.Context) {

// }
func (h *Handler) getAllReview(c *gin.Context) {
	placeId, err := strconv.Atoi(c.Param("place-id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	guideId, err := strconv.Atoi(c.Param("guide-id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	reviews, err := h.service.Review.GetAllReview(placeId,guideId,offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "reviews", reviews)
}
