package server

import (
	"fmt"
	"log/slog"
	"movie-service/internal/dal"
	"movie-service/internal/db"
	"movie-service/internal/handler"
	"movie-service/internal/service"
	"net/http"
	"os"
)

type Server interface {
	Start() error
	Shutdown()
	registerRoutes() error
}

type server struct {
	mux *http.ServeMux
	cfg *config

	movieHandler handler.MovieHandler
}

func NewServer(cfg *config) Server {
	return &server{
		mux: http.NewServeMux(),
		cfg: cfg,
	}
}

func (s *server) Start() error {
	slog.Info("Registering routes ...")
	err := s.registerRoutes()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	slog.Info(fmt.Sprintf("Starting server on port %s", s.cfg.Port))
	return http.ListenAndServe(s.cfg.Port, s.mux)
}
func (s *server) Shutdown() {
	os.Exit(0)
}

func (s *server) registerRoutes() error {
	db, err := db.ConnectMongo(s.cfg.DbUri, s.cfg.DbName)
	if err != nil {
		return err
	}

	movieRepository := dal.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository)
	s.movieHandler = handler.NewMovieHandler(movieService)

	s.mux.HandleFunc("/add", s.movieHandler.HandleAddMovie)
	// s.mux.HandleFunc("/get", s.movieHandler.)
	// s.mux.HandleFunc("/update/{id}" , s.movieHandler.)
	// s.mux.HandleFunc("/delete/{id}", s.movieHandler.)

	return nil
}
