package handler

import (
	"net/http"

	"movie-service/internal/models"
	"movie-service/internal/service"
	"movie-service/utils"

	"github.com/gin-gonic/gin"
)

// TODO: add logger and return statement with status code
type MovieHandler interface {
	HandleAddBatchOfMovie(c *gin.Context)
	HandleAddMovie(c *gin.Context)
	HandleGetAllMovie(c *gin.Context)
	HadleGetMovieById(c *gin.Context)
	HandleUpdateMovieById(c *gin.Context)
	HandleDeleteMovieByID(c *gin.Context)
	HandleDeleteAllMovie(c *gin.Context)
}

type movieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(s service.MovieService) MovieHandler {
	return &movieHandler{
		movieService: s,
	}
}

func (h *movieHandler) HandleAddBatchOfMovie(c *gin.Context) {
	var movies []models.Movie
	err := c.ShouldBindJSON(&movies)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.movieService.AddBatchOfMovie(movies)
	if err != nil {
		if _, clientError := utils.BadRequestMovieErrors[err]; clientError {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_ids": res.InsertedIDs})
}

func (h *movieHandler) HandleAddMovie(c *gin.Context) {
	var movie models.Movie
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertRes, err := h.movieService.AddMovie(movie)
	if err != nil {
		_, clientError := utils.BadRequestMovieErrors[err]
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

// GET /movie/get => get all movies
func (h *movieHandler) HandleGetAllMovie(c *gin.Context) {
	data, err := h.movieService.GetAllMovie()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", data)
}

// PAST /movie/:id
func (h *movieHandler) HadleGetMovieById(c *gin.Context) {
	id := c.Param("id")

	data, err := h.movieService.GetMovieById(id)
	if err != nil {
		if _, clientError := utils.BadRequestMovieErrors[err]; clientError {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Data(http.StatusOK, "application/json", data)
}

// PUT /movie/update/:id => update movie information by id
func (h *movieHandler) HandleUpdateMovieById(c *gin.Context) {
	id := c.Param("id")

	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.movieService.UpdateMovieById(id, &movie)
	if err != nil {
		if _, clientError := utils.BadRequestMovieErrors[err]; clientError {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"matched_count": res.MatchedCount})
}

// DELETE /movie/delete/:id => delete movie
func (h *movieHandler) HandleDeleteMovieByID(c *gin.Context) {
	id := c.Param("id")

	deleteres, err := h.movieService.DeleteMovieById(id)
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

func (h *movieHandler) HandleDeleteAllMovie(c *gin.Context) {
	deleteres, err := h.movieService.DeleteAllMovie()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusNoContent, gin.H{"deleted_count": deleteres.DeletedCount})
}
