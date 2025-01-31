package server

import (
	"os"

	"reservation-service/internal/dal"
	"reservation-service/internal/handler"
	"reservation-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/movie-reservation-system/pkg/db"
	"github.com/sherinur/movie-reservation-system/pkg/logging"
)

var log = logging.GetLogger()

type Server interface {
	Start() error
	Shutdown()
	registerRoutes() error
}

type server struct {
	router *gin.Engine
	cfg    *config

	handler handler.ReservationHandler
}

func NewServer(cfg *config) Server {
	return &server{
		router: gin.Default(),
		cfg:    cfg,
	}
}

func (s *server) Start() error {
	err := s.registerRoutes()
	if err != nil {
		log.Errorf("Could not register routes: %s", err.Error())
	}

	log.Info("Sarting server at the port" + s.cfg.Port)

	err = s.router.Run(s.cfg.Port)
	if err != nil {
		log.Errorf("Error starting server: %s", err.Error())
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

	log.Info("Registering routes..")

	repository := dal.NewReservationRepository(database)
	service := service.NewReservationService(repository)
	s.handler = handler.NewReservationHandler(service)

	s.router.GET("/booking", s.handler.GetReservations)
	s.router.GET("/booking/:id", s.handler.GetReservation)
	s.router.POST("/booking", s.handler.AddReservation)
	s.router.PUT("/booking/:id", s.handler.PayReservation)
	s.router.DELETE("/booking/delete/:id", s.handler.DeleteReservation)

	return nil
}
