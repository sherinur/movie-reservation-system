package server

import (
	"os"

	"movie-service/internal/dal"
	"movie-service/internal/handler"
	"movie-service/internal/service"

	"github.com/Depado/ginprom"
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
	cfg    *Config
	log    *logging.Logger

	movieHandler   handler.MovieHandler
	cinemaHandler  handler.CinemaHandler
	sessionHandler handler.SessionHandler
}

func NewServer(cfg *Config) Server {
	r := gin.Default()

	// Prometheus middleware
	prom := ginprom.New(
		ginprom.Engine(r),
		ginprom.Subsystem("movie_service"),
	)
	r.Use(prom.Instrument())

	corsConfig := &middleware.CorsConfig{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}

	// CORS middleware
	middleware.SetCorsConfig(corsConfig)
	r.Use(middleware.CorsMiddleware())

	// JWT middleware
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

	s.log.Info("Starting server on port " + s.cfg.Port)

	err = s.router.Run(s.cfg.Port)
	if err != nil {
		s.log.Errorf("Cannot start the server: %s", err.Error())
		return err
	}

	return nil
}

func (s *server) Shutdown() {
	os.Exit(1)
}

func (s *server) registerRoutes() error {
	s.log.Info(s.cfg.DbName, s.cfg.DbUri)
	db, err := db.ConnectMongo(s.cfg.DbUri, s.cfg.DbName)
	if err != nil {
		return err
	}

	// Register routes
	movieRepository := dal.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository)
	s.movieHandler = handler.NewMovieHandler(movieService, s.log)

	cinemaRepository := dal.NewCinemaRepository(db)
	cinemaService := service.NewCinemaService(cinemaRepository)
	s.cinemaHandler = handler.NewCinemaHandler(cinemaService, s.log)

	sessionRepository := dal.NewSessionRepository(db, cinemaRepository)
	sessionService := service.NewSessionService(sessionRepository)
	s.sessionHandler = handler.NewSessionHandler(sessionService, s.log)

	// Basic CRUD operation routes for movie and cinema
	s.router.POST("/movie", s.movieHandler.HandleAddMovie)
	s.router.POST("/movie/batch", s.movieHandler.HandleAddBatchOfMovie)
	s.router.GET("/movie", s.movieHandler.HandleGetAllMovie)
	s.router.GET("/movie/:id", s.movieHandler.HadleGetMovieById)
	s.router.PUT("/movie/:id", s.movieHandler.HandleUpdateMovieById)
	s.router.DELETE("/movie/:id", s.movieHandler.HandleDeleteMovieByID)
	s.router.DELETE("/movie", s.movieHandler.HandleDeleteAllMovie)

	s.router.POST("/cinema", s.cinemaHandler.HandleAddCinema)
	s.router.GET("/cinema", s.cinemaHandler.HandleGetAllCinema)
	s.router.GET("/cinema/:id", s.cinemaHandler.HadleGetCinemaById)
	s.router.PUT("/cinema/:id", s.cinemaHandler.HandleUpdateCinema)
	s.router.DELETE("/cinema/:id", s.cinemaHandler.HandleDeleteCinema)
	s.router.DELETE("/cinema", s.cinemaHandler.HandleDeleteAllCinema)

	s.router.POST("/cinema/:id/hall", s.cinemaHandler.HandleAddHall)
	s.router.GET("/cinema/:id/hall_list", s.cinemaHandler.HandleGetAllHall)
	s.router.GET("/cinema/:id/hall/:hallNumber", s.cinemaHandler.HandleGetHall)
	s.router.PUT("/cinema/:id/hall/:hallNumber", s.cinemaHandler.HandleUpdateHall)
	s.router.DELETE("/cinema/:id/hall/:hallNumber", s.cinemaHandler.HandleDeleteHall)

	s.router.POST("/session", s.sessionHandler.HandleAddSession)
	s.router.GET("/session", s.sessionHandler.HandleGetAllSession)
	s.router.GET("/session/:id", s.sessionHandler.HandleGetSessionByID)
	s.router.PUT("/session/:id", s.sessionHandler.HandleUpdateSessionByID)
	s.router.DELETE("/session/:id", s.sessionHandler.HandleDeleteSessionByID)
	s.router.DELETE("/session", s.sessionHandler.HandleDeleteAllSession)

	s.router.GET("/session/:id/seat", s.sessionHandler.HandleGetSeats)
	s.router.GET("/session/movie/:id", s.sessionHandler.HandleGetSessionsByMovieID)
	s.router.POST("session/:id/close", s.sessionHandler.HandlePostSeatClose)
	s.router.POST("session/:id/open", s.sessionHandler.HandlePostSeatClose)

	// Health check route
	s.router.GET("/health", handler.GetHealth)

	return nil
}
