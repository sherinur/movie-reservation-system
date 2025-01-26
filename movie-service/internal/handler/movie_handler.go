package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"movie-service/internal/models"
	"movie-service/internal/service"
)

// TODO: add logger and return statement with status code
type MovieHandler interface {
	HandleAddMovie(w http.ResponseWriter, r *http.Request)
	HandleGetAllMovie(w http.ResponseWriter, r *http.Request)
	HandleUpdateMovieById(w http.ResponseWriter, r *http.Request)
	HandleDeleteMovieByID(w http.ResponseWriter, r *http.Request)
}

type movieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(s service.MovieService) MovieHandler {
	return &movieHandler{
		movieService: s,
	}
}

// Post /movie/add => add new movie
func (h movieHandler) HandleAddMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var movie []models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.movieService.AddMovie(movie)
	if err != nil {
		switch err {
		case service.ErrBadRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", res.InsertedIDs...)))
}

// GET /movie/get => get all movies
func (h movieHandler) HandleGetAllMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := h.movieService.GetAllMovie()
	if err != nil {
		switch err {
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// PUT /movie/update/{id} => update movie information by id
func (h movieHandler) HandleUpdateMovieById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var movie *models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")
	res, err := h.movieService.UpdateMovieById(id, movie)
	if err != nil {
		switch err {
		case service.ErrBadRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", res.MatchedCount)))
}

// DELETE /movie/delete/{id} => delete movie
func (h movieHandler) HandleDeleteMovieByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")

	res, err := h.movieService.DeleteMovieById(id)
	if err != nil {
		switch err {
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("%v", res.DeletedCount)))
}
