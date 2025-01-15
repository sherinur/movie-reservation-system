package handler

import (
	"fmt"
	"net/http"

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
	w.WriteHeader(200)
	w.Write([]byte("Login"))
}

// POST /register => new user registration
func (h *userHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	// var user *models.User

	result, err := h.userService.Register()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(fmt.Sprintf("%v", result.InsertedID)))
}

// GET /profile => get user profile data (need JWT)
func (h *userHandler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Profile"))
}
