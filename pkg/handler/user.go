package handler

import (
	"net/http"
	"strconv"
	"traveland/ent"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addUser(c *gin.Context) {
}

func (h *Handler) delteUser(c *gin.Context) {

}
func (h *Handler) updateUser(c *gin.Context) {
	var input ent.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	end, err := h.service.User.UpdateUserInfo(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "", end)

}
func (h *Handler) getUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	user, err := h.service.GetUserByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "users", user)
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "users", users)
}

func (h *Handler) getUsersByRole(c *gin.Context) {
	role_id, err := strconv.Atoi(c.Param("role-id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	users, err := h.service.GetUsersByRole(role_id, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "users", users)
}
