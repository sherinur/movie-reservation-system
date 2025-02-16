package handler

import (
	"net/http"

	"reservation-service/internal/models"
	"reservation-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/movie-reservation-system/pkg/logging"
)

type PromotionHandler interface {
	GetPromotions(c *gin.Context)
	GetPromotion(c *gin.Context)
	AddPromotion(c *gin.Context)
	UpdatePromotion(c *gin.Context)
	DeletePromotion(c *gin.Context)
}

type promotionHandler struct {
	promotionService service.PromotionService
	log              *logging.Logger
}

func NewPromotionHandler(s service.PromotionService, Log *logging.Logger) PromotionHandler {
	return &promotionHandler{
		promotionService: s,
		log:              Log,
	}
}

// GET /promotion --> returns all promotions of a user
func (rh *promotionHandler) GetPromotions(c *gin.Context) {
	result, err := rh.promotionService.GetPromotions(c.Request.Context())
	if err != nil {
		rh.log.Warnf("Error getting reservations: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error"})
		return
	}

	if result == nil {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		rh.log.Infof("Successfully got %d promotion objects", len(result))
		c.JSON(http.StatusOK, result)
	}
}

// GET /promotion/id --> returns specific promotion
func (rh *promotionHandler) GetPromotion(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		rh.log.Infof("Error getting promotion: %s", ErrNoId.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoId.Error(), "message": "Invalid request"})
		return
	}

	result, err := rh.promotionService.GetPromotion(c.Request.Context(), id)
	if err != nil {
		rh.log.Warnf("Error getting promotion: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal Server Error"})
		return
	}

	rh.log.Info("Successfully returned promotion objects")
	c.JSON(http.StatusOK, result)
}

// AddPromotion to create new promotion
func (rh *promotionHandler) AddPromotion(c *gin.Context) {
	var requestBody models.PromotionRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		rh.log.Infof("Error creating promotion: %s", ErrEmptyData.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrEmptyData.Error(), "message": "Invalid Request Body"})
		return
	}

	result, err := rh.promotionService.AddPromotion(c.Request.Context(), requestBody)
	if err != nil {
		rh.log.Warn("Error adding new process: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Reserving error"})
		return
	}

	rh.log.Info("Successfully created new promotion")
	c.JSON(http.StatusAccepted, result)
}

// UpdatePromotion to update promotion
func (rh *promotionHandler) UpdatePromotion(c *gin.Context) {
	var requestBody models.Promotion
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		rh.log.Warn("Error paying promotion: " + ErrEmptyData.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrEmptyData.Error(), "message": "Invalid Request Body"})
		return
	}

	id := c.Param("id")
	if id == "" {
		rh.log.Infof("Error getting promotion: %s", ErrNoId.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoId.Error(), "message": "Invalid request"})
		return
	}

	result, err := rh.promotionService.UpdatePromotion(c.Request.Context(), id, requestBody)
	if err != nil {
		rh.log.Warn("Error paying the promotion: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Updating error"})
		return
	}

	rh.log.Info("Successfully paid processing and created promotion")
	c.JSON(http.StatusOK, result)
}

// DeletePromotion to delete promotion by id
func (rh *promotionHandler) DeletePromotion(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		rh.log.Infof("Error getting promotion: %s", ErrNoId.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNoId.Error(), "message": "Invalid request"})
		return
	}

	result, err := rh.promotionService.DeletePromotion(c.Request.Context(), id)
	if err != nil {
		rh.log.Warnf("Error getting promotion: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Deleting error"})
		return
	}

	rh.log.Info("Successfully deleted promotion")
	c.JSON(http.StatusOK, result)
}
