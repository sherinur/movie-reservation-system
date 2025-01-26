package service

import (
	"fmt"
	"movie-service/internal/dal"
	"movie-service/internal/models"
	"movie-service/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: Implement validation for service
type MovieService interface {
	AddMovie(movielist []models.Movie) (*mongo.InsertManyResult, error)
	GetAllMovie() ([]byte, error)
	UpdateMovieById(id string, movie *models.Movie) (*mongo.UpdateResult, error)
	DeleteMovieById(id string) (*mongo.DeleteResult, error)
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
		emptyField := utils.IsEmpty(movie)
		if emptyField != "" {
			return nil, fmt.Errorf("empty field: %s", emptyField)
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

func (s *movieService) UpdateMovieById(id string, movie *models.Movie) (*mongo.UpdateResult, error) {
	emptyField := utils.IsEmpty(movie)
	if emptyField != "" {
		return nil, fmt.Errorf("empty field: %s", emptyField)
	}

	res, err := s.movieRepository.UpdateMovieById(id, movie)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *movieService) DeleteMovieById(id string) (*mongo.DeleteResult, error) {
	if id == "" {
		return nil, ErrInvalidId
	}

	res, err := s.movieRepository.DeleteMovieById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
