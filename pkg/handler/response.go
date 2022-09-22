package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"messsage"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, messsage string) {
	logrus.Error(messsage)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: messsage})
}
