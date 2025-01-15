package service

import (
	"reservation-service/reservation-service/internal/dal"
)

type ReservationService interface{}

type reservationService struct {
	reservationRepository dal.ReservationRepository
}

func NewReservationService(r *dal.ReservationRepository) ReservationService {
	return &reservationService{
		reservationRepository: r,
	}
}
