package handler

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
	roleId, err := strconv.Atoi(c.Param("role-id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	users, err := h.service.GetUsersByRole(roleId, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newResponse(c, "users", users)
}

//DONT WORK
func (h *Handler) addPhoto(c *gin.Context) {

	// form reader
	userId,_ := strconv.Atoi(c.PostForm("user-id"))

	form, _ := c.MultipartForm()
	// filename
	var fileName string
	imgExt := "jpeg"

	// берем первое имя файла из присланного списка
	for key := range form.File {
		fileName = key
		// извлекаем расширение файла
		arr := strings.Split(fileName, ".")
		if len(arr) > 1 {
			imgExt = arr[len(arr)-1]
		}
		continue
	}

	// извлекаем содержание присланного файла по названию файла
	file, _, err := c.Request.FormFile(fileName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	// читаем содержание присланного файл в []byte
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := h.service.AddPhoto(userId, fileBytes, imgExt)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	// берем первое имя файла из присланного списка
	newResponse(c, "", result)
}
