package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"reservation-service/internal/models"
	"reservation-service/internal/service"
	"reservation-service/internal/utilits"

	"github.com/sherinur/movie-reservation-system/pkg/logging"
)

var log = logging.GetLogger()

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
		utilits.WriteErrorResponse(http.StatusMethodNotAllowed, "method not allowed", ErrMethodNotPost, w, r)
		log.Warn()
		return
	}

	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		utilits.WriteErrorResponse(http.StatusBadRequest, "invalid request body", ErrEmptyData, w, r)
		return
	}
	result, err := rh.reservationService.AddReservation(booking)
	if err != nil {
		log.Info("error adding new process: " + err.Error())
		utilits.WriteErrorResponse(http.StatusInternalServerError, "reserving error", err, w, r)
		return
	}

	jsonResponse, err := utilits.ConvertToJson(result)
	if err != nil {
		log.Info("error while coverting response to json: " + err.Error())
		utilits.WriteErrorResponse(http.StatusInternalServerError, "converting error", err, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (rh *reservationHandler) PayReservation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utilits.WriteErrorResponse(http.StatusMethodNotAllowed, "method not allowed", ErrMethodNotPut, w, r)
		return
	}

	var paying models.Paying
	if err := json.NewDecoder(r.Body).Decode(&paying); err != nil {
		utilits.WriteErrorResponse(http.StatusBadRequest, "invalid request body", ErrEmptyData, w, r)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/booking/")
	if id == "" {
		utilits.WriteErrorResponse(http.StatusBadRequest, "invalid request", ErrNoId, w, r)
		return
	}

	result, err := rh.reservationService.PayReservation(id, paying)
	if err != nil {
		log.Info("error paying the reservation: " + err.Error())
		utilits.WriteErrorResponse(http.StatusInternalServerError, "updating error", err, w, r)
		return
	}

	jsonResponse, err := utilits.ConvertToJson(result)
	if err != nil {
		log.Info("error while coverting response to json: " + err.Error())
		utilits.WriteErrorResponse(http.StatusInternalServerError, "converting error", err, w, r)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(jsonResponse)
}

func (rh *reservationHandler) DeleteReservation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utilits.WriteErrorResponse(http.StatusMethodNotAllowed, "method not allowed", ErrMethodNotDelete, w, r)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/booking/delete/")
	if id == "" {
		utilits.WriteErrorResponse(http.StatusBadRequest, "invalid request", ErrNoId, w, r)
		return
	}
	err := rh.reservationService.DeleteReservation(id)
	if err != nil {
		utilits.WriteErrorResponse(http.StatusInternalServerError, "deleting error", err, w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}
