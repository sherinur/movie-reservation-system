package server

import (
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
}

func NewServer(cfg *config) Server {
	return &server{
		mux: http.NewServeMux(),
		cfg: cfg,
	}
}

func (s *server) Start() error {
	return http.ListenAndServe(s.cfg.Port, s.mux)
}

func (s *server) Shutdown() {
	os.Exit(0)
}

func (s *server) registerRoutes() error {

	return nil
}
