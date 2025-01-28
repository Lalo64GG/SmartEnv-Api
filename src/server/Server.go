package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/SmartEnv-api/src/config"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http/routes"
	record "github.com/lalo64/SmartEnv-api/src/records/infraestructure/http/routes"
)

type Server struct {
	engine   *gin.Engine
	http     string
	port     string
	httpAddr string
}

func NewServer(http, port string) Server {

	gin.SetMode(gin.ReleaseMode)

	srv := Server{
		engine:   gin.New(),
		http:     http,
		port:     port,
		httpAddr: http + ":" + port,
	}

	config.Connect()
	srv.engine.Use(gin.Logger())
	srv.engine.RedirectTrailingSlash = true
	srv.registerRoutes()

	return srv
}


func (s *Server) registerRoutes() {
    userRoutes := s.engine.Group("/v1/users")
	recordRoutes := s.engine.Group("/v1/records")

	
	record.RecordRoutes(recordRoutes)
	routes.UserRoutes(userRoutes)

	
}

func (s *Server) Run() error {
	log.Println("Server running on " + s.httpAddr)
	return s.engine.Run(s.httpAddr)
}