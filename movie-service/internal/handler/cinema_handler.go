package handler

import (
	"net/http"

	"movie-service/internal/models"
	"movie-service/internal/service"
	"movie-service/utils"

	"github.com/gin-gonic/gin"
)

type CinemaHandler interface {
	HandleAddCinema(c *gin.Context)
	HandleAddHall(c *gin.Context)
	HandleGetAllCinema(c *gin.Context)
	HandleUpdateCinema(c *gin.Context)
	HandleDeleteCinema(c *gin.Context)
	HandleDeleteAllCinema(c *gin.Context)
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
		case utils.ErrBadRequest:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_id": res.InsertedID})
}

func (h *cinemaHandler) HandleAddHall(c *gin.Context) {
	id := c.Param("id")

	var hall models.Hall
	err := c.ShouldBindJSON(&hall)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateResult, err := h.cinemaService.AddHall(id, hall)
	if err != nil {
		_, clientError := utils.BadRequestMovieErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"updated": updateResult})
}

// GET /cinema/get => get all cinema
func (h *cinemaHandler) HandleGetAllCinema(c *gin.Context) {
	data, err := h.cinemaService.GetAllCinema()
	if err != nil {
		if _, clientError := utils.BadRequestCinemaErrors[err]; clientError {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Data(http.StatusOK, "application/json", data)
}

func (h *cinemaHandler) HadleGetCinemaById(c *gin.Context) {
	id := c.Param("id")

	data, err := h.cinemaService.GetCinemaById(id)
	if err != nil {
		if _, clientError := utils.BadRequestCinemaErrors[err]; clientError {
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
		if _, clientError := utils.BadRequestCinemaErrors[err]; clientError {
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

	deleteres, err := h.cinemaService.DeleteCinemaById(id)
	if err != nil {
		if _, clientError := utils.BadRequestMovieErrors[err]; clientError {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"deleted_count": deleteres.DeletedCount})
}

func (h *cinemaHandler) HandleDeleteAllCinema(c *gin.Context) {
	deleteres, err := h.cinemaService.DeleteAllCinema()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusNoContent, gin.H{"deleted_count": deleteres.DeletedCount})
}
