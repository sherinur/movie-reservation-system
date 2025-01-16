package service

import (
	"errors"
	"reservation-service/reservation-service/internal/dal"
	"reservation-service/reservation-service/internal/models"
)

type ReservationService interface {
	AddReservation(reservation models.Reservation) error
	DeleteReservation(id string) error
}

type reservationService struct {
	reservationRepository dal.ReservationRepository
}

func NewReservationService(r dal.ReservationRepository) ReservationService {
	return &reservationService{
		reservationRepository: r,
	}
}

func (s *reservationService) AddReservation(reservation models.Reservation) error {
	if reservation.MovieTitle == "" || reservation.Email == "" {
		return errors.New("reservation id or name or email is empty")
	}

	return s.reservationRepository.Add(reservation)
}

func (s *reservationService) DeleteReservation(id string) error {
	if id == "" {
		return errors.New("reservation id is empty")
	}

	return s.reservationRepository.Delete(id)
}
