package server

import (
	"os"

	"reservation-service/internal/dal"
	"reservation-service/internal/handler"
	"reservation-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/movie-reservation-system/pkg/db"
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
	cfg    *config
	log    *logging.Logger

	handler handler.ReservationHandler
}

func NewServer(cfg *config) Server {
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
	middleware.SetSecret([]byte(cfg.SecretKey))
	return &server{
		router: r,
		cfg:    cfg,
		log:    logging.NewLogger("dev"),
	}
}

func (s *server) Start() error {
	err := s.registerRoutes()
	if err != nil {
		s.log.Errorf("Could not register routes: %s", err.Error())
	}

	s.log.Info("Sarting server at the port" + s.cfg.Port)

	err = s.router.Run(s.cfg.Port)
	if err != nil {
		s.log.Errorf("Error starting server: %s", err.Error())
	}

	return nil
}

func (s *server) Shutdown() {
	os.Exit(0)
}

func (s *server) registerRoutes() error {
	database, err := db.ConnectMongo(s.cfg.DBuri, s.cfg.DBname)
	if err != nil {
		return err
	}

	s.log.Info("Registering routes..")

	repository := dal.NewReservationRepository(database)
	service := service.NewReservationService(repository)
	s.handler = handler.NewReservationHandler(service, s.log)

	autorized := s.router.Group("/booking")
	autorized.Use(middleware.JwtMiddleware())
	{
		autorized.POST("/", s.handler.AddReservation)
		autorized.GET("/", s.handler.GetReservations)
		autorized.GET("/:id", s.handler.GetReservation)
		autorized.PUT("/:id", s.handler.PayReservation)
		autorized.DELETE("/delete/:id", s.handler.DeleteReservation)
	}

	return nil
}
