package server

import (
	"net/http"
	"os"

	"reservation-service/internal/dal"
	"reservation-service/internal/handler"
	"reservation-service/internal/service"

	"github.com/sherinur/movie-reservation-system/pkg/logging"

	"github.com/sherinur/movie-reservation-system/pkg/db"
)

var log = logging.GetLogger()

type Server interface {
	Start() error
	Shutdown()
	registerRoutes() error
}

type server struct {
	mux *http.ServeMux
	cfg *config

	handler handler.ReservationHandler
}

func NewServer(cfg *config) Server {
	return &server{
		mux: http.NewServeMux(),
		cfg: cfg,
	}
}

func (s *server) Start() error {
	err := s.registerRoutes()
	if err != nil {
		log.Errorf("Could not register routes: %s", err.Error())
	}

	log.Info("Sarting server at the port" + s.cfg.Port)

	err = http.ListenAndServe(s.cfg.Port, s.mux)
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

	s.mux.HandleFunc("/booking", s.handler.AddReservation)
	s.mux.HandleFunc("/booking/", s.handler.PayReservation)
	s.mux.HandleFunc("/booking/delete/", s.handler.DeleteReservation)

	return nil
}
