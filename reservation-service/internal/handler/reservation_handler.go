package handler

import (
	"net/http"

	"reservation-service/internal/models"
	"reservation-service/internal/service"
	"reservation-service/internal/utilits"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/movie-reservation-system/pkg/logging"
)

var log = logging.GetLogger()

type ReservationHandler interface {
	AddReservation(c *gin.Context)
	PayReservation(c *gin.Context)
	DeleteReservation(c *gin.Context)
}

type reservationHandler struct {
	reservationService service.ReservationService
}

func NewReservationHandler(s service.ReservationService) ReservationHandler {
	return &reservationHandler{
		reservationService: s,
	}
}

func (rh *reservationHandler) AddReservation(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		utilits.WriteErrorResponse(http.StatusBadRequest, "invalid request body", ErrEmptyData, c.Writer, c.Request)
		return
	}
	result, err := rh.reservationService.AddReservation(booking)
	if err != nil {
		log.Info("error adding new process: " + err.Error())
		utilits.WriteErrorResponse(http.StatusInternalServerError, "reserving error", err, c.Writer, c.Request)
		return
	}

	// jsonResponse, err := utilits.ConvertToJson(result)
	// if err != nil {
	// 	log.Info("error while converting response to json: " + err.Error())
	// 	utilits.WriteErrorResponse(http.StatusInternalServerError, "converting error", err, c.Writer, c.Request)
	// 	return
	// }

	c.JSON(http.StatusAccepted, result)
}

func (rh *reservationHandler) PayReservation(c *gin.Context) {
	var paying models.Paying
	if err := c.ShouldBindJSON(&paying); err != nil {
		utilits.WriteErrorResponse(http.StatusBadRequest, "invalid request body", ErrEmptyData, c.Writer, c.Request)
		return
	}

	id := c.Param("id")
	if id == "" {
		utilits.WriteErrorResponse(http.StatusBadRequest, "invalid request", ErrNoId, c.Writer, c.Request)
		return
	}

	result, err := rh.reservationService.PayReservation(id, paying)
	if err != nil {
		log.Info("error paying the reservation: " + err.Error())
		utilits.WriteErrorResponse(http.StatusInternalServerError, "updating error", err, c.Writer, c.Request)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (rh *reservationHandler) DeleteReservation(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utilits.WriteErrorResponse(http.StatusBadRequest, "invalid request", ErrNoId, c.Writer, c.Request)
		return
	}
	err := rh.reservationService.DeleteReservation(id)
	if err != nil {
		utilits.WriteErrorResponse(http.StatusInternalServerError, "deleting error", err, c.Writer, c.Request)
		return
	}

	c.Status(http.StatusNoContent)
}
