package service

import (
	"fmt"
	"time"

	"reservation-service/internal/dal"
	"reservation-service/internal/models"
	"reservation-service/internal/utilits"

	"go.mongodb.org/mongo-driver/mongo"
)

type ReservationService interface {
	AddReservation(booking models.Booking) (*mongo.InsertOneResult, error)
	PayReservation(id string, paying models.Paying) (*mongo.UpdateResult, error)
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

func (s *reservationService) AddReservation(booking models.Booking) (*mongo.InsertOneResult, error) {
	if len(booking.Tickets) == 0 {
		return nil, ErrEmptyData
	}
	if booking.ScreeningID == "" || booking.UserID == "" {
		return nil, ErrEmptyData
	}
	for _, ticket := range booking.Tickets {
		if ticket.SeatColumn == "" || ticket.SeatRow == "" || ticket.Price <= 0 || ticket.SeatType == "" || ticket.UserType == "" {
			return nil, ErrEmptyData
		}
	}

	process := models.Process{
		ScreeningID: booking.ScreeningID,
		UserID:      booking.UserID,
		Status:      "processing",
		Tickets:     booking.Tickets,
		ExpiringAt:  time.Now().Add(20 * time.Second),
	}

	for _, ticket := range booking.Tickets {
		process.TotalPrice += ticket.Price
	}

	result, err := s.reservationRepository.Add(process)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *reservationService) PayReservation(id string, paying models.Paying) (*mongo.UpdateResult, error) {
	if id == "" {
		return nil, ErrNoId
	}

	process, err := s.reservationRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	if process.Status != "processing" {
		return nil, ErrPaidReservation
	}

	reservation := models.Reservation{
		ScreeningID: process.ScreeningID,
		UserID:      process.UserID,
		Email:       paying.Email,
		PhoneNumber: paying.Email,
		Status:      "paid",
		Tickets:     process.Tickets,
		TotalPrice:  process.TotalPrice,
		BoughtTime:  time.Now(),
	}

	qrData := fmt.Sprintf("Reservation for %d seats at %s\nStatus: %s", len(reservation.Tickets), reservation.BoughtTime, reservation.Status)
	QR, err := utilits.GenerateQR(qrData)
	if err != nil {
		return nil, err
	}
	reservation.QRCode = QR

	result, err := s.reservationRepository.Update(id, reservation)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *reservationService) DeleteReservation(id string) error {
	if id == "" {
		return ErrNoId
	}

	return s.reservationRepository.Delete(id)
}
