package server

import (
	"net/http"
	"os"

	"user-service/internal/dal"
	"user-service/internal/db"
	"user-service/internal/handler"
	"user-service/internal/service"
	"user-service/internal/utils"
)

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

var Logger = utils.NewLogger(true, true)

func (s *server) Start() error {
	Logger.PrintInfoMsg("Registering routes...")
	err := s.registerRoutes()
	if err != nil {
		Logger.PrintErrorMsg("Could not register routes: " + err.Error())
		return err
	}

	Logger.PrintInfoMsg("Starting server on port " + s.cfg.Port)

	err = http.ListenAndServe(s.cfg.Port, s.mux)
	if err != nil {
		Logger.PrintErrorMsg("Can not start the server: " + err.Error())
		return err
	}

	return nil
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
	userService := service.NewUserService(userRepository, s.cfg.JwtSecretKey)
	s.userHandler = handler.NewUserHandler(userService)

	s.mux.HandleFunc("/login", s.userHandler.HandleLogin)
	s.mux.HandleFunc("/register", s.userHandler.HandleRegister)
	s.mux.HandleFunc("/profile", s.userHandler.HandleProfile)

	// other routes
	s.mux.HandleFunc("/health", handler.GetHealth)

	return nil
}
