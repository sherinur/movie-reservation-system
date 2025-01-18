package service

import (
	"errors"
	"fmt"

	"reservation-service/internal/dal"
	"reservation-service/internal/models"
	"reservation-service/internal/utilits"
	"time"
)

type ReservationService interface {
	AddReservation(booking models.Booking) error
	PayReservation(id string) error
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
		Status:     "Processing",
		BoughtTime: time.Now().String(),
		Tickets:    booking.Tickets,
	}

	for _, ticket := range reservation.Tickets {
		reservation.TotalPrice += ticket.Price
	}

	qrData := fmt.Sprintf("Reservation for %d seats on %s at %s\nStatus: %s", len(reservation.Tickets), reservation.MovieTitle, reservation.BoughtTime, reservation.Status)
	QR, err := utilits.GenerateQR(qrData)
	if err != nil {
		return err
	}
	reservation.QRCode = QR

	return s.reservationRepository.Add(reservation)
}

func (s *reservationService) PayReservation(id string) error {
	if id == "" {
		return errors.New("id is empty")
	}

	reservation, err := s.reservationRepository.GetById(id)
	if err != nil {
		return err
	}

	err = utilits.SendMail(reservation.Email, "Good boooy", reservation.MovieTitle)

	return s.reservationRepository.Update(id)
}

func (s *reservationService) DeleteReservation(id string) error {
	if id == "" {
		return errors.New("reservation id is empty")
	}

	return s.reservationRepository.Delete(id)
}
