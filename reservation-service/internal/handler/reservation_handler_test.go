package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reservation-service/internal/dal"
	"reservation-service/internal/models"
	"reservation-service/internal/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sherinur/movie-reservation-system/pkg/db"
	"github.com/sherinur/movie-reservation-system/pkg/logging"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() (*gin.Engine, ReservationHandler, error) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// secret := "a5d52d1471164c78450ee0f6095cfN2f2c712e45525010b0e46e936cc61e6d205"
	// middleware.SetSecret([]byte(secret))
	// r.Use(middleware.JwtMiddleware())

	db, err := db.ConnectMongo("mongodb://localhost:27017", "test")
	if err != nil {
		return nil, nil, err
	}

	err = db.Collection("reservations").Drop(context.TODO())
	if err != nil {
		return nil, nil, err
	}

	repo := dal.NewReservationRepository(db)
	serv := service.NewReservationService(repo)
	handler := NewReservationHandler(serv, logging.NewLogger("test"))

	r.POST("booking", handler.AddReservation)
	r.GET("booking", handler.GetReservations)
	r.GET("booking/:id", handler.GetReservation)
	r.PUT("booking/:id", handler.PayReservation)
	r.DELETE("booking/delete/:id", handler.DeleteReservation)

	return r, handler, nil
}

func TestAddReservation(t *testing.T) {
	_, handler, err := setupTestRouter()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		req    string
		userId string

		expectedStatus  int
		expectedError   string
		expectedMessage string
	}{
		{
			name:            "User not autorized",
			req:             `{ "screening_id": "679bg09j53ae5cc94c021c5d", "tickets": [ { "seat_row": "A", "seat_column": "3", "price": 1500.00, "seat_type": "common", "user_type": "adult" } ] }`,
			userId:          "",
			expectedStatus:  http.StatusUnauthorized,
			expectedError:   "not autorized",
			expectedMessage: "Not Autorized",
		},
		{
			name:            "Not provided all data",
			req:             `{ "tickets": [ { "seat_row": "A", "seat_column": "3", "price": 1500.00, "seat_type": "common", "user_type": "adult" } ] }`,
			userId:          "679dc7d7ebe4f308cb076e7f",
			expectedStatus:  http.StatusBadRequest,
			expectedError:   "not provided all data",
			expectedMessage: "Reserving error",
		},
		{
			name:            "No request data",
			req:             ``,
			userId:          "679dc7d7ebe4f308cb076e7f",
			expectedStatus:  http.StatusBadRequest,
			expectedError:   "incoming data must be entered",
			expectedMessage: "Invalid Request Body",
		},
		// {
		// 	name: "",
		// 	req: `{}`,
		// 	userId: "",
		// 	expectedStatus: ,
		// 	expectedError: ,
		// 	expectedMessage: ,
		// },
		{
			name:            "Successful creating",
			req:             `{ "screening_id": "679bg09j53ae5cc94c021c5d", "tickets": [ { "seat_row": "A", "seat_column": "3", "price": 1500.00, "seat_type": "common", "user_type": "adult" } ] }`,
			userId:          "679dc7d7ebe4f308cb076e7f",
			expectedStatus:  http.StatusAccepted,
			expectedError:   "",
			expectedMessage: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			body := bytes.NewReader([]byte(test.req))
			req, _ := http.NewRequest(http.MethodPost, "/booking", body)

			c, _ := gin.CreateTestContext(w)
			c.Request = req
			if test.userId != "" {
				c.Set("user_id", test.userId)
			}

			handler.AddReservation(c)

			assert.Equal(t, test.expectedStatus, w.Code)

			var response map[string]string
			err = json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			if test.expectedError != "" {
				assert.Equal(t, test.expectedError, response["error"])
			}

			if test.expectedMessage != "" {
				assert.Equal(t, test.expectedMessage, response["message"])
			}
		})
	}
}

func TestGetReservations(t *testing.T) {
	_, handler, err := setupTestRouter()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		userId string

		expectedStatus  int
		expectedError   string
		expectedMessage string
	}{
		{
			name:            "User not autorized",
			userId:          "",
			expectedStatus:  http.StatusUnauthorized,
			expectedError:   "not autorized",
			expectedMessage: "Not Autorized",
		},
		{
			name:            "Successful getting reservations",
			userId:          "679dc7d7ebe4f308cb076e7f",
			expectedStatus:  http.StatusOK,
			expectedError:   "",
			expectedMessage: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodGet, "/booking", nil)

			c, _ := gin.CreateTestContext(w)
			c.Request = req
			if test.userId != "" {
				c.Set("user_id", test.userId)
			}

			handler.GetReservations(c)

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

func TestGetReservation(t *testing.T) {
	r, handler, err := setupTestRouter()
	if err != nil {
		t.Fatal(err)
	}

	body := bytes.NewReader([]byte(`{ "screening_id": "679bg09j53ae5cc94c021c5d", "tickets": [ { "seat_row": "A", "seat_column": "3", "price": 1500.00, "seat_type": "common", "user_type": "adult" } ] }`))

	req, _ := http.NewRequest(http.MethodPost, "/booking", body)
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("user_id", "679dc7d7ebe4f308cb076e7f")

	handler.AddReservation(c)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	insertedID := response["InsertedID"].(string)

	req, _ = http.NewRequest(http.MethodGet, "/booking/"+insertedID, nil)
	w = httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var Response models.Reservation
	err = json.Unmarshal(w.Body.Bytes(), &Response)
	assert.NoError(t, err)
	//assert.Equal(t, newRes.ScreeningID, Response.ScreeningID)
}
