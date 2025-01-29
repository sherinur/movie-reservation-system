package service

import (
	"movie-service/internal/dal"
	"movie-service/internal/models"
	"movie-service/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type CinemaService interface {
	AddCinema(cinema models.Cinema) (*mongo.InsertOneResult, error)
	AddHall(id string, hall models.Hall) (*mongo.UpdateResult, error)
	GetAllCinema() ([]byte, error)
	GetCinemaById(id string) ([]byte, error)
	UpdateCinemaById(id string, cinema *models.Cinema) (*mongo.UpdateResult, error)
	DeleteCinemaById(id string) (*mongo.DeleteResult, error)
	DeleteAllCinema() (*mongo.DeleteResult, error)
}

type cinemaService struct {
	cinemaRepository dal.CinemaRepository
}

func NewCinemaService(r dal.CinemaRepository) CinemaService {
	return &cinemaService{
		cinemaRepository: r,
	}
}

func (s *cinemaService) AddCinema(cinema models.Cinema) (*mongo.InsertOneResult, error) {
	err := utils.ValidateCinema(cinema)
	if err != nil {
		return nil, err
	}

	res, err := s.cinemaRepository.AddCinema(cinema)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *cinemaService) AddHall(id string, hall models.Hall) (*mongo.UpdateResult, error) {
	if id == "" {
		return nil, utils.ErrInvalidId
	}

	err := utils.ValidateHall(hall)
	if err != nil {
		return nil, err
	}

	updateResult, err := s.cinemaRepository.AddHall(id, hall)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (s cinemaService) GetAllCinema() ([]byte, error) {
	data, err := s.cinemaRepository.GetAllCinema()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s cinemaService) GetCinemaById(id string) ([]byte, error) {
	if id == "" {
		return nil, utils.ErrInvalidId
	}

	data, err := s.cinemaRepository.GetCinemaById(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *cinemaService) UpdateCinemaById(id string, cinema *models.Cinema) (*mongo.UpdateResult, error) {
	if id == "" {
		return nil, utils.ErrInvalidId
	}

	err := utils.ValidateCinema(*cinema)
	if err != nil {
		return nil, err
	}

	res, err := s.cinemaRepository.UpdateCinemaById(id, cinema)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *cinemaService) DeleteCinemaById(id string) (*mongo.DeleteResult, error) {
	if id == "" {
		return nil, utils.ErrInvalidId
	}

	deleteres, err := s.cinemaRepository.DeleteCinemaById(id)
	if err != nil {
		return nil, err
	}

	return deleteres, nil
}

func (s *cinemaService) DeleteAllCinema() (*mongo.DeleteResult, error) {
	deleteres, err := s.cinemaRepository.DeleteAllCinema()
	if err != nil {
		return nil, err
	}

	return deleteres, err
}
