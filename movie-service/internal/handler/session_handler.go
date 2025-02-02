package handler

import (
	"movie-service/internal/models"
	"movie-service/internal/service"
	"movie-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionHandler interface{}

type sessionHandler struct {
	sessionHandler service.SessionService
}

func NewSessionHandler(s service.SessionService) SessionHandler {
	return &sessionHandler{
		sessionHandler: s,
	}
}

func (h *sessionHandler) HandleAddSession(c *gin.Context) {
	var session models.Session
	err := c.ShouldBindJSON(session)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertRes, err := h.sessionHandler.AddSession(session)
	if err != nil {
		_, clientError := utils.BadRequestSessionErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_id": insertRes.InsertedID})
}
