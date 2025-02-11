package server

import (
	"os"

	"user-service/configs"
	"user-service/internal/handler"

	"github.com/gin-gonic/gin"

	"github.com/sherinur/movie-reservation-system/pkg/logging"
	"github.com/sherinur/movie-reservation-system/pkg/middleware"
)

type Server interface {
	Start() error
	Shutdown()
	registerRoutes() error
}

type server struct {
	router *gin.Engine
	cfg    *configs.Config
	log    *logging.Logger

	userHandler handler.UserHandler
}

func NewServer(cfg *configs.Config) Server {
	r := gin.Default()

	corsConfig := &middleware.CorsConfig{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}

	// cors middleware
	middleware.SetCorsConfig(corsConfig)
	r.Use(middleware.CorsMiddleware())

	// jwt middleware
	middleware.SetSecret([]byte(cfg.JwtAccessSecret))

	return &server{
		router: r,
		cfg:    cfg,
		log:    logging.NewLogger("dev"),
	}
}

func (s *server) Start() error {
	s.log.Info("Registering routes...")
	err := s.registerRoutes()
	if err != nil {
		s.log.Errorf("Could not register routes: %s", err.Error())
		return err
	}

	s.log.Info("Starting server on port" + s.cfg.Port)

	err = s.router.Run(s.cfg.Port)
	if err != nil {
		s.log.Errorf("Can not start the server: %s", err.Error())
		return err
	}

	return nil
}

// TODO: Write gracefull shutdown
func (s *server) Shutdown() {
	os.Exit(1)
}
