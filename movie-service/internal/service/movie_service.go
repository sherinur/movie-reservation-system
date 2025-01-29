package service

import (
	"movie-service/internal/dal"
	"movie-service/internal/models"
	"movie-service/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: Implement validation for service
type MovieService interface {
	AddMovie(movielist []models.Movie) (*mongo.InsertManyResult, error)
	GetAllMovie() ([]byte, error)
	GetMovieById(id string) ([]byte, error)
	UpdateMovieById(id string, movie *models.Movie) (*mongo.UpdateResult, error)
	DeleteMovieById(id string) (*mongo.DeleteResult, error)
	DeleteAllMovie() (*mongo.DeleteResult, error)
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
		err := utils.ValidateMovie(movie)
		if err != nil {
			return nil, err
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

func (s *movieService) GetMovieById(id string) ([]byte, error) {
	if id == "" {
		return nil, utils.ErrInvalidId
	}

	data, err := s.movieRepository.GetMovieById(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *movieService) UpdateMovieById(id string, movie *models.Movie) (*mongo.UpdateResult, error) {
	err := utils.ValidateMovie(*movie)
	if err != nil {
		return nil, err
	}

	res, err := s.movieRepository.UpdateMovieById(id, movie)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *movieService) DeleteMovieById(id string) (*mongo.DeleteResult, error) {
	if id == "" {
		return nil, utils.ErrInvalidId
	}

	deleteres, err := s.movieRepository.DeleteMovieById(id)
	if err != nil {
		return nil, err
	}

	return deleteres, nil
}

func (s *movieService) DeleteAllMovie() (*mongo.DeleteResult, error) {
	deleteres, err := s.movieRepository.DeleteAllMovie()
	if err != nil {
		return nil, err
	}

	return deleteres, nil
}
