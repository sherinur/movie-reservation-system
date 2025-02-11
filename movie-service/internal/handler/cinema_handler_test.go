package handler

// func setupCinemaTestRouter() (*gin.Engine, error) {
// 	gin.SetMode(gin.TestMode)
// 	r := gin.New()

// 	db, err := db.ConnectMongo("mongodb://localhost:27017", "cinematest")
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := db.Collection("cinemas").Drop(context.TODO()); err != nil {
// 		return nil, err
// 	}

// 	repo := dal.NewCinemaRepository(db)
// 	service := service.NewCinemaService(repo)
// 	handler := NewCinemaHandler(service, logging.NewLogger("test"))

// 	r.POST("/cinema", handler.HandleAddCinema)
// 	r.GET("/cinemalist", handler.HandleGetAllCinema)
// 	r.GET("/cinema/:id", handler.HadleGetCinemaById)
// 	r.PUT("/cinema/:id", handler.HandleUpdateCinema)
// 	r.DELETE("/cinema/:id", handler.HandleDeleteCinema)
// 	r.DELETE("/cinemalist", handler.HandleDeleteAllCinema)

// 	r.POST("/cinema/:id/hall", handler.HandleAddHall)
// 	r.GET("/cinema/:id/hall_list", handler.HandleGetAllHall)
// 	r.GET("/cinema/:id/hall/:hallNumber", handler.HandleGetHall)
// 	r.DELETE("/cinema/:id/hall/:hallNumber", handler.HandleDeleteHall)

// 	return r, nil
// }

// func TestHandleAddCinema(t *testing.T) {
// 	r, err := setupCinemaTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	cinema := models.Cinema{
// 		Name:    "Kinopark 6 Keruencity",
// 		City:    "Astana",
// 		Address: "Ramstore - Mega, Qorghalzhyn Hwy 1, Astana 010000",
// 		Rating:  4.6,
// 		HallList: []models.Hall{
// 			{
// 				Number: 1,
// 				Seats: []models.Seat{
// 					{Row: "A", Column: "1", Status: "available"},
// 					{Row: "A", Column: "2", Status: "available"},
// 					{Row: "A", Column: "3", Status: "available"},
// 					{Row: "A", Column: "4", Status: "available"},
// 					{Row: "A", Column: "5", Status: "available"},
// 					{Row: "A", Column: "6", Status: "available"},
// 					{Row: "A", Column: "7", Status: "available"},
// 					{Row: "A", Column: "8", Status: "available"},
// 					{Row: "A", Column: "9", Status: "available"},
// 					{Row: "A", Column: "10", Status: "available"},
// 				},
// 			},
// 		},
// 	}

// 	body, _ := json.Marshal(cinema)
// 	req, _ := http.NewRequest("POST", "/cinema", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, response["inserted_id"])
// }

// func TestHandleGetAllCinema(t *testing.T) {
// 	r, err := setupCinemaTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req, _ := http.NewRequest("GET", "/cinemalist", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response []models.Cinema
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, response)
// }

// func TestHandleGetCinemaById(t *testing.T) {
// 	r, err := setupCinemaTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	cinema := models.Cinema{
// 		Name:    "Kinopark 6 Keruencity",
// 		City:    "Astana",
// 		Address: "Ramstore - Mega, Qorghalzhyn Hwy 1, Astana 010000",
// 		Rating:  4.6,
// 		HallList: []models.Hall{
// 			{
// 				Number: 1,
// 				Seats: []models.Seat{
// 					{Row: "A", Column: "1", Status: "available"},
// 					{Row: "A", Column: "2", Status: "available"},
// 					{Row: "A", Column: "3", Status: "available"},
// 					{Row: "A", Column: "4", Status: "available"},
// 					{Row: "A", Column: "5", Status: "available"},
// 					{Row: "A", Column: "6", Status: "available"},
// 					{Row: "A", Column: "7", Status: "available"},
// 					{Row: "A", Column: "8", Status: "available"},
// 					{Row: "A", Column: "9", Status: "available"},
// 					{Row: "A", Column: "10", Status: "available"},
// 				},
// 			},
// 		},
// 	}

// 	body, _ := json.Marshal(cinema)
// 	req, _ := http.NewRequest("POST", "/cinema", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	insertedID := response["inserted_id"].(string)

// 	req, _ = http.NewRequest("GET", "/cinema/"+insertedID, nil)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var cinemaResponse models.Cinema
// 	err = json.Unmarshal(w.Body.Bytes(), &cinemaResponse)
// 	assert.NoError(t, err)
// 	assert.Equal(t, cinema.Name, cinemaResponse.Name)
// }

// func TestHandleUpdateCinemaById(t *testing.T) {
// 	r, err := setupCinemaTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	cinema := models.Cinema{
// 		Name:    "Kinopark 6 Keruencity",
// 		City:    "Astana",
// 		Address: "Ramstore - Mega, Qorghalzhyn Hwy 1, Astana 010000",
// 		Rating:  4.6,
// 		HallList: []models.Hall{
// 			{
// 				Number: 1,
// 				Seats: []models.Seat{
// 					{Row: "A", Column: "1", Status: "available"},
// 					{Row: "A", Column: "2", Status: "available"},
// 					{Row: "A", Column: "3", Status: "available"},
// 					{Row: "A", Column: "4", Status: "available"},
// 					{Row: "A", Column: "5", Status: "available"},
// 					{Row: "A", Column: "6", Status: "available"},
// 					{Row: "A", Column: "7", Status: "available"},
// 					{Row: "A", Column: "8", Status: "available"},
// 					{Row: "A", Column: "9", Status: "available"},
// 					{Row: "A", Column: "10", Status: "available"},
// 				},
// 			},
// 		},
// 	}

// 	body, _ := json.Marshal(cinema)
// 	req, _ := http.NewRequest("POST", "/cinema", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	insertedID := response["inserted_id"].(string)

// 	updatedCinema := models.Cinema{
// 		Name:    "Kinopark 7 Keruencity",
// 		City:    "Astana",
// 		Address: "Ramstore - Mega, Qorghalzhyn Hwy 1, Astana 010000",
// 		Rating:  4.7,
// 		HallList: []models.Hall{
// 			{
// 				Number: 1,
// 				Seats: []models.Seat{
// 					{Row: "A", Column: "1", Status: "available"},
// 					{Row: "A", Column: "2", Status: "available"},
// 					{Row: "A", Column: "3", Status: "available"},
// 					{Row: "A", Column: "4", Status: "available"},
// 					{Row: "A", Column: "5", Status: "available"},
// 					{Row: "A", Column: "6", Status: "available"},
// 					{Row: "A", Column: "7", Status: "available"},
// 					{Row: "A", Column: "8", Status: "available"},
// 					{Row: "A", Column: "9", Status: "available"},
// 					{Row: "A", Column: "10", Status: "available"},
// 				},
// 			},
// 		},
// 	}

// 	updatedBody, _ := json.Marshal(updatedCinema)
// 	req, _ = http.NewRequest("PUT", "/cinema/"+insertedID, bytes.NewReader(updatedBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	req, _ = http.NewRequest("GET", "/cinema/"+insertedID, nil)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var cinemaResponse models.Cinema
// 	err = json.Unmarshal(w.Body.Bytes(), &cinemaResponse)
// 	assert.NoError(t, err)
// 	assert.Equal(t, updatedCinema.Name, cinemaResponse.Name)
// }

// func TestHandleDeleteCinemaByID(t *testing.T) {
// 	r, err := setupCinemaTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	cinema := models.Cinema{
// 		Name:    "Kinopark 6 Keruencity",
// 		City:    "Astana",
// 		Address: "Ramstore - Mega, Qorghalzhyn Hwy 1, Astana 010000",
// 		Rating:  4.6,
// 		HallList: []models.Hall{
// 			{
// 				Number: 1,
// 				Seats: []models.Seat{
// 					{Row: "A", Column: "1", Status: "available"},
// 					{Row: "A", Column: "2", Status: "available"},
// 					{Row: "A", Column: "3", Status: "available"},
// 					{Row: "A", Column: "4", Status: "available"},
// 					{Row: "A", Column: "5", Status: "available"},
// 					{Row: "A", Column: "6", Status: "available"},
// 					{Row: "A", Column: "7", Status: "available"},
// 					{Row: "A", Column: "8", Status: "available"},
// 					{Row: "A", Column: "9", Status: "available"},
// 					{Row: "A", Column: "10", Status: "available"},
// 				},
// 			},
// 		},
// 	}

// 	body, _ := json.Marshal(cinema)
// 	req, _ := http.NewRequest("POST", "/cinema", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	insertedID := response["inserted_id"].(string)

// 	req, _ = http.NewRequest("DELETE", "/cinema/"+insertedID, nil)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNoContent, w.Code)
// }

// func TestHandleDeleteAllCinema(t *testing.T) {
// 	r, err := setupCinemaTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req, _ := http.NewRequest("DELETE", "/cinemalist", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNoContent, w.Code)
// }

// func TestHandleAddHall(t *testing.T) {
// 	r, err := setupCinemaTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	cinema := models.Cinema{
// 		Name:    "Kinopark 6 Keruencity",
// 		City:    "Astana",
// 		Address: "Ramstore - Mega, Qorghalzhyn Hwy 1, Astana 010000",
// 		Rating:  4.6,
// 		HallList: []models.Hall{
// 			{
// 				Number: 1,
// 				Seats: []models.Seat{
// 					{Row: "A", Column: "1", Status: "available"},
// 					{Row: "A", Column: "2", Status: "available"},
// 					{Row: "A", Column: "3", Status: "available"},
// 					{Row: "A", Column: "4", Status: "available"},
// 					{Row: "A", Column: "5", Status: "available"},
// 					{Row: "A", Column: "6", Status: "available"},
// 					{Row: "A", Column: "7", Status: "available"},
// 					{Row: "A", Column: "8", Status: "available"},
// 					{Row: "A", Column: "9", Status: "available"},
// 					{Row: "A", Column: "10", Status: "available"},
// 				},
// 			},
// 		},
// 	}

// 	body, _ := json.Marshal(cinema)
// 	req, _ := http.NewRequest("POST", "/cinema", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	insertedID := response["inserted_id"].(string)

// 	hall := models.Hall{
// 		Number: 2,
// 		Seats: []models.Seat{
// 			{Row: "A", Column: "1", Status: "available"},
// 			{Row: "A", Column: "2", Status: "available"},
// 			{Row: "A", Column: "3", Status: "available"},
// 			{Row: "A", Column: "4", Status: "available"},
// 			{Row: "A", Column: "5", Status: "available"},
// 			{Row: "A", Column: "6", Status: "available"},
// 			{Row: "A", Column: "7", Status: "available"},
// 			{Row: "A", Column: "8", Status: "available"},
// 			{Row: "A", Column: "9", Status: "available"},
// 			{Row: "A", Column: "10", Status: "available"},
// 		},
// 	}

// 	hallBody, _ := json.Marshal(hall)
// 	req, _ = http.NewRequest("POST", "/cinema/"+insertedID+"/hall", bytes.NewReader(hallBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var hallResponse map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &hallResponse)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, hallResponse["updated"])
// }

// func TestHandleGetAllHall(t *testing.T) {
// 	r, err := setupCinemaTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	cinema := models.Cinema{
// 		Name:    "Kinopark 6 Keruencity",
// 		City:    "Astana",
// 		Address: "Ramstore - Mega, Qorghalzhyn Hwy 1, Astana 010000",
// 		Rating:  4.6,
// 		HallList: []models.Hall{
// 			{
// 				Number: 1,
// 				Seats: []models.Seat{
// 					{Row: "A", Column: "1", Status: "available"},
// 					{Row: "A", Column: "2", Status: "available"},
// 					{Row: "A", Column: "3", Status: "available"},
// 					{Row: "A", Column: "4", Status: "available"},
// 					{Row: "A", Column: "5", Status: "available"},
// 					{Row: "A", Column: "6", Status: "available"},
// 					{Row: "A", Column: "7", Status: "available"},
// 					{Row: "A", Column: "8", Status: "available"},
// 					{Row: "A", Column: "9", Status: "available"},
// 					{Row: "A", Column: "10", Status: "available"},
// 				},
// 			},
// 		},
// 	}

// 	body, _ := json.Marshal(cinema)
// 	req, _ := http.NewRequest("POST", "/cinema", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	insertedID := response["inserted_id"].(string)

// 	req, _ = http.NewRequest("GET", "/cinema/"+insertedID+"/hall_list", nil)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var hallsResponse []models.Hall_list
// 	err = json.Unmarshal(w.Body.Bytes(), &hallsResponse)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, hallsResponse)
// }

// func TestHandleGetHall(t *testing.T) {
// 	r, err := setupCinemaTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	cinema := models.Cinema{
// 		Name:    "Kinopark 6 Keruencity",
// 		City:    "Astana",
// 		Address: "Ramstore - Mega, Qorghalzhyn Hwy 1, Astana 010000",
// 		Rating:  4.6,
// 		HallList: []models.Hall{
// 			{
// 				Number: 1,
// 				Seats: []models.Seat{
// 					{Row: "A", Column: "1", Status: "available"},
// 					{Row: "A", Column: "2", Status: "available"},
// 					{Row: "A", Column: "3", Status: "available"},
// 					{Row: "A", Column: "4", Status: "available"},
// 					{Row: "A", Column: "5", Status: "available"},
// 					{Row: "A", Column: "6", Status: "available"},
// 					{Row: "A", Column: "7", Status: "available"},
// 					{Row: "A", Column: "8", Status: "available"},
// 					{Row: "A", Column: "9", Status: "available"},
// 					{Row: "A", Column: "10", Status: "available"},
// 				},
// 			},
// 		},
// 	}

// 	body, _ := json.Marshal(cinema)
// 	req, _ := http.NewRequest("POST", "/cinema", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	insertedID := response["inserted_id"].(string)

// 	req, _ = http.NewRequest("GET", "/cinema/"+insertedID+"/hall/1", nil)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var hallResponse models.Hall
// 	err = json.Unmarshal(w.Body.Bytes(), &hallResponse)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 1, hallResponse.Number)
// }

// func TestHandleDeleteHall(t *testing.T) {
// 	r, err := setupCinemaTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	cinema := models.Cinema{
// 		Name:    "Kinopark 6 Keruencity",
// 		City:    "Astana",
// 		Address: "Ramstore - Mega, Qorghalzhyn Hwy 1, Astana 010000",
// 		Rating:  4.6,
// 		HallList: []models.Hall{
// 			{
// 				Number: 1,
// 				Seats: []models.Seat{
// 					{Row: "A", Column: "1", Status: "available"},
// 					{Row: "A", Column: "2", Status: "available"},
// 					{Row: "A", Column: "3", Status: "available"},
// 					{Row: "A", Column: "4", Status: "available"},
// 					{Row: "A", Column: "5", Status: "available"},
// 					{Row: "A", Column: "6", Status: "available"},
// 					{Row: "A", Column: "7", Status: "available"},
// 					{Row: "A", Column: "8", Status: "available"},
// 					{Row: "A", Column: "9", Status: "available"},
// 					{Row: "A", Column: "10", Status: "available"},
// 				},
// 			},
// 		},
// 	}

// 	body, _ := json.Marshal(cinema)
// 	req, _ := http.NewRequest("POST", "/cinema", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	insertedID := response["inserted_id"].(string)

// 	req, _ = http.NewRequest("DELETE", "/cinema/"+insertedID+"/hall/1", nil)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNoContent, w.Code)

// 	req, _ = http.NewRequest("GET", "/cinema/"+insertedID, nil)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var cinemaResponse models.Cinema
// 	err = json.Unmarshal(w.Body.Bytes(), &cinemaResponse)
// 	assert.NoError(t, err)
// 	assert.Empty(t, cinemaResponse.HallList)
// }
