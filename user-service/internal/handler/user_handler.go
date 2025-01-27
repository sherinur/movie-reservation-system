package handler

import (
	"net/http"

	"user-service/internal/models"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/movie-reservation-system/pkg/logging"
)

// TODO: handlers must implement http.Handler interface, not custom interface

// TODO: Add to handler : kitHTTP.NewServer -> middleware, go kit, server options

// TODO: Use kitHTTP to unmarshal binary from req.body: - go kit -> decodeRequest

type UserHandler interface {
	HandleLogin(c *gin.Context)
	HandleRegister(c *gin.Context)
	HandleProfile(c *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(s service.UserService) UserHandler {
	return &userHandler{
		userService: s,
	}
}

var log = logging.GetLogger()

// POST /login => auth and give JWT
func (h *userHandler) HandleLogin(c *gin.Context) {
	var logReq models.LoginRequest

	if err := c.ShouldBindJSON(&logReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body"})
		return
	}

	jwtToken, err := h.userService.Authorize(&logReq)
	if err != nil {
		log.Infof("Failed authentication attempt to the profile with email %s from IP %s, error: %s", logReq.Email, c.ClientIP(), err.Error())
		switch err {
		case service.ErrWrongPassword:
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "Invalid password"})
			return
		case service.ErrNoUser:
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "User not found"})
			return
		default:
			log.Errorf("User authentication error: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal server error"})
			return
		}
	}

	log.Infof("User with email %s logged in from %s", logReq.Email, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

// POST /register => new user registration
func (h *userHandler) HandleRegister(c *gin.Context) {
	var regReq models.RegisterRequest

	if err := c.ShouldBindJSON(&regReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body"})
		return
	}

	_, err := h.userService.Register(&regReq)
	if err != nil {
		log.Infof("Failed registration attempt from IP %s, error: %s", c.ClientIP(), err.Error())
		switch err {
		case service.ErrInvalidPassword:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid password"})
			return
		case service.ErrInvalidUsername:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid username"})
			return
		case service.ErrUserExists:
			c.JSON(http.StatusConflict, gin.H{"error": err.Error(), "message": "User already exists"})
			return
		default:
			log.Errorf("User registration error: %s", err.Error())
			c.JSON(http.StatusConflict, gin.H{"error": err.Error(), "message": "Internal server error"})
			return
		}
	}

	log.Infof("Registered an user with email %s from IP %s", regReq.Email, c.ClientIP())
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// GET /profile => get user profile data (need JWT)
func (h *userHandler) HandleProfile(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error(), "message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, users)
}
