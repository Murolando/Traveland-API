package handler

import (
	// "strconv"
	// "io/ioutil"
	"net/http"
	"traveland/ent"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input ent.User

	// form read
	// roleId,_ := strconv.Atoi(c.PostForm("role-id"))
	// sex,_  := strconv.ParseBool(c.PostForm("sex"))
	// input = ent.User{
	// 	Name:     c.PostForm("name"),
	// 	LastName: c.PostForm("last-name"),
	// 	Role_id:  roleId,
	// 	Mail:     c.PostForm("mail"),
	// 	Password: c.Request.FormValue("password"),
	// 	Sex:      sex,
	// }

	// // fmt.Println(input)
	// file, _, err := c.Request.FormFile("image")
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	// defer file.Close()

	// fileBytes, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	// input.Photo = fileBytes

	// json read
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "id", id)
}

type singInInput struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input ent.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err, userId := h.service.Authorization.GenerateToken(input.Mail, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "", map[string]interface{}{"token": token, "user-id": userId})
}
