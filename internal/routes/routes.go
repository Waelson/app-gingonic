package routes

import (
	"github.com/Waelson/internal/controller"
	"github.com/gin-gonic/gin"
)

type Option = func(*routes)

type Routes interface {
	Config(engine *gin.Engine)
}

type routes struct {
	engine           *gin.Engine
	pingController   controller.PingController
	clientController controller.ClientController
}

func (c *routes) Config(engine *gin.Engine) {
	v1 := c.engine.Group("/api/v1")
	v1.POST("/clients", c.clientController.Save)
	c.engine.GET("/ping", c.pingController.Handler)
}

func NewRoutes(options ...Option) Routes {
	result := new(routes)
	for _, opt := range options {
		opt(result)
	}
	return result
}

func WithPingController(controller controller.PingController) Option {
	return func(r *routes) {
		r.pingController = controller
	}
}

func WithClientController(controller controller.ClientController) Option {
	return func(r *routes) {
		r.clientController = controller
	}
}

func WithEngine(engine *gin.Engine) Option {
	return func(r *routes) {
		r.engine = engine
	}
}
