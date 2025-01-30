package server

import (
	"os"

	"user-service/internal/dal"
	"user-service/internal/handler"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"

	"github.com/sherinur/movie-reservation-system/pkg/db"
	"github.com/sherinur/movie-reservation-system/pkg/logging"
	"github.com/sherinur/movie-reservation-system/pkg/middleware"
)

var log = logging.GetLogger()

type Server interface {
	Start() error
	Shutdown()
	registerRoutes() error
}

type server struct {
	router *gin.Engine
	cfg    *Config

	userHandler handler.UserHandler
}

func NewServer(cfg *Config) Server {
	r := gin.Default()

	corsConfig := &middleware.CorsConfig{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedMethods: []string{"GET", "POST", "UPDATE", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}

	// cors middleware
	middleware.SetCorsConfig(corsConfig)
	r.Use(middleware.CorsMiddleware())

	// jwt middleware
	middleware.SetSecret([]byte(cfg.JwtSecretKey))

	return &server{
		router: r,
		cfg:    cfg,
	}
}

func (s *server) registerRoutes() error {
	db, err := db.ConnectMongo(s.cfg.DbUri, s.cfg.DbName)
	if err != nil {
		return err
	}

	userRepository := dal.NewUserRepository(db)
	userService := service.NewUserService(userRepository, s.cfg.JwtSecretKey)
	s.userHandler = handler.NewUserHandler(userService)

	s.router.GET("/health", handler.GetHealth)

	s.router.POST("/register", s.userHandler.HandleRegister)
	s.router.POST("/login", s.userHandler.HandleLogin)
	s.router.GET("/users/me", middleware.JwtMiddleware(), s.userHandler.HandleProfile)
	s.router.PUT("/users/me/password", middleware.JwtMiddleware(), s.userHandler.HandleUpdatePassword)
	s.router.PUT("/users/me/email", middleware.JwtMiddleware(), s.userHandler.HandleUpdatePassword)
	s.router.DELETE("/users/me", middleware.JwtMiddleware(), s.userHandler.HandleDeleteProfile)

	return nil
}

func (s *server) Start() error {
	log.Info("Registering routes...")
	err := s.registerRoutes()
	if err != nil {
		log.Errorf("Could not register routes: %s", err.Error())
		return err
	}

	log.Info("Starting server on port" + s.cfg.Port)

	err = s.router.Run(s.cfg.Port)
	if err != nil {
		log.Errorf("Can not start the server: %s", err.Error())
		return err
	}

	return nil
}

// TODO: Write gracefull shutdown
func (s *server) Shutdown() {
	os.Exit(1)
}
