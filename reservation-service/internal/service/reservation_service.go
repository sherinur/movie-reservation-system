package service

import (
	"errors"
	"reservation-service/reservation-service/internal/dal"
	"reservation-service/reservation-service/internal/models"
	"time"
)

type ReservationService interface {
	AddReservation(booking models.Booking) error
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

func (s *reservationService) AddReservation(booking models.Booking) error {
	if booking.MovieTitle == "" || booking.Email == "" || len(booking.Tickets) == 0 {
		return errors.New("booking is empty")
	}
	reservation := models.Reservation{
		MovieTitle: booking.MovieTitle,
		Email:      booking.Email,
		Status:     "Bought",
		BoughtTime: time.Now().String(),
		Tickets:    booking.Tickets,
	}
	return s.reservationRepository.Add(reservation)
}

func (s *reservationService) DeleteReservation(id string) error {
	if id == "" {
		return errors.New("reservation id is empty")
	}

	return s.reservationRepository.Delete(id)
}
