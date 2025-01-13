package handler

import (
	"encoding/json"
	"fmt"
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

var jwtSecret = []byte("secretkey")

// POST /login => auth and give JWT
func (h *userHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req *models.LoginRequest

	// TODO: Fix decoding the request body
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.userService.Authorize(req)
	if err != nil {
		switch err {
		case service.ErrWrongPassword:
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		case service.ErrNoUser:
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("%s, %s, %s, %s", user.ID, user.Username, user.Email, user.Password)))
}

// POST /register => new user registration
func (h *userHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	defer r.Body.Close()

	var user *models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.userService.Register(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("%v", result.InsertedID)))
}

// GET /profile => get user profile data (need JWT)
func (h *userHandler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	w.WriteHeader(200)
	w.Write([]byte("Profile"))
}
