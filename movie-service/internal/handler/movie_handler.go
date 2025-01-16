package handler

import (
	"encoding/json"
	"fmt"
	"movie-service/internal/models"
	"movie-service/internal/service"
	"net/http"
)

type MovieHandler interface {
	HandleAddMovie(w http.ResponseWriter, r *http.Request)
	HandleGetMovie(w http.ResponseWriter, r *http.Request)
	HandleUpdateMovie(w http.ResponseWriter, r *http.Request)
	HandleDeleteMovie(w http.ResponseWriter, r *http.Request)
}

type movieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(s service.MovieService) MovieHandler {
	return &movieHandler{
		movieService: s,
	}
}

// Post /addmovie => add new movie
func (h movieHandler) HandleAddMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	defer r.Body.Close()

	var movie []models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.movieService.AddMovie(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", res.InsertedIDs...)))
}

// GET /movies => get all movies
func (h movieHandler) HandleGetMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
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

// PUT /update/{id} => update movie information
func (h movieHandler) HandleUpdateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}

func (h movieHandler) HandleDeleteMovie(w http.ResponseWriter, r *http.Request) {

}
