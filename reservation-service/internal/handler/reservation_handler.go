package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"reservation-service/internal/models"
	"reservation-service/internal/service"
)

type ReservationHandler interface {
	AddReservation(w http.ResponseWriter, r *http.Request)
	PayReservation(w http.ResponseWriter, r *http.Request)
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

func (rh *reservationHandler) AddReservation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported.", http.StatusMethodNotAllowed)
		return
	}

	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := rh.reservationService.AddReservation(booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Reservation added successfully"))
}

func (rh *reservationHandler) PayReservation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT method is supported.", http.StatusMethodNotAllowed)
		return
	}

	var paying models.Paying
	if err := json.NewDecoder(r.Body).Decode(&paying); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/booking/")
	if id == "" {
		http.Error(w, "Missing update ID", http.StatusBadRequest)
		return
	}

	err := rh.reservationService.PayReservation(id, paying)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating reservation: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Reservation updated successfully"))
}

func (rh *reservationHandler) DeleteReservation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE method is supported.", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/booking/delete/")
	if id == "" {
		http.Error(w, "Missing reservation ID", http.StatusBadRequest)
		return
	}
	err := rh.reservationService.DeleteReservation(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting reservation: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reservation deleted successfully"))
}
