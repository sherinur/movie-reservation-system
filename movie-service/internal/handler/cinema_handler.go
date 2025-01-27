package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"movie-service/internal/models"
	"movie-service/internal/service"
)

// TODO: add logger and return statement with status code
type CinemaHandler interface {
	HandleAddCinema(w http.ResponseWriter, r *http.Request)
	HandleGetAllCinema(w http.ResponseWriter, r *http.Request)
	HandleUpdateCinema(w http.ResponseWriter, r *http.Request)
	HandleDeleteCinema(w http.ResponseWriter, r *http.Request)
}

type cinemaHandler struct {
	cinemaService service.CinemaService
}

func NewCinemaHandler(s service.CinemaService) CinemaHandler {
	return &cinemaHandler{
		cinemaService: s,
	}
}

// Post /cinema/add => add new cinema
func (h *cinemaHandler) HandleAddCinema(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported.", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	var cinema models.Cinema

	if err := json.NewDecoder(r.Body).Decode(&cinema); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.cinemaService.AddCinema(cinema)
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
	w.Write([]byte(fmt.Sprintf("%v", res.InsertedID)))
}

// GET /cinema/get => get all cinema
func (h *cinemaHandler) HandleGetAllCinema(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported.", http.StatusMethodNotAllowed)
		return
	}

	data, err := h.cinemaService.GetAllCinema()
	if err != nil {
		_, clientError := service.BadRequestCinemaErrors[err]
		switch {
		case clientError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// PUT /cinema/update/{id} => update cinema information by id
func (h *cinemaHandler) HandleUpdateCinema(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT method is supported.", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var cinema *models.Cinema
	if err := json.NewDecoder(r.Body).Decode(&cinema); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")
	res, err := h.cinemaService.UpdateCinemaById(id, cinema)
	if err != nil {
		_, clientError := service.BadRequestCinemaErrors[err]
		switch {
		case clientError:
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

// DELETE /cinema/delete/{id} => delete cinema
func (h *cinemaHandler) HandleDeleteCinema(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE method is supported.", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	res, err := h.cinemaService.DeleteCinemaById(id)
	if err != nil {
		_, clientError := service.BadRequestMovieErrors[err]
		switch {
		case clientError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("%v", res.DeletedCount)))
}
