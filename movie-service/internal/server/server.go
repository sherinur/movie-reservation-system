package server

import (
	"os"

	"movie-service/internal/dal"
	"movie-service/internal/handler"
	"movie-service/internal/service"

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
	cfg    *Config

	movieHandler  handler.MovieHandler
	cinemaHandler handler.CinemaHandler
}

func NewServer(cfg *Config) Server {
	return &server{
		router: gin.Default(),
		cfg:    cfg,
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

	err = s.router.Run(s.cfg.Port)
	if err != nil {
		log.Errorf("Can not start the server: %s", err.Error())
		return err
	}

	return nil
}

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
	s.router.POST("/movie", s.movieHandler.HandleAddMovie)
	s.router.POST("/movie/batch", s.movieHandler.HandleAddBatchOfMovie)
	s.router.GET("/movielist", s.movieHandler.HandleGetAllMovie)
	s.router.GET("/movie/:id", s.movieHandler.HadleGetMovieById)
	s.router.PUT("/movie/update/:id", s.movieHandler.HandleUpdateMovieById)
	s.router.DELETE("/movie/delete/:id", s.movieHandler.HandleDeleteMovieByID)
	s.router.DELETE("/movie/delete", s.movieHandler.HandleDeleteAllMovie)

	s.router.POST("/cinema/add", s.cinemaHandler.HandleAddCinema)
	s.router.GET("/cinemalist", s.cinemaHandler.HandleGetAllCinema)
	s.router.GET("/cinema/:id", s.movieHandler.HadleGetMovieById)
	s.router.PUT("/cinema/update/:id", s.cinemaHandler.HandleUpdateCinema)
	s.router.DELETE("/cinema/delete/:id", s.cinemaHandler.HandleDeleteCinema)
	s.router.DELETE("/cinema/delete", s.cinemaHandler.HandleDeleteAllCinema)

	s.router.POST("/cinema/:id/hall", s.cinemaHandler.HandleAddHall)
	s.router.GET("/cinema/:id/hall_list", s.cinemaHandler.HandleGetAllHall)
	s.router.GET("/cinema/:id/hall/:hallNumber", s.cinemaHandler.HandleGetHall)
	s.router.DELETE("/cinema/:id/hall/:hallNumber", s.cinemaHandler.HandleDeleteHall)

	// other routes
	s.router.GET("/health", handler.GetHealth)

	return nil
}
