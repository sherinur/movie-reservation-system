package service

import (
	"movie-service/internal/dal"
	"movie-service/internal/models"
	"movie-service/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type CinemaService interface {
	AddCinema(cinemalist []models.Cinema) (*mongo.InsertManyResult, error)
	GetAllCinema() ([]byte, error)
	UpdateCinemaById(id string, cinema *models.Cinema) (*mongo.UpdateResult, error)
	DeleteCinemaById(id string) (*mongo.DeleteResult, error)
}

type cinemaService struct {
	cinemaRepository dal.CinemaRepository
}

func NewCinemaService(r dal.CinemaRepository) CinemaService {
	return &cinemaService{
		cinemaRepository: r,
	}
}

func (s *cinemaService) AddCinema(cinemalist []models.Cinema) (*mongo.InsertManyResult, error) {
	for _, movie := range cinemalist {
		isempty := utils.IsEmpty(movie)
		if !isempty {
			return nil, ErrBadRequest
		}
	}

	res, err := s.cinemaRepository.AddCinema(cinemalist)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s cinemaService) GetAllCinema() ([]byte, error) {
	data, err := s.cinemaRepository.GetAllCinema()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *cinemaService) UpdateCinemaById(id string, cinema *models.Cinema) (*mongo.UpdateResult, error) {
	isempty := utils.IsEmpty(cinema)
	if !isempty {
		return nil, ErrBadRequest
	}

	res, err := s.cinemaRepository.UpdateCinemaById(id, cinema)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *cinemaService) DeleteCinemaById(id string) (*mongo.DeleteResult, error) {
	if id == "" {
		return nil, ErrInvalidId
	}

	res, err := s.cinemaRepository.DeleteCinemaById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
