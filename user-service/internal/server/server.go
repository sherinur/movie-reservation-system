package server

import (
	"fmt"
	"log/slog"
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
	cfg config
}

func NewServer(cfg *config) Server {
	return &server{
		mux: http.NewServeMux(),
		cfg: *cfg,
	}
}

func (s *server) Start() error {
	// TODO: log the start
	slog.Info(fmt.Sprintf("Starting server on port %s", s.cfg.Port))
	return http.ListenAndServe(s.cfg.Port, s.mux)
}

// TODO: Write gracefull shutdown
func (s *server) Shutdown() {
	os.Exit(1)
}

func (s *server) registerRoutes() error {
	return nil
}
