package handler

import (
	"net/http"

	"reservation-service/internal/models"
	"reservation-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/movie-reservation-system/pkg/logging"
)

type ReservationHandler interface {
	GetReservations(c *gin.Context)
	GetReservation(c *gin.Context)
	AddReservation(c *gin.Context)
	PayReservation(c *gin.Context)
	DeleteReservation(c *gin.Context)
}

type reservationHandler struct {
	reservationService service.ReservationService
	log                *logging.Logger
}

func NewReservationHandler(s service.ReservationService, Log *logging.Logger) ReservationHandler {
	return &reservationHandler{
		reservationService: s,
		log:                Log,
	}
}

// GET /booking --> returns all reservation of a user
func (rh *reservationHandler) GetReservations(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		rh.log.Warnf("Error getting reservations: %s", ErrNotAutorized.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": ErrNotAutorized.Error(), "message": "Not Autorized"})
		return
	}

	result, err := rh.reservationService.GetReservations(c.Request.Context(), userId.(string))
	if err != nil {
		rh.log.Warnf("Error getting reservations: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error"})
		return
	}

	if result == nil {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		rh.log.Infof("Successfully got %d reservation objects", len(result))
		c.JSON(http.StatusOK, result)
	}
}

// GET /booking/id --> returns specific reservation
func (rh *reservationHandler) GetReservation(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		rh.log.Infof("Error getting reservation: %s", ErrNoId.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoId.Error(), "message": "Invalid request"})
		return
	}

	result, err := rh.reservationService.GetReservation(c.Request.Context(), id)
	if err != nil {
		rh.log.Warnf("Error getting reservation: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error"})
		return
	}

	rh.log.Info("Successfully returned reservation objects")
	c.JSON(http.StatusOK, result)
}

// AddReservation to create new processing
func (rh *reservationHandler) AddReservation(c *gin.Context) {
	var requestBody models.ProcessingRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		rh.log.Infof("Error creating processing: %s", ErrEmptyData.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrEmptyData.Error(), "message": "Invalid Request Body"})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		rh.log.Warnf("Error getting reservations: %s", ErrNotAutorized.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": ErrNotAutorized.Error(), "message": "Not Autorized"})
		return
	}

	requestBody.UserID = userId.(string)
	result, err := rh.reservationService.AddReservation(c.Request.Context(), requestBody)
	if err != nil {
		rh.log.Warn("Error adding new process: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Reserving error"})
		return
	}

	rh.log.Info("Successfully created new processing")
	c.JSON(http.StatusAccepted, result)
}

// PayReservation to update processing and make it reservation by id
func (rh *reservationHandler) PayReservation(c *gin.Context) {
	var requestBody models.ReservationRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		rh.log.Warn("Error paying reservation: " + ErrEmptyData.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrEmptyData.Error(), "message": "Invalid Request Body"})
		return
	}

	id := c.Param("id")
	if id == "" {
		rh.log.Infof("Error getting reservation: %s", ErrNoId.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoId.Error(), "message": "Invalid request"})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		rh.log.Warnf("Error getting reservations: %s", ErrNotAutorized.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": ErrNotAutorized.Error(), "message": "Not Autorized"})
		return
	}

	requestBody.UserID = userId.(string)
	result, err := rh.reservationService.PayReservation(c.Request.Context(), id, requestBody)
	if err != nil {
		rh.log.Warn("Error paying the reservation: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Updating error"})
		return
	}

	rh.log.Info("Successfully paid processing and created reservation")
	c.JSON(http.StatusOK, result)
}

// DeleteReservation to delete reservation by id
func (rh *reservationHandler) DeleteReservation(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		rh.log.Infof("Error getting reservation: %s", ErrNoId.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoId.Error(), "message": "Invalid request"})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		rh.log.Warnf("Error getting reservations: %s", ErrNotAutorized.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": ErrNotAutorized.Error(), "message": "Not Autorized"})
		return
	}

	result, err := rh.reservationService.DeleteReservation(c.Request.Context(), id, userId.(string))
	if err != nil {
		rh.log.Warnf("Error getting reservation: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Deleting error"})
		return
	}

	rh.log.Info("Successfully deleted reservation")
	c.JSON(http.StatusOK, result)
}
