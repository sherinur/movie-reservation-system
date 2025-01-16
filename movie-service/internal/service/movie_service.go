package service

import (
	"movie-service/internal/dal"
	"movie-service/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type MovieService interface {
	AddMovie(movielist []models.Movie) (*mongo.InsertManyResult, error)
	GetAllMovie() ([]byte, error)
	UpdateMovie()
	DeleteMovie()
}

type movieService struct {
	movieRepository dal.MovieRepository
}

func NewMovieService(r dal.MovieRepository) MovieService {
	return &movieService{
		movieRepository: r,
	}
}

func (s *movieService) AddMovie(movielist []models.Movie) (*mongo.InsertManyResult, error) {

	for _, movie := range movielist {
		//TODO write checking for empty values
		if movie.Title == "" || movie.Description == "" || movie.Genre == "" {
			return nil, ErrBadRequest
		}
	}

	res, err := s.movieRepository.AddMovie(movielist)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *movieService) GetAllMovie() ([]byte, error) {
	data, err := s.movieRepository.GetAllMovie()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *movieService) UpdateMovie() {

}

func (s *movieService) DeleteMovie() {

}
