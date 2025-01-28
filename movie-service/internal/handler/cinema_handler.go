package handler

import (
	"net/http"

	"movie-service/internal/models"
	"movie-service/internal/service"

	"github.com/gin-gonic/gin"
)

type CinemaHandler interface {
	HandleAddCinema(c *gin.Context)
	HandleGetAllCinema(c *gin.Context)
	HandleUpdateCinema(c *gin.Context)
	HandleDeleteCinema(c *gin.Context)
}

type cinemaHandler struct {
	cinemaService service.CinemaService
}

func NewCinemaHandler(s service.CinemaService) CinemaHandler {
	return &cinemaHandler{
		cinemaService: s,
	}
}

// POST /cinema/add => add new cinema
func (h *cinemaHandler) HandleAddCinema(c *gin.Context) {
	var cinema models.Cinema
	if err := c.ShouldBindJSON(&cinema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.cinemaService.AddCinema(cinema)
	if err != nil {
		switch err {
		case service.ErrBadRequest:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_id": res.InsertedID})
}

// GET /cinema/get => get all cinema
func (h *cinemaHandler) HandleGetAllCinema(c *gin.Context) {
	data, err := h.cinemaService.GetAllCinema()
	if err != nil {
		if _, clientError := service.BadRequestCinemaErrors[err]; clientError {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Data(http.StatusOK, "application/json", data)
}

// PUT /cinema/update/:id => update cinema information by id
func (h *cinemaHandler) HandleUpdateCinema(c *gin.Context) {
	var cinema models.Cinema
	if err := c.ShouldBindJSON(&cinema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	res, err := h.cinemaService.UpdateCinemaById(id, &cinema)
	if err != nil {
		if _, clientError := service.BadRequestCinemaErrors[err]; clientError {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"matched_count": res.MatchedCount})
}

// DELETE /cinema/delete/:id => delete cinema
func (h *cinemaHandler) HandleDeleteCinema(c *gin.Context) {
	id := c.Param("id")
	res, err := h.cinemaService.DeleteCinemaById(id)
	if err != nil {
		if _, clientError := service.BadRequestMovieErrors[err]; clientError {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"deleted_count": res.DeletedCount})
}
