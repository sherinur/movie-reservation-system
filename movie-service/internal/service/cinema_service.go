package service

import (
	"strconv"

	"movie-service/internal/dal"
	"movie-service/internal/models"
	"movie-service/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type CinemaService interface {
	AddCinema(cinema models.Cinema) (*mongo.InsertOneResult, error)
	GetAllCinema() ([]byte, error)
	GetCinemaById(id string) ([]byte, error)
	UpdateCinemaById(id string, cinema *models.Cinema) (*mongo.UpdateResult, error)
	DeleteCinemaById(id string) (*mongo.DeleteResult, error)
	DeleteAllCinema() (*mongo.DeleteResult, error)

	AddHall(id string, hall models.Hall) (*mongo.UpdateResult, error)
	GetHall(cinemaID string, hallNumber string) ([]byte, error)
	GetAllHall(cinemaID string) ([]models.Hall, error)
	DeleteHall(cinemaID string, hallNumber string) (*mongo.UpdateResult, error)
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

func (s *cinemaService) AddHall(cinemaID string, hall models.Hall) (*mongo.UpdateResult, error) {
	if cinemaID == "" {
		return nil, utils.ErrInvalidId
	}

	err := utils.ValidateHall(hall)
	if err != nil {
		return nil, err
	}

	_, err = s.cinemaRepository.GetHall(cinemaID, hall.Number)
	if err == nil {
		return nil, utils.ErrHallAlreadyExist
	}

	updateResult, err := s.cinemaRepository.AddHall(cinemaID, hall)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (s *cinemaService) GetHall(cinemaID string, hallNumber string) ([]byte, error) {
	if cinemaID == "" {
		return nil, utils.ErrInvalidId
	}

	if hallNumber == "" {
		return nil, utils.ErrInvalidId
	}

	num, err := strconv.Atoi(hallNumber)
	if err != nil {
		return nil, err
	}

	data, err := s.cinemaRepository.GetHall(cinemaID, num)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *cinemaService) GetAllHall(cinemaID string) ([]models.Hall, error) {
	var halls []models.Hall
	if cinemaID == "" {
		return halls, utils.ErrInvalidId
	}

	halls, err := s.cinemaRepository.GetAllHall(cinemaID)
	if err != nil {
		return halls, err
	}

	return halls, nil
}

func (s *cinemaService) DeleteHall(cinemaID string, hallNumber string) (*mongo.UpdateResult, error) {
	if cinemaID == "" {
		return nil, utils.ErrInvalidId
	}

	if hallNumber == "" {
		return nil, utils.ErrInvalidId
	}

	num, err := strconv.Atoi(hallNumber)
	if err != nil {
		return nil, err
	}

	updateResult, err := s.cinemaRepository.DeleteHall(cinemaID, num)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}
