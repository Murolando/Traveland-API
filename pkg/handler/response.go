package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context,statusCode int, message string) {
	fmt.Println(message)

	c.AbortWithStatusJSON(statusCode,Error{message})
}
