package handler

import (
	"reservation-service/reservation-service/internal/service"
)

type ReservationHandler interface{}

type reservationHandler struct {
	reservationService service.ReservationService
}

func NewReservationHandler(s *service.ReservationService) ReservationHandler {
	return &reservationHandler{
		reservationService: s,
	}
}
