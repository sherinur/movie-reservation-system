package service

import (
	"errors"
	"fmt"
	"time"

	"reservation-service/internal/dal"
	"reservation-service/internal/models"
	"reservation-service/internal/utilits"
)

type ReservationService interface {
	AddReservation(booking models.Booking) error
	PayReservation(id string, paying models.Paying) error
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
	if len(booking.Tickets) == 0 {
		return errors.New("booking is empty")
	}
	for _, ticket := range booking.Tickets {
		if ticket.SeatColumn == "" || ticket.SeatRow == "" || ticket.Price < 0 || ticket.Type == "" {
			return errors.New("not provided all seat data")
		}
	}

	process := models.Process{
		ScreeningID: booking.ScreeningID,
		Status:      "processing",
		Tickets:     booking.Tickets,
		CreatedTime: time.Now().String(),
	}

	for _, ticket := range booking.Tickets {
		process.TotalPrice += ticket.Price
	}

	return s.reservationRepository.Add(process)
}

func (s *reservationService) PayReservation(id string, paying models.Paying) error {
	if id == "" {
		return errors.New("id is empty")
	}

	process, err := s.reservationRepository.GetById(id)
	if err != nil {
		return err
	}

	reservation := models.Reservation{
		ScreeningID: process.ScreeningID,
		Email:       paying.Email,
		PhoneNumber: paying.Email,
		Status:      "paid",
		Tickets:     process.Tickets,
		TotalPrice:  process.TotalPrice,
		BoughtTime:  time.Now().String(),
	}

	qrData := fmt.Sprintf("Reservation for %d seats at %s\nStatus: %s", len(reservation.Tickets), reservation.BoughtTime, reservation.Status)
	QR, err := utilits.GenerateQR(qrData)
	if err != nil {
		return err
	}
	reservation.QRCode = QR

	return s.reservationRepository.Update(id, reservation)
}

func (s *reservationService) DeleteReservation(id string) error {
	if id == "" {
		return errors.New("id is empty")
	}

	return s.reservationRepository.Delete(id)
}
