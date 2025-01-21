package server

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"reservation-service/internal/dal"
	"reservation-service/internal/db"
	"reservation-service/internal/handler"
	"reservation-service/internal/service"
)

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
		return err
	}

	return http.ListenAndServe(s.cfg.Port, s.mux)
}

func (s *server) Shutdown() {
	os.Exit(0)
}

func (s *server) registerRoutes() error {
	database, err := db.ConnectMongo(s.cfg.DBuri, s.cfg.DBname)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Registering routes..")

	repository := dal.NewReservationRepository(database)
	service := service.NewReservationService(repository)
	s.handler = handler.NewReservationHandler(service)

	s.mux.HandleFunc("/booking", s.handler.AddReservation)
	s.mux.HandleFunc("/booking/", s.handler.PayReservation)
	s.mux.HandleFunc("/booking/delete/", s.handler.DeleteReservation)

	return nil
}
