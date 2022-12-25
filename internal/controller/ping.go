package controller

import (
	"github.com/Waelson/internal/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingController interface {
	Handler(c *gin.Context)
}

type pingController struct {
	logger log.Logger
}

func (a *pingController) Handler(c *gin.Context) {
	a.logger.Error(c.Request.Context(), "Ping")
	c.JSON(http.StatusOK, gin.H{"message": "ping"})
}

func NewPingController() PingController {
	result := new(pingController)
	result.logger = log.NewLogger("PingController")
	return result
}
