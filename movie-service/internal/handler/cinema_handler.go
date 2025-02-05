package handler

import (
	"net/http"
	"strconv"

	"movie-service/internal/models"
	"movie-service/internal/service"
	"movie-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/movie-reservation-system/pkg/logging"
)

type CinemaHandler interface {
	HandleAddCinema(c *gin.Context)
	HandleGetAllCinema(c *gin.Context)
	HadleGetCinemaById(c *gin.Context)
	HandleUpdateCinema(c *gin.Context)
	HandleDeleteCinema(c *gin.Context)
	HandleDeleteAllCinema(c *gin.Context)

	HandleAddHall(c *gin.Context)
	HandleGetHall(c *gin.Context)
	HandleGetAllHall(c *gin.Context)
	HandleDeleteHall(c *gin.Context)
}

type cinemaHandler struct {
	cinemaService service.CinemaService
	log           *logging.Logger
}

func NewCinemaHandler(s service.CinemaService, logger *logging.Logger) CinemaHandler {
	return &cinemaHandler{
		cinemaService: s,
		log:           logger,
	}
}

func (h *cinemaHandler) HandleAddCinema(c *gin.Context) {
	var cinema models.Cinema

	if err := c.ShouldBindJSON(&cinema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertResult, err := h.cinemaService.AddCinema(cinema)
	if err != nil {
		h.log.Infof("Failed to add cinema from IP %s , error: %s", c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestCinemaErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Cinema add with ID %s from IP %s", insertResult.InsertedID, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"inserted_id": insertResult.InsertedID})
}

func (h *cinemaHandler) HandleGetAllCinema(c *gin.Context) {
	data, err := h.cinemaService.GetAllCinema()
	if err != nil {
		h.log.Infof("Failed to get cinema from IP %s, error: %s", c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestCinemaErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Cinema get from IP %s ", c.ClientIP())
	c.Data(http.StatusOK, "application/json", data)
}

func (h *cinemaHandler) HadleGetCinemaById(c *gin.Context) {
	id := c.Param("id")

	data, err := h.cinemaService.GetCinemaById(id)
	if err != nil {
		h.log.Infof("Failed to get cinema with ID %s from IP %s", id, c.ClientIP())
		_, clientError := utils.BadRequestCinemaErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Cinema get with ID %s from IP %s", id, c.ClientIP())
	c.Data(http.StatusOK, "application/json", data)
}

func (h *cinemaHandler) HandleUpdateCinema(c *gin.Context) {
	var cinema models.Cinema
	if err := c.ShouldBindJSON(&cinema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	updateResult, err := h.cinemaService.UpdateCinemaById(id, &cinema)
	if err != nil {
		h.log.Infof("Failed to update cinema with ID %s from IP %s, error: %s", id, c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestCinemaErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Cinema updated with ID %s from IP %s", id, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"matched_count": updateResult.MatchedCount})
}

func (h *cinemaHandler) HandleDeleteCinema(c *gin.Context) {
	id := c.Param("id")

	deleteResult, err := h.cinemaService.DeleteCinemaById(id)
	if err != nil {
		h.log.Infof("Failed to delete cinema with ID %s from IP %s", id, c.ClientIP())
		_, clientError := utils.BadRequestMovieErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Cinema delete with ID %s from IP %s", id, c.ClientIP())
	c.JSON(http.StatusNoContent, gin.H{"deleted_count": deleteResult.DeletedCount})
}

func (h *cinemaHandler) HandleDeleteAllCinema(c *gin.Context) {
	deleteResult, err := h.cinemaService.DeleteAllCinema()
	if err != nil {
		h.log.Infof("Failed to delete all cinema from IP %s, error: %s", c.ClientIP(), err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	h.log.Infof("%s deleted from IP %s", strconv.Itoa(int(deleteResult.DeletedCount)), c.ClientIP())
	c.JSON(http.StatusNoContent, gin.H{"deleted_count": deleteResult.DeletedCount})
}

// HandleAddHall handles the addition of a new hall to an existing cinema.
// It retrieves the cinema ID from the URL parameters and binds the JSON payload to a Hall model.
// If the JSON binding fails, it returns a 400 Bad Request error.
// It then calls the cinemaService to add the hall to the specified cinema.
// If the cinemaService returns an error, it checks if it's a client error and returns a 400 Bad Request error, otherwise, it returns a 500 Internal Server Error.
// On success, it returns a 200 OK status with the update result.
func (h *cinemaHandler) HandleAddHall(c *gin.Context) {
	cinemaID := c.Param("id")

	var hall models.Hall
	err := c.ShouldBindJSON(&hall)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateResult, err := h.cinemaService.AddHall(cinemaID, hall)
	if err != nil {
		h.log.Infof("Failed to add hall in cinema with ID %s from IP %s, error: %s", cinemaID, c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestCinemaErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Hall added in cinema with ID %s from IP %s", cinemaID, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"updated": updateResult})
}

// HandleGetHall handles the retrieval of a specific hall from a cinema.
// It retrieves the cinema ID and hall number from the URL parameters.
// It then calls the cinemaService to get the hall data.
// If the cinemaService returns an error, it checks if it's a client error and returns a 400 Bad Request error, otherwise, it returns a 500 Internal Server Error.
// On success, it returns a 200 OK status with the hall data in JSON format.
func (h *cinemaHandler) HandleGetHall(c *gin.Context) {
	cinemaID := c.Param("id")
	hallNumber := c.Param("hallNumber")

	data, err := h.cinemaService.GetHall(cinemaID, hallNumber)
	if err != nil {
		h.log.Infof("Failed to get hall %s from cinema with ID %s, error: %s", hallNumber, cinemaID, c.ClientIP())
		_, clientError := utils.BadRequestCinemaErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Hall %s get from cinema with ID %s form IP %s", hallNumber, cinemaID, c.ClientIP())
	c.Data(http.StatusOK, "application/json", data)
}

func (h *cinemaHandler) HandleGetAllHall(c *gin.Context) {
	cinemaID := c.Param("id")

	halls, err := h.cinemaService.GetAllHall(cinemaID)

	if err != nil {
		h.log.Infof("Failed to get all hall from cinema wit ID %s form IP %s,error: %s", cinemaID, c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestCinemaErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// c.Data(http.StatusOK, "application/json", halls)
	h.log.Infof("All hall get from cinema with ID %s from IP %s", cinemaID, c.ClientIP())
	c.JSON(http.StatusOK, halls)
}

// HandleDeleteHall handles the deletion of a specific hall from a cinema.
// It retrieves the cinema ID and hall number from the URL parameters.
// It then calls the cinemaService to delete the hall from the specified cinema.
// If the cinemaService returns an error, it checks if it's a client error and returns a 400 Bad Request error, otherwise, it returns a 500 Internal Server Error.
// On success, it returns a 204 No Content status with the count of deleted halls.
func (h *cinemaHandler) HandleDeleteHall(c *gin.Context) {
	cinemaID := c.Param("id")
	hallNumber := c.Param("hallNumber")

	updateResult, err := h.cinemaService.DeleteHall(cinemaID, hallNumber)
	if err != nil {
		h.log.Infof("Failed to delete hall %s in cinema wiht ID %s from IP %s, error: %s", hallNumber, cinemaID, c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestCinemaErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	h.log.Infof("Hall %s deleted in cinema with ID %s form IP %s", hallNumber, cinemaID, c.ClientIP())
	c.JSON(http.StatusNoContent, gin.H{"deleted_count": updateResult.ModifiedCount})
}
