package handler

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"movie-service/internal/dal"
// 	"movie-service/internal/models"
// 	"movie-service/internal/service"

// 	"github.com/gin-gonic/gin"
// 	"github.com/sherinur/movie-reservation-system/pkg/db"
// 	"github.com/sherinur/movie-reservation-system/pkg/logging"
// 	"github.com/stretchr/testify/assert"
// )

// func setupSessionTestRouter() (*gin.Engine, error) {
// 	gin.SetMode(gin.TestMode)
// 	r := gin.New()

// 	db, err := db.ConnectMongo("mongodb://localhost:27017", "sessiontest")
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := db.Collection("sessions").Drop(context.TODO()); err != nil {
// 		return nil, err
// 	}

// 	repo := dal.NewSessionRepository(db)
// 	service := service.NewSessionService(repo)
// 	handler := NewSessionHandler(service, logging.NewLogger("test"))

// 	r.POST("/session", handler.HandleAddSession)
// 	r.GET("/sessions", handler.HandleGetAllSession)
// 	r.PUT("/session/:id", handler.HandleUpdateSessionByID)
// 	r.DELETE("/session/:id", handler.HandleDeleteSessionByID)
// 	r.DELETE("/sessions", handler.HandleDeleteAllSession)

// 	return r, nil
// }

// func TestHandleAddSession(t *testing.T) {
// 	r, err := setupSessionTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	body := []byte(`
// 		"movie_id":        "60d5ec49f9a1c72d4c8e4b8b",
// 		"cinema_id":       "60d5ec49f9a1c72d4c8e4b8c",
// 		"hall_number":     1,
// 		"start_time":      "2023-10-10T14:00:00Z",
// 		"end_time":        "2023-10-10T16:00:00Z",
// 		"available_seats": 50`)
// 	req, _ := http.NewRequest("POST", "/session", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, response["inserted_id"])
// }

// func TestHandleGetAllSession(t *testing.T) {
// 	r, err := setupSessionTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req, _ := http.NewRequest("GET", "/sessions", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response []models.Session
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, response)
// }

// func TestHandleGetSessionByID(t *testing.T) {
// 	r, err := setupSessionTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	session := models.Session{
// 		MovieID:        "60d5ec49f9a1c72d4c8e4b8b",
// 		CinemaID:       "60d5ec49f9a1c72d4c8e4b8c",
// 		HallNumber:     1,
// 		AvailableSeats: 50,
// 	}

// 	body, _ := json.Marshal(session)
// 	req, _ := http.NewRequest("POST", "/session", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	insertedID := response["inserted_id"].(string)

// 	req, _ = http.NewRequest("GET", "/session/"+insertedID, nil)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var sessionResponse models.Session
// 	err = json.Unmarshal(w.Body.Bytes(), &sessionResponse)
// 	assert.NoError(t, err)
// 	assert.Equal(t, session.MovieID, sessionResponse.MovieID)
// }

// func TestHandleUpdateSessionByID(t *testing.T) {
// 	r, err := setupSessionTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	body := []byte(`
// 		"movie_id":        "60d5ec49f9a1c72d4c8e4b8b",
// 		"cinema_id":       "60d5ec49f9a1c72d4c8e4b8c",
// 		"hall_number":     1,
// 		"start_time":      "2023-10-10T14:00:00Z",
// 		"end_time":        "2023-10-10T16:00:00Z",
// 		"available_seats": 50`)

// 	req, _ := http.NewRequest("POST", "/session", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	insertedID := response["inserted_id"].(string)

// 	updatedBody := []byte(`
// 	"movie_id":        "60d5ec49f9a1c72d4c8e4b8b",
// 	"cinema_id":       "60d5ec49f9a1c72d4c8e4b8c",
// 	"hall_number":     1,
// 	"start_time":      "2023-10-10T14:00:00Z",
// 	"end_time":        "2023-10-10T16:00:00Z",
// 	"available_seats": 50`)

// 	updatedSession := models.Session{
// 		AvailableSeats: 50,
// 	}

// 	req, _ = http.NewRequest("PUT", "/session/"+insertedID, bytes.NewReader(updatedBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	req, _ = http.NewRequest("GET", "/session/"+insertedID, nil)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var sessionResponse models.Session
// 	err = json.Unmarshal(w.Body.Bytes(), &sessionResponse)
// 	assert.NoError(t, err)
// 	assert.Equal(t, updatedSession.AvailableSeats, sessionResponse.AvailableSeats)
// }

// func TestHandleDeleteSessionByID(t *testing.T) {
// 	r, err := setupSessionTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	body := []byte(`
// 	"movie_id":        "60d5ec49f9a1c72d4c8e4b8b",
// 	"cinema_id":       "60d5ec49f9a1c72d4c8e4b8c",
// 	"hall_number":     1,
// 	"start_time":      "2023-10-10T14:00:00Z",
// 	"end_time":        "2023-10-10T16:00:00Z",
// 	"available_seats": 50`)

// 	req, _ := http.NewRequest("POST", "/session", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	insertedID := response["inserted_id"].(string)

// 	req, _ = http.NewRequest("DELETE", "/session/"+insertedID, nil)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNoContent, w.Code)
// }

// func TestHandleDeleteAllSession(t *testing.T) {
// 	r, err := setupSessionTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req, _ := http.NewRequest("DELETE", "/sessions", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNoContent, w.Code)
// }
