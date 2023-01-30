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


	numId, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "userCtx not found")
		return
	}
	userId := numId.(int)
	input.UserId = userId
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
	numId, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "userCtx not found")
		return
	}
	userId := numId.(int)
	t,err := h.service.Review.DeleteReview(id,userId)
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
	params,ok :=c.Keys["reviewQueryParams"].(*ent.ReviewQueryParams)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "reviewQueryParams not found")
		return
	}

	reviews, err := h.service.Review.GetAllReview(params)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "reviews", reviews)
}
