package handler

import (
	"encoding/json"
	"net/http"
	"reservation-service/reservation-service/internal/models"
	"reservation-service/reservation-service/internal/service"
)

type ReservationHandler interface {
	HandleBooking(w http.ResponseWriter, r *http.Request)
}

type reservationHandler struct {
	reservationService service.ReservationService
}

func NewReservationHandler(s *service.ReservationService) ReservationHandler {
	return &reservationHandler{
		reservationService: s,
	}
}

func (rh *reservationHandler) HandleBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Only GET and POST method is supported.", http.StatusMethodNotAllowed)
		return
	}

	var tickets models.Ticket

	err := json.NewDecoder(r.Body).Decode(&tickets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
