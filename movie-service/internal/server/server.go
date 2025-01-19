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

	movieHandler  handler.MovieHandler
	cinemaHandler handler.CinemaHandler
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

	//Registr routes
	movieRepository := dal.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository)
	s.movieHandler = handler.NewMovieHandler(movieService)

	cinemaRepository := dal.NewCinemaRepository(db)
	cinemaService := service.NewCinemaService(cinemaRepository)
	s.cinemaHandler = handler.NewCinemaHandler(cinemaService)

	//Basic crud operation routes for movie and cinema
	//TODO: test routes and add validation in service
	s.mux.HandleFunc("/movie/add", s.movieHandler.HandleAddMovie)
	s.mux.HandleFunc("/movie/get", s.movieHandler.HandleGetAllMovie)
	s.mux.HandleFunc("/movie/update/{id}", s.movieHandler.HandleUpdateMovieById)
	s.mux.HandleFunc("/movie/delete/{id}", s.movieHandler.HandleDeleteMovieByID)

	//! this routes not tested
	s.mux.HandleFunc("/cinema/add", s.cinemaHandler.HandleAddCinema)
	s.mux.HandleFunc("/cinema/get", s.cinemaHandler.HandleGetAllCinema)
	s.mux.HandleFunc("/cinema/update/{id}", s.cinemaHandler.HandleUpdateCinema)
	s.mux.HandleFunc("/cinema/delete/{id}", s.cinemaHandler.HandleDeleteCinema)

	// other routes
	s.mux.HandleFunc("/health", handler.GetHealth)

	return nil
}
