package controller

import (
	"github.com/Waelson/internal/service"
	"github.com/gin-gonic/gin"
)

type Option = func(controller *clientController)

type ClientController interface {
	Save(c *gin.Context)
}

type clientController struct {
	clientService service.ClientService
}

func (a *clientController) Save(c *gin.Context) {

}

func NewClientController(options ...Option) ClientController {
	result := new(clientController)
	for _, opt := range options {
		opt(result)
	}
	return result
}

func WithClientService(service service.ClientService) Option {
	return func(c *clientController) {
		c.clientService = service
	}
}
