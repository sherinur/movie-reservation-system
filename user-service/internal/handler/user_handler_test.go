package handler

import (
	"bytes"
	"context"
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

	if err := db.Collection("users").Drop(context.TODO()); err != nil {
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

func TestHandleRegister_TableDriven(t *testing.T) {
	r, err := setupTestRouter()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		regReq string

		expectedStatus  int
		expectedError   string
		expectedMessage string
	}{
		{
			name:            "Empty Body",
			regReq:          ``,
			expectedStatus:  http.StatusBadRequest,
			expectedError:   "EOF",
			expectedMessage: "Invalid request body",
		},
		{
			name:            "Empty JSON",
			regReq:          `{}`,
			expectedStatus:  http.StatusBadRequest,
			expectedError:   "email is not valid",
			expectedMessage: "Invalid email",
		},
		{
			name:            "Invalid Password",
			regReq:          `{"username": "john_doe", "email": "john@example.com", "password": "12345678910"}`,
			expectedStatus:  http.StatusBadRequest,
			expectedError:   "password is not valid",
			expectedMessage: "Invalid password",
		},
		{
			name:            "Invalid Password",
			regReq:          `{"username": "john_doe", "email": "john@example.com", "password": "johndoe678910"}`,
			expectedStatus:  http.StatusBadRequest,
			expectedError:   "password is not valid",
			expectedMessage: "Invalid password",
		},
		{
			name:            "Invalid Password",
			regReq:          `{"username": "john_doe", "email": "john@example.com", "password": "Johndoe678910"}`,
			expectedStatus:  http.StatusBadRequest,
			expectedError:   "password is not valid",
			expectedMessage: "Invalid password",
		},
		{
			name:            "Invalid Email",
			regReq:          `{"username": "john_doe", "email": "invalid-emailexample.com", "password": "X@vier!Pass12"}`,
			expectedStatus:  http.StatusBadRequest,
			expectedError:   "email is not valid",
			expectedMessage: "Invalid email",
		},
		{
			name:            "Create",
			regReq:          `{"username": "john_doe", "email": "john@example.com", "password": "X@vier!Pass12"}`,
			expectedStatus:  http.StatusCreated,
			expectedError:   "",
			expectedMessage: "User created successfully",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			body := bytes.NewReader([]byte(test.regReq))

			req, _ := http.NewRequest("POST", "/users/register", body)
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatus, w.Code)

			var response map[string]string
			err = json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			if test.expectedError != "" {
				assert.Equal(t, test.expectedError, response["error"])
			}

			assert.Equal(t, test.expectedMessage, response["message"])
		})
	}
}

func TestHandleRegister_AlreadyExists(t *testing.T) {
	r, err := setupTestRouter()
	if err != nil {
		t.Fatal(err)
	}

	body := []byte(`{"username": "john_doe", "email": "john@example.com", "password": "X@vier!Pass12" }`)

	// the first request
	w := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/users/register", bytes.NewReader(body))
	req1.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req1)

	assert.Equal(t, http.StatusCreated, w.Code)

	// the second request
	w = httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", "/users/register", bytes.NewReader(body))
	req2.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req2)

	assert.Equal(t, http.StatusConflict, w.Code)

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "the user is already exists", response["error"])
	assert.Equal(t, "User already exists", response["message"])
}

func TestHandleLogin_TableDriven(t *testing.T) {
	r, err := setupTestRouter()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		logReq string

		expectedStatus  int
		expectedError   string
		expectedMessage string
	}{
		{
			name:            "Empty Body",
			logReq:          ``,
			expectedStatus:  http.StatusBadRequest,
			expectedError:   "invalid body",
			expectedMessage: "Invalid request body",
		},
		{
			name:            "Empty JSON",
			logReq:          `{}`,
			expectedStatus:  http.StatusUnauthorized,
			expectedError:   "user not found",
			expectedMessage: "User not found",
		},
		{
			name:            "No User",
			logReq:          `{}`,
			expectedStatus:  http.StatusUnauthorized,
			expectedError:   "user not found",
			expectedMessage: "User not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			body := bytes.NewReader([]byte(test.logReq))

			req, _ := http.NewRequest("POST", "/users/login", body)
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatus, w.Code)

			var response map[string]string
			err = json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			if test.expectedError != "" {
				assert.Equal(t, test.expectedError, response["error"])
			}

			assert.Equal(t, test.expectedMessage, response["message"])
		})
	}
}
