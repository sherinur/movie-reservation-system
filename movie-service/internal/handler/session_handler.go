package handler

import (
	"net/http"

	"movie-service/internal/models"
	"movie-service/internal/service"
	"movie-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/movie-reservation-system/pkg/logging"
)

// TODO:Test and debug routers
type SessionHandler interface {
	HandleAddSession(c *gin.Context)
	HandleGetAllSession(c *gin.Context)
	HandleGetSessionByID(c *gin.Context)
	HandleUpdateSessionByID(c *gin.Context)
	HandleDeleteSessionByID(c *gin.Context)
	HandleDeleteAllSession(c *gin.Context)

	HandleGetSeats(c *gin.Context)
	HandleGetSessionsByMovieID(c *gin.Context)
}

type sessionHandler struct {
	sessionHandler service.SessionService
	log            *logging.Logger
}

func NewSessionHandler(s service.SessionService, logger *logging.Logger) SessionHandler {
	return &sessionHandler{
		sessionHandler: s,
		log:            logger,
	}
}

func (h *sessionHandler) HandleAddSession(c *gin.Context) {
	var session models.Session
	err := c.ShouldBindJSON(&session)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertResult, err := h.sessionHandler.AddSession(session)
	if err != nil {
		h.log.Infof("Failed to add session from IP %s,error: %s", c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestSessionErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Session added wit ID %s from IP %s", insertResult.InsertedID, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"inserted_id": insertResult.InsertedID})
}

func (h *sessionHandler) HandleGetAllSession(c *gin.Context) {
	session, err := h.sessionHandler.GetAllSession()
	if err != nil {
		h.log.Infof("Failed to get all session from IP %s, error: %s", c.ClientIP(), err.Error())
		switch {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		}
	}

	h.log.Infof("All session get from IP %s", c.ClientIP())
	c.JSON(http.StatusOK, session)
}

func (h *sessionHandler) HandleGetSessionByID(c *gin.Context) {
	sessionID := c.Param("id")

	session, err := h.sessionHandler.GetSessionByID(sessionID)
	if err != nil {
		h.log.Infof("Failed to get session with ID %s from IP %s, error: %s", sessionID, c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestSessionErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Session retrieved with ID %s from IP %s", sessionID, c.ClientIP())
	c.JSON(http.StatusOK, session)
}

func (h *sessionHandler) HandleGetSeats(c *gin.Context) {
	sessionID := c.Param("id")

	seats, err := h.sessionHandler.GetSeat(sessionID)
	if err != nil {
		h.log.Infof("Failed to get seats with ID %s from IP %s, error: %s", sessionID, c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestSessionErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Seats retrieved with ID %s from IP %s", sessionID, c.ClientIP())
	c.JSON(http.StatusOK, seats)
}

func (h *sessionHandler) HandleUpdateSessionByID(c *gin.Context) {
	sessionID := c.Param("id")

	var session models.Session
	err := c.ShouldBindJSON(&session)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateResult, err := h.sessionHandler.UpdateSessionByID(sessionID, session)
	if err != nil {
		h.log.Infof("Failed to update session with ID %s from IP %s, error: %s", sessionID, c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestSessionErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Session updated with ID %s from IP %s", sessionID, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"matched_count": updateResult.MatchedCount})
}

func (h *sessionHandler) HandleDeleteSessionByID(c *gin.Context) {
	sessionID := c.Param("id")

	deleteResult, err := h.sessionHandler.DeleteSessionByID(sessionID)
	if err != nil {
		h.log.Infof("Failed to delete session with ID %s from IP %s, error: %s", sessionID, c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestSessionErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Session deleted with ID %s from IP %s", sessionID, c.ClientIP())
	c.JSON(http.StatusNoContent, gin.H{"deleted_count": deleteResult.DeletedCount})
}

func (h *sessionHandler) HandleDeleteAllSession(c *gin.Context) {
	deleteResult, err := h.sessionHandler.DeleteAllSession()
	if err != nil {
		h.log.Infof("Failed to delete all session from IP %s, error: %s", c.ClientIP(), err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	h.log.Infof("All session deleted from IP %s", c.ClientIP())
	c.JSON(http.StatusNoContent, gin.H{"deleted_count": deleteResult.DeletedCount})
}

func (h *sessionHandler) HandleGetSessionsByMovieID(c *gin.Context) {
	movieID := c.Param("id")

	sessions, err := h.sessionHandler.GetSessionsByMovieID(movieID)
	if err != nil {
		h.log.Infof("Failed to get sessions with movieID %s from IP %s, error: %s", movieID, c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestSessionErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Sessions retrieved with movieID %s from IP %s", movieID, c.ClientIP())
	c.JSON(http.StatusOK, sessions)
}
