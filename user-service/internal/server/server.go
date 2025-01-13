package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"user-service/internal/dal"
	"user-service/internal/db"
	"user-service/internal/handler"
	"user-service/internal/service"
)

type Server interface {
	Start() error
	Shutdown()
	registerRoutes() error
}

type server struct {
	mux *http.ServeMux
	cfg *config

	userHandler handler.UserHandler
}

func NewServer(cfg *config) Server {
	return &server{
		mux: http.NewServeMux(),
		cfg: cfg,
	}
}

func (s *server) Start() error {
	// TODO: log the start
	slog.Info("Registering routes ...")
	err := s.registerRoutes()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	slog.Info(fmt.Sprintf("Starting server on port %s", s.cfg.Port))
	return http.ListenAndServe(s.cfg.Port, s.mux)
}

// TODO: Write gracefull shutdown
func (s *server) Shutdown() {
	os.Exit(1)
}

func (s *server) registerRoutes() error {
	db, err := db.ConnectMongo(s.cfg.DbUri, s.cfg.DbName)
	if err != nil {
		return err
	}

	// user routes
	userRepository := dal.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	s.userHandler = handler.NewUserHandler(userService)

	s.mux.HandleFunc("/login", s.userHandler.HandleLogin)
	s.mux.HandleFunc("/register", s.userHandler.HandleRegister)
	s.mux.HandleFunc("/profile", s.userHandler.HandleProfile)

	// other routes
	s.mux.HandleFunc("/health", handler.GetHealth)

	return nil
}
