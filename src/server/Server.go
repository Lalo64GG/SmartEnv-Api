package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/SmartEnv-api/src/config"
	"github.com/lalo64/SmartEnv-api/src/kafka"
	record "github.com/lalo64/SmartEnv-api/src/records/infraestructure/http/routes"
	"github.com/lalo64/SmartEnv-api/src/users/infraestructure/http/routes"
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
		engine:   gin.Default(),
		http:     http,
		port:     port,
		httpAddr: http + ":" + port,
	}

	config.Connect()
	srv.engine.RedirectTrailingSlash = true
	srv.registerRoutes()
	
	// Iniciar el consumidor en segundo plano
	
	return srv
}

func (s *Server) registerRoutes() {
	userRoutes := s.engine.Group("/v1/users")
	recordRoutes := s.engine.Group("/v1/records")
	kafkaRoute := s.engine.Group("/v1/consumer")

	// Ruta para obtener los registros
	kafkaRoute.GET("/message", kafka.GetRecords)

	// Otras rutas
	record.RecordRoutes(recordRoutes)
	routes.UserRoutes(userRoutes)
}

func (s *Server) Run() error {
	
	log.Println("Server running on " + s.httpAddr)
	return s.engine.Run(s.httpAddr)
}
