package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addUser(c *gin.Context) {
}

func (h *Handler) delteUser(c *gin.Context) {

}
func (h *Handler) updateUser(c *gin.Context) {

}
func (h *Handler) getUserByID(c *gin.Context) {
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	user, err := h.service.GetUserByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
		"error": map[string]int{"code": 200,"description":0},
	})
	
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
		"error": map[string]int{"code": 200,"description":0},
	})
}

func (h *Handler) getUsersByRole(c *gin.Context) {
	role_id,err := strconv.Atoi(c.Param("role-id"))
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	users, err := h.service.GetUsersByRole(role_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
		"error": map[string]int{"code": 200,"description":0},
	})
}
