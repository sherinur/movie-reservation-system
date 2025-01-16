package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reservation-service/reservation-service/internal/models"
	"reservation-service/reservation-service/internal/service"
	"strings"
)

type ReservationHandler interface {
	HandleBooking(w http.ResponseWriter, r *http.Request)
	AddReservation(w http.ResponseWriter, r *http.Request)
	DeleteReservation(w http.ResponseWriter, r *http.Request)
}

type reservationHandler struct {
	reservationService service.ReservationService
}

func NewReservationHandler(s service.ReservationService) ReservationHandler {
	return &reservationHandler{
		reservationService: s,
	}
}

func (rh *reservationHandler) HandleBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported.", http.StatusMethodNotAllowed)
		return
	}

	var tickets models.Reservation

	err := json.NewDecoder(r.Body).Decode(&tickets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(tickets)
}

func (rh *reservationHandler) AddReservation(w http.ResponseWriter, r *http.Request) {
	var reservation models.Reservation
	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := rh.reservationService.AddReservation(reservation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Reservation added successfully"))
}

func (rh *reservationHandler) DeleteReservation(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/booking/delete/")
	if id == "" {
		http.Error(w, "Missing reservation ID", http.StatusBadRequest)
		return
	}
	err := rh.reservationService.DeleteReservation(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s lol %s", err.Error(), id), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reservation deleted successfully"))
}
