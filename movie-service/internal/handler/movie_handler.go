package handler

import "movie-service/internal/service"

type MovieHandler interface {
}

type movieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(s *service.MovieService) MovieHandler {
	return &movieHandler{
		movieService: s,
	}
}
