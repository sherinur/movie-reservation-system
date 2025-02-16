package handler

import (
	"net/http"

	"reservation-service/internal/models"
	"reservation-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/movie-reservation-system/pkg/logging"
)

type PaymentHandler interface {
	GetPayments(c *gin.Context)
	GetPayment(c *gin.Context)
	AddPayment(c *gin.Context)
	UpdatePayment(c *gin.Context)
	DeletePayment(c *gin.Context)
}

type paymentHandler struct {
	paymentService service.PaymentService
	log            *logging.Logger
}

func NewPaymentHandler(s service.PaymentService, Log *logging.Logger) PaymentHandler {
	return &paymentHandler{
		paymentService: s,
		log:            Log,
	}
}

func (rh *paymentHandler) GetPayments(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		rh.log.Warnf("Error getting Payments: %s", ErrNotAutorized.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": ErrNotAutorized.Error(), "message": "Not Autorized"})
		return
	}

	result, err := rh.paymentService.GetPayments(c.Request.Context(), userId.(string))
	if err != nil {
		rh.log.Warnf("Error getting Payments: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error"})
		return
	}

	if result == nil {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		rh.log.Infof("Successfully got %d Payment objects", len(result))
		c.JSON(http.StatusOK, result)
	}
}

// GET /booking/id --> returns specific Payment
func (rh *paymentHandler) GetPayment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		rh.log.Infof("Error getting Payment: %s", ErrNoId.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoId.Error(), "message": "Invalid request"})
		return
	}

	result, err := rh.paymentService.GetPayment(c.Request.Context(), id)
	if err != nil {
		rh.log.Warnf("Error getting Payment: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error"})
		return
	}

	rh.log.Info("Successfully returned Payment objects")
	c.JSON(http.StatusOK, result)
}

func (rh *paymentHandler) AddPayment(c *gin.Context) {
	var requestBody models.PaymentRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		rh.log.Infof("Error creating processing: %s", ErrEmptyData.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrEmptyData.Error(), "message": "Invalid Request Body"})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		rh.log.Warnf("Error getting Payments: %s", ErrNotAutorized.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": ErrNotAutorized.Error(), "message": "Not Autorized"})
		return
	}

	requestBody.UserId = userId.(string)
	result, err := rh.paymentService.AddPayment(c.Request.Context(), requestBody)
	if err != nil {
		rh.log.Warn("Error adding new process: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Reserving error"})
		return
	}

	rh.log.Info("Successfully created new processing")
	c.JSON(http.StatusAccepted, result)
}

// PayPayment to update processing and make it Payment by id
func (rh *paymentHandler) UpdatePayment(c *gin.Context) {
	var requestBody models.Payment
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		rh.log.Warn("Error paying Payment: " + ErrEmptyData.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrEmptyData.Error(), "message": "Invalid Request Body"})
		return
	}

	id := c.Param("id")
	if id == "" {
		rh.log.Infof("Error getting Payment: %s", ErrNoId.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoId.Error(), "message": "Invalid request"})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		rh.log.Warnf("Error getting Payments: %s", ErrNotAutorized.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": ErrNotAutorized.Error(), "message": "Not Autorized"})
		return
	}

	requestBody.UserId = userId.(string)
	result, err := rh.paymentService.UpdatePayment(c.Request.Context(), id, requestBody)
	if err != nil {
		rh.log.Warn("Error paying the Payment: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Updating error"})
		return
	}

	rh.log.Info("Successfully paid processing and created Payment")
	c.JSON(http.StatusOK, result)
}

func (rh *paymentHandler) DeletePayment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		rh.log.Infof("Error getting Payment: %s", ErrNoId.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoId.Error(), "message": "Invalid request"})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		rh.log.Warnf("Error getting Payments: %s", ErrNotAutorized.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": ErrNotAutorized.Error(), "message": "Not Autorized"})
		return
	}

	result, err := rh.paymentService.DeletePayment(c.Request.Context(), id, userId.(string))
	if err != nil {
		rh.log.Warnf("Error getting Payment: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Deleting error"})
		return
	}

	rh.log.Info("Successfully deleted Payment")
	c.JSON(http.StatusOK, result)
}
