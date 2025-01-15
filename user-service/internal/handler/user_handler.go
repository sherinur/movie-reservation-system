package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"user-service/internal/models"
	"user-service/internal/service"
)

type UserHandler interface {
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandleRegister(w http.ResponseWriter, r *http.Request)
	HandleProfile(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(s service.UserService) UserHandler {
	return &userHandler{
		userService: s,
	}
}

// POST /login => auth and give JWT
func (h *userHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.LoginRequest

	// TODO: Fix decoding the request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	jwtToken, err := h.userService.Authorize(&req)
	if err != nil {
		switch err {
		case service.ErrWrongPassword:
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		case service.ErrNoUser:
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		default:
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(200)
	w.Write([]byte(jwtToken))
}

// POST /register => new user registration
func (h *userHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var req *models.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.userService.Register(req)
	if err != nil {
		switch err {
		case service.ErrInvalidPassword:
			http.Error(w, "Invalid password", http.StatusBadRequest)
			return
		case service.ErrInvalidUsername:
			http.Error(w, "Invalid username", http.StatusBadRequest)
			return
		case service.ErrUserExists:
			http.Error(w, "User already exists", http.StatusConflict)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("%v", result.InsertedID)))
}

// GET /profile => get user profile data (need JWT)
func (h *userHandler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := h.userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&users)
}
