package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/skip2/go-qrcode"
	"reservation-service/internal/dal"
	"reservation-service/internal/models"
	"time"
)

type ReservationService interface {
	AddReservation(booking models.Booking) error
	DeleteReservation(id string) error
	UpdateStatus(id string) error
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

	qrData := fmt.Sprintf("Reservation for %d seats on %s at %s\nStatus: %s", len(reservation.Tickets), reservation.MovieTitle, reservation.BoughtTime, reservation.Status)

	var png []byte
	png, err := qrcode.Encode(qrData, qrcode.Medium, 256)
	if err != nil {
		return err
	}

	reservation.QRCode = base64.StdEncoding.EncodeToString(png)

	return s.reservationRepository.Add(reservation)
}

func (s *reservationService) UpdateStatus(id string) error {
	if id == "" {
		return errors.New("id is empty")
	}

	return s.reservationRepository.Update(id)
}

func (s *reservationService) DeleteReservation(id string) error {
	if id == "" {
		return errors.New("reservation id is empty")
	}

	return s.reservationRepository.Delete(id)
}
