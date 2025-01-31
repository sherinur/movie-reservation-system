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
	GetReservations(c *gin.Context)
	GetReservation(c *gin.Context)
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

// GetReservations to handle getting all existing reservations
func (rh *reservationHandler) GetReservations(c *gin.Context) {
	result, err := rh.reservationService.GetReservations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GerReservation to handle getting reservation by id
func (rh *reservationHandler) GetReservation(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoId.Error(), "message": "Invalid request"})
		return
	}

	result, err := rh.reservationService.GetReservation(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error"})
	}

	c.JSON(http.StatusOK, result)
}

// AddReservation to create new processing
func (rh *reservationHandler) AddReservation(c *gin.Context) {
	var requestBody models.ProcessingRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utilits.WriteErrorResponse(http.StatusBadRequest, "invalid request body", ErrEmptyData, c.Writer, c.Request)
		return
	}
	result, err := rh.reservationService.AddReservation(requestBody)
	if err != nil {
		log.Info("error adding new process: " + err.Error())
		utilits.WriteErrorResponse(http.StatusInternalServerError, "reserving error", err, c.Writer, c.Request)
		return
	}

	c.JSON(http.StatusAccepted, result)
}

// PayReservation to update processing and make it reservation by id
func (rh *reservationHandler) PayReservation(c *gin.Context) {
	var requestBody models.ReservationRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utilits.WriteErrorResponse(http.StatusBadRequest, "invalid request body", ErrEmptyData, c.Writer, c.Request)
		return
	}

	id := c.Param("id")
	if id == "" {
		utilits.WriteErrorResponse(http.StatusBadRequest, "invalid request", ErrNoId, c.Writer, c.Request)
		return
	}

	result, err := rh.reservationService.PayReservation(id, requestBody)
	if err != nil {
		log.Info("error paying the reservation: " + err.Error())
		utilits.WriteErrorResponse(http.StatusInternalServerError, "updating error", err, c.Writer, c.Request)
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteReservation to delete reservation by id
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
