package server

import (
	"net/http"
	"os"

	"movie-service/internal/dal"
	"movie-service/internal/db"
	"movie-service/internal/handler"

	"movie-service/internal/service"

	"github.com/sherinur/movie-reservation-system/pkg/logging"
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
	log.Info("Registering routes...")
	err := s.registerRoutes()
	if err != nil {
		log.Errorf("Could not register routes: %s", err.Error())
		return err
	}

	log.Info("Starting server on port" + s.cfg.Port)

	err = http.ListenAndServe(s.cfg.Port, s.mux)
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

// opentelemetry/otel
// load balancer ++ nginx

func (s *server) registerRoutes() error {
	db, err := db.ConnectMongo(s.cfg.DbUri, s.cfg.DbName)
	if err != nil {
		return err
	}

	// Registr routes
	movieRepository := dal.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository)
	s.movieHandler = handler.NewMovieHandler(movieService)

	cinemaRepository := dal.NewCinemaRepository(db)
	cinemaService := service.NewCinemaService(cinemaRepository)
	s.cinemaHandler = handler.NewCinemaHandler(cinemaService)

	// Basic crud operation routes for movie and cinema
	// TODO: test routes and add validation in service
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
