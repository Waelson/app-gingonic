package app

import (
	"context"
	"fmt"
	"github.com/Waelson/internal/controller"
	"github.com/Waelson/internal/log"
	"github.com/Waelson/internal/routes"
	"github.com/Waelson/internal/service"
	"github.com/gin-gonic/gin"
	"time"
)

type Option func(*application)

type Application interface {
	Start() error
}

type application struct {
	host          string
	port          int
	timeout       time.Duration
	maxConnection int
	engine        *gin.Engine
	routes        routes.Routes
	logger        log.Logger
}

func (a *application) Start() error {
	a.logger.Infof(context.Background(), "Starting server on port %d", a.port)
	error := a.engine.Run(fmt.Sprintf(":%d", a.port))
	return error
}

func (a *application) configDependencies() {

	//Services
	clientService := service.NewClientService()

	//Controllers
	clientController := controller.NewClientController(controller.WithClientService(clientService))
	pingController := controller.NewPingController()

	//Gin
	a.engine = gin.New()

	//Routes
	a.routes = routes.NewRoutes(routes.WithClientController(clientController),
		routes.WithPingController(pingController), routes.WithEngine(a.engine))

	//Setting routes
	a.routes.Config()
}

// NewApplication Create application instance
func NewApplication(options ...Option) Application {
	result := new(application)
	result.logger = log.NewLogger("Application")
	for _, opt := range options {
		opt(result)
	}
	result.configDependencies()
	return result
}

func WithHost(host string) Option {
	return func(s *application) {
		s.host = host
	}
}

func WithPort(port int) Option {
	return func(s *application) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *application) {
		s.timeout = timeout
	}
}

func WithMaxConnection(maxConnection int) Option {
	return func(s *application) {
		s.maxConnection = maxConnection
	}
}
