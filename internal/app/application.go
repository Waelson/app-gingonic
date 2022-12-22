package app

import (
	"fmt"
	"github.com/Waelson/internal/controller"
	"github.com/Waelson/internal/routes"
	"github.com/Waelson/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Option = func(*application)

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
}

func (a *application) Start() error {
	log.Println("Starting server...")
	a.initialize()
	error := a.engine.Run(fmt.Sprintf(":%d", a.port))
	return error
}

func (a *application) initialize() {
	engine := gin.Default()

	clientService := service.NewClientService()
	clientController := controller.NewClientController(controller.WithClientService(clientService))
	pingController := controller.NewPingController()

	routes := routes.NewRoutes(routes.WithClientController(clientController),
		routes.WithPingController(pingController), routes.WithEngine(engine))

	a.routes = routes
	a.engine = engine
}

func NewApplication(options ...Option) Application {
	result := new(application)
	for _, opt := range options {
		opt(result)
	}
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

func WithMaxConn(maxConnection int) Option {
	return func(s *application) {
		s.maxConnection = maxConnection
	}
}
