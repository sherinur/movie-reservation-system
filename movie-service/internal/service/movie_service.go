package service

import (
	"movie-service/internal/dal"
	"movie-service/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type MovieService interface {
	AddMovie(movielist *models.MovieList) (*mongo.InsertManyResult, error)
	GetMovie() (*models.MovieList, error)
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

func (s *movieService) AddMovie(movielist *models.MovieList) (*mongo.InsertManyResult, error) {

	var movies []interface{}

	for _, r := range movielist.List {
		//TODO write checking for empty values
		if r.Title == "" || r.Description == "" || r.Genre == "" {
			return nil, ErrBadRequest
		}
		movies = append(movies, r)
	}

	res, err := s.movieRepository.AddMovie(movies)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *movieService) GetMovie() (*models.MovieList, error) {
	data, err := s.movieRepository.GetMovie()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *movieService) UpdateMovie() {

}

func (s *movieService) DeleteMovie() {

}
