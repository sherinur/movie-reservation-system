package server

import (
	"net/http"
	"os"

	"user-service/internal/dal"
	"user-service/internal/handler"
	"user-service/internal/service"

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
	mux *http.ServeMux
	cfg *Config

	userHandler handler.UserHandler
}

func NewServer(cfg *Config) Server {
	return &server{
		mux: http.NewServeMux(),
		cfg: cfg,
	}
}

func (s *server) registerRoutes() error {
	db, err := db.ConnectMongo(s.cfg.DbUri, s.cfg.DbName)
	if err != nil {
		return err
	}

	// user routes
	userRepository := dal.NewUserRepository(db)
	userService := service.NewUserService(userRepository, s.cfg.JwtSecretKey)
	s.userHandler = handler.NewUserHandler(userService)

	s.mux.HandleFunc("/login", s.userHandler.HandleLogin)
	s.mux.HandleFunc("/register", s.userHandler.HandleRegister)
	s.mux.Handle("/profile", handler.JwtMiddleware(s.cfg.JwtSecretKey)(http.HandlerFunc(s.userHandler.HandleProfile)))

	// other routes
	s.mux.HandleFunc("/health", handler.GetHealth)

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

	err = http.ListenAndServe(s.cfg.Port, handler.CorsMiddleware(s.mux))
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
