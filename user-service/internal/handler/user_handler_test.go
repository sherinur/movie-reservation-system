package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"user-service/internal/dal"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/movie-reservation-system/pkg/db"
	"github.com/sherinur/movie-reservation-system/pkg/logging"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() (*gin.Engine, error) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	secret := "a5d52d1471164c78450ee0f6095cfN2f2c712e45525010b0e46e936cc61e6d205"

	db, err := db.ConnectMongo("mongodb://localhost:27017", "test")
	if err != nil {
		return nil, err
	}

	repo := dal.NewUserRepository(db)
	service := service.NewUserService(repo, secret)
	handler := NewUserHandler(service, logging.NewLogger("test"))

	r.POST("/users/register", handler.HandleRegister)
	r.POST("/users/login", handler.HandleLogin)
	r.GET("/users/me", handler.HandleProfile)
	r.PUT("/users/me/password", handler.HandleUpdatePassword)
	r.PUT("/users/me/email", handler.HandleUpdatePassword)
	r.DELETE("/users/me", handler.HandleDeleteProfile)

	return r, nil
}

func TestHandleLogin(t *testing.T) {
	r, err := setupTestRouter()
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/login", nil)
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "invalid body", response["error"])
	assert.Equal(t, "Invalid request body", response["message"])
}
