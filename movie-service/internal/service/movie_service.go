package service

import "movie-service/internal/dal"

type MovieService interface {
}

type movieService struct {
	movieRepository dal.MovieRepository
}

func NewMovieService(r *dal.MovieRepository) MovieService {
	return &movieService{
		movieRepository: r,
	}
}
