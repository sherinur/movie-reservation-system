package service

import (
	"context"
	"fmt"
	"time"

	"reservation-service/internal/dal"
	"reservation-service/internal/models"
	"reservation-service/internal/utilits"

	"go.mongodb.org/mongo-driver/mongo"
)

type ReservationService interface {
	GetReservations(ctx context.Context, userId string) ([]models.Reservation, error)
	GetReservation(ctx context.Context, id string) (*models.Reservation, error)
	AddReservation(ctx context.Context, booking models.ProcessingRequest) (*mongo.InsertOneResult, error)
	PayReservation(ctx context.Context, id string, paying models.ReservationRequest) (*mongo.UpdateResult, error)
	DeleteReservation(ctx context.Context, id, userID string) error
}

type reservationService struct {
	reservationRepository dal.ReservationRepository
}

func NewReservationService(r dal.ReservationRepository) ReservationService {
	return &reservationService{
		reservationRepository: r,
	}
}

func (s *reservationService) GetReservations(ctx context.Context, userId string) ([]models.Reservation, error) {
	return s.reservationRepository.GetByUserId(ctx, userId)
}

func (s *reservationService) GetReservation(ctx context.Context, id string) (*models.Reservation, error) {
	if id == "" {
		return nil, ErrNoId
	}

	return s.reservationRepository.GetById(ctx, id)
}

func (s *reservationService) AddReservation(ctx context.Context, requestBody models.ProcessingRequest) (*mongo.InsertOneResult, error) {
	if len(requestBody.Tickets) == 0 {
		return nil, ErrEmptyData
	}
	if requestBody.ScreeningID == "" {
		return nil, ErrEmptyData
	}
	for _, ticket := range requestBody.Tickets {
		if ticket.SeatColumn == "" || ticket.SeatRow == "" || ticket.Price <= 0 || ticket.SeatType == "" || ticket.UserType == "" {
			return nil, ErrEmptyData
		}
	}

	process := models.Reservation{
		ScreeningID: requestBody.ScreeningID,
		UserID:      requestBody.UserID,
		Status:      "processing",
		Tickets:     requestBody.Tickets,
		ExpiringAt:  time.Now().Add(10 * time.Minute),
	}

	for _, ticket := range requestBody.Tickets {
		process.TotalPrice += ticket.Price
	}

	result, err := s.reservationRepository.AddReservation(ctx, process)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *reservationService) PayReservation(ctx context.Context, id string, requestBody models.ReservationRequest) (*mongo.UpdateResult, error) {
	if id == "" {
		return nil, ErrNoId
	}

	process, err := s.reservationRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if process.Status == "paid" {
		return nil, ErrPaidReservation
	}
	if process.UserID != requestBody.UserID {
		return nil, ErrWrongUser
	}

	reservation := models.Reservation{
		ScreeningID: process.ScreeningID,
		UserID:      process.UserID,
		Email:       requestBody.Email,
		PhoneNumber: requestBody.PhoneNumber,
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

	result, err := s.reservationRepository.UpdateReservation(ctx, id, reservation)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *reservationService) DeleteReservation(ctx context.Context, id, userID string) error {
	if id == "" {
		return ErrNoId
	}

	reservation, err := s.reservationRepository.GetById(ctx, id)
	if err != nil {
		return err
	}
	if reservation.UserID != userID {
		return ErrWrongUser
	}

	return s.reservationRepository.Delete(ctx, id)
}
