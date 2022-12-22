package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingController interface {
	Handler(c *gin.Context)
}

type pingController struct {
}

func (a *pingController) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ping"})
}

func NewPingController() PingController {
	result := new(pingController)
	return result
}
