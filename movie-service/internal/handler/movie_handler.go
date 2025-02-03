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

var log = logging.GetLogger()

func (h *movieHandler) HandleAddBatchOfMovie(c *gin.Context) {
	var movies []models.Movie
	err := c.ShouldBindJSON(&movies)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertResult, err := h.movieService.AddBatchOfMovie(movies)
	if err != nil {
		log.Infof("Failed to add batch of movie from IP %s, error: %s", c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestMovieErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	log.Infof("Batch of movies added with IDs: %s from IP %s", insertResult.InsertedIDs, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"inserted_ids": insertResult.InsertedIDs})
}

func (h *movieHandler) HandleAddMovie(c *gin.Context) {
	var movie models.Movie
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertResult, err := h.movieService.AddMovie(movie)
	if err != nil {
		log.Infof("Failed to add movie from IP %s, error: %s", c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestMovieErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	log.Infof("Movie added with ID %s, from IP %s", insertResult.InsertedID, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"inserted_id": insertResult.InsertedID})
}

func (h *movieHandler) HandleGetAllMovie(c *gin.Context) {
	data, err := h.movieService.GetAllMovie()
	if err != nil {
		log.Infof("Failed to get all movie from IP %s, error: %s", c.ClientIP(), err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", data)
}

func (h *movieHandler) HadleGetMovieById(c *gin.Context) {
	id := c.Param("id")

	data, err := h.movieService.GetMovieById(id)
	if err != nil {
		log.Infof("Faile to get movie by ID %s from IP %s, error: %s", id, c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestMovieErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Data(http.StatusOK, "application/json", data)
}

func (h *movieHandler) HandleUpdateMovieById(c *gin.Context) {
	id := c.Param("id")

	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deleteResult, err := h.movieService.UpdateMovieById(id, &movie)
	if err != nil {
		log.Infof("Failed to update movie with ID %s from IP %s, error: %s", id, c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestMovieErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	log.Infof("Movie updated with ID %s from IP %s", id, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"matched_count": deleteResult.MatchedCount})

}

func (h *movieHandler) HandleDeleteMovieByID(c *gin.Context) {
	id := c.Param("id")

	deleteResult, err := h.movieService.DeleteMovieById(id)
	if err != nil {
		log.Infof("Failed to delete movie with ID %s from IP %s, error: %s", id, c.ClientIP(), err.Error())
		_, clientError := utils.BadRequestMovieErrors[err]
		switch {
		case clientError:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	log.Infof("Movie delete with ID %s from IP %s", id, c.ClientIP())
	c.JSON(http.StatusNoContent, gin.H{"deleted_count": deleteResult.DeletedCount})
}

func (h *movieHandler) HandleDeleteAllMovie(c *gin.Context) {
	deleteResult, err := h.movieService.DeleteAllMovie()
	if err != nil {
		log.Infof("Failed to delete all movies from IP %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	log.Infof("%s movie deleted from IP %s", strconv.Itoa(int(deleteResult.DeletedCount)), c.ClientIP())
	c.JSON(http.StatusNoContent, gin.H{"deleted_count": deleteResult.DeletedCount})
}
