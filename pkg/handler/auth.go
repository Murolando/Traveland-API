package handler

import (
	"net/http"
	"traveland/ent"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input ent.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c,"id",id)
}

type singInInput struct{
	Mail        string	`json:"mail"`
	Password	string 	`json:"password"`
}
func (h *Handler) signIn(c *gin.Context) {
	var input ent.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}
	token, err, userId := h.service.Authorization.GenerateToken(input.Mail, input.Password)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c,"",map[string]interface{}{"token":token,"user-id":userId})
}
