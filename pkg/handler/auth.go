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

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
		"error": map[string]int{"code": 200,"description":0},
	})
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
	token, err := h.service.Authorization.GenerateToken(input.Mail, input.Password)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, map[string]interface{}{
		"result": token,
		"error": map[string]int{"code": 200,"description":0},
	})
}
