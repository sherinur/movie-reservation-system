package handler

import (
	"net/http"

	"user-service/internal/models"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/movie-reservation-system/pkg/logging"
)

type UserHandler interface {
	HandleLogin(c *gin.Context)
	HandleRegister(c *gin.Context)
	HandleProfile(c *gin.Context)
	HandleUpdatePassword(c *gin.Context)
	HandleUpdateEmail(c *gin.Context)
	HandleDeleteProfile(c *gin.Context)
}

type userHandler struct {
	userService  service.UserService
	tokenService service.TokenService
	log          *logging.Logger
}

func NewUserHandler(u service.UserService, t service.TokenService, logger *logging.Logger) UserHandler {
	return &userHandler{
		userService:  u,
		tokenService: t,
		log:          logger,
	}
}

// POST /login => auth and give JWT
func (h *userHandler) HandleLogin(c *gin.Context) {
	var logReq models.LoginRequest

	if err := c.ShouldBindJSON(&logReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body", "message": "Invalid request body"})
		return
	}

	user, err := h.userService.Authorize(c.Request.Context(), &logReq)
	if err != nil {
		h.log.Infof("Failed authentication attempt to the profile with email %s from IP %s, error: %s", logReq.Email, c.ClientIP(), err.Error())
		switch err {
		case service.ErrWrongPassword:
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "Invalid password"})
			return
		case service.ErrNoUser:
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "User not found"})
			return
		default:
			h.log.Errorf("User authentication error: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal server error"})
			return
		}
	}

	payload := h.tokenService.CreatePayload(user)
	accessToken, refreshToken, err := h.tokenService.GenerateTokens(payload)
	if err != nil {
		h.log.Errorf("User authentication error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal server error"})
		return
	}

	h.log.Infof("User with email %s logged in from %s", logReq.Email, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}

// POST /register => new user registration
func (h *userHandler) HandleRegister(c *gin.Context) {
	var regReq models.RegisterRequest

	if err := c.ShouldBindJSON(&regReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body"})
		return
	}

	err := h.userService.Register(c.Request.Context(), &regReq)
	if err != nil {
		h.log.Infof("Failed registration attempt from IP %s, error: %s", c.ClientIP(), err.Error())
		switch err {
		case service.ErrInvalidEmail:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid email"})
			return
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
			h.log.Errorf("User registration error: %s", err.Error())
			c.JSON(http.StatusConflict, gin.H{"error": err.Error(), "message": "Internal server error"})
			return
		}
	}

	h.log.Infof("Registered an user with email %s from IP %s", regReq.Email, c.ClientIP())
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// GET /users/me => get user profile data (need JWT)
func (h *userHandler) HandleProfile(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userIdStr, ok := userId.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
	}

	user, err := h.userService.GetUser(c.Request.Context(), userIdStr)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error(), "message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GET /users => get users profile data (need JWT with admin role)
func (h *userHandler) HandleGetUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error(), "message": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// DELETE /users/me => delete user profile (need JWT)
func (h *userHandler) HandleDeleteProfile(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userIdStr, ok := userId.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
	}

	err := h.userService.DeleteUser(c.Request.Context(), userIdStr)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error(), "message": "Internal server error"})
		return
	}

	c.Status(http.StatusNoContent)
}

// PUT /users/me/password => update user password (need JWT)
func (h *userHandler) HandleUpdatePassword(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req models.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body"})
		return
	}

	userIdStr, ok := userId.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
		return
	}

	err := h.userService.UpdatePasswordById(c.Request.Context(), userIdStr, req.Password)
	if err != nil {
		switch err {
		case service.ErrInvalidPassword:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid password"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Internal server error"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

// PUT /users/me/email => update user email (need JWT)
func (h *userHandler) HandleUpdateEmail(c *gin.Context) {
	// userId, exists := c.Get("user_id")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	// 	return
	// }

	// userIdStr, ok := userId.(string)
	// if !ok {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
	// }

	// err := h.userService.DeleteUser(userIdStr)
	// if err != nil {
	// 	c.JSON(http.StatusConflict, gin.H{"error": err.Error(), "message": "Internal server error"})
	// 	return
	// }

	c.Status(http.StatusNoContent)
}
