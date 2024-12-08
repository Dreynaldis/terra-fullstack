package server

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/delivery/route"
	"backend/internal/repository"
	"backend/internal/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	cfg    *config.Config
	db     database.Service
}

func NewServer(cfg *config.Config, db database.Service) *Server {
	s := &Server{
		router: gin.Default(),
		cfg:    cfg,
		db:     db,
	}

	repo := repository.NewUserRepository(db.Queries())
	userUsecase := usecase.NewUserUsecase(repo)

	route.RegisterRoutes(s.router, userUsecase)
	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {
	s.router.GET("/health", s.healthCheck)
}

func (s *Server) healthCheck(c *gin.Context) {
	health := s.db.Health()
	status := http.StatusOK
	if health["status"] == "down" {
		status = http.StatusServiceUnavailable
	}
	c.JSON(status, health)
}

func (s *Server) Run() error {
	addr := ":" + s.cfg.AppPort
	log.Printf("Server is running on %s", addr)
	return s.router.Run(addr)
}
