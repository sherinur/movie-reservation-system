package handler

// func setupMovieTestRouter() (*gin.Engine, error) {
// 	gin.SetMode(gin.TestMode)
// 	r := gin.New()

// 	db, err := db.ConnectMongo("mongodb://localhost:27017", "movietest")
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := db.Collection("movies").Drop(context.TODO()); err != nil {
// 		return nil, err
// 	}

// 	repo := dal.NewMovieRepository(db)
// 	service := service.NewMovieService(repo)
// 	handler := NewMovieHandler(service, logging.NewLogger("test"))

// 	r.POST("/movie", handler.HandleAddMovie)
// 	r.GET("/movielist", handler.HandleGetAllMovie)
// 	r.GET("/movie/:id", handler.HadleGetMovieById)
// 	r.PUT("/movie/:id", handler.HandleUpdateMovieById)
// 	r.DELETE("/movie/:id", handler.HandleDeleteMovieByID)
// 	r.DELETE("/movie", handler.HandleDeleteAllMovie)

// 	return r, nil
// }

// func TestHandleAddMovie(t *testing.T) {
// 	r, err := setupMovieTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	movie := models.Movie{
// 		Title:       "The Shawshank Redemption",
// 		Genre:       "Drama",
// 		Description: "Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.",
// 		PosterImage: "https://m.media-amazon.com/images/M/MV5BMDAyY2FhYjctNDc5OS00MDNlLThiMGUtY2UxYWVkNGY2ZjljXkEyXkFqcGc@._V1_QL75_UX140_CR0,1,140,207_.jpg",
// 		Duration:    142,
// 		Language:    "English",
// 		ReleaseDate: "1994",
// 		Rating:      "9.3",
// 		PGrating:    "R",
// 		Production:  "Castle Rock Entertainment",
// 		Producer:    "Niki Marvin",
// 		Status:      "Released",
// 	}

// 	body, _ := json.Marshal(movie)
// 	req, _ := http.NewRequest("POST", "/movie", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, response["inserted_id"])
// }

// func TestHandleGetAllMovie(t *testing.T) {
// 	r, err := setupMovieTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req, _ := http.NewRequest("GET", "/movielist", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response []models.Movie
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, response)
// }

// func TestHandleGetMovieById(t *testing.T) {
// 	r, err := setupMovieTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	movie := models.Movie{
// 		Title:       "The Godfather",
// 		Genre:       "Crime, Drama",
// 		Description: "The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.",
// 		PosterImage: "https://m.media-amazon.com/images/M/MV5BNGEwYjgwOGQtYjg5ZS00Njc1LTk2ZGEtM2QwZWQ2NjdhZTE5XkEyXkFqcGc@._V1_QL75_UY207_CR3,0,140,207_.jpg",
// 		Duration:    175,
// 		Language:    "English, Italian",
// 		ReleaseDate: "1972",
// 		Rating:      "9.2",
// 		PGrating:    "R",
// 		Production:  "Paramount Pictures",
// 		Producer:    "Albert S. Ruddy",
// 		Status:      "Released",
// 	}

// 	body, _ := json.Marshal(movie)
// 	req, _ := http.NewRequest("POST", "/movie", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	insertedID := response["inserted_id"].(string)

// 	req, _ = http.NewRequest("GET", "/movie/"+insertedID, nil)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var movieResponse models.Movie
// 	err = json.Unmarshal(w.Body.Bytes(), &movieResponse)
// 	assert.NoError(t, err)
// 	assert.Equal(t, movie.Title, movieResponse.Title)
// }

// func TestHandleUpdateMovieById(t *testing.T) {
// 	r, err := setupMovieTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	movie := models.Movie{
// 		Title:       "The Chaos Class",
// 		Genre:       "Comedy",
// 		Description: "Lazy, uneducated students share a very close bond. They live together in the dormitory, where they plan their latest pranks. When a new headmaster arrives, the students naturally try to overthrow him. A comic war of nitwits follows.",
// 		PosterImage: "https://m.media-amazon.com/images/M/MV5BZTdhN2ViYjctMGZlZi00ZGRmLWIxMWQtNzJiMDY1ZDcxNjJmXkEyXkFqcGc@._V1_QL75_UX140_CR0,1,140,207_.jpg",
// 		Duration:    85,
// 		Language:    "English, Italian",
// 		ReleaseDate: "1975",
// 		Rating:      "9.2",
// 		PGrating:    "R",
// 		Production:  "Unknown",
// 		Producer:    "Ertem Egilmez",
// 		Status:      "Released",
// 	}

// 	body, _ := json.Marshal(movie)
// 	req, _ := http.NewRequest("POST", "/movie", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	insertedID := response["inserted_id"].(string)

// 	updatedMovie := models.Movie{
// 		Title:       "The Dark Knight",
// 		Genre:       "Action, Crime, Drama",
// 		Description: "When a menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman, James Gordon and Harvey Dent must work together to put an end to the madness.",
// 		PosterImage: "https://m.media-amazon.com/images/M/MV5BMTMxNTMwODM0NF5BMl5BanBnXkFtZTcwODAyMTk2Mw@@._V1_QL75_UX140_CR0,0,140,207_.jpg",
// 		Duration:    152,
// 		Language:    "English",
// 		ReleaseDate: "2008",
// 		Rating:      "9.0",
// 		PGrating:    "PG-13",
// 		Production:  "Warner Bros.",
// 		Producer:    "Christopher Nolan, Emma Thomas",
// 		Status:      "Released",
// 	}

// 	updatedBody, _ := json.Marshal(updatedMovie)
// 	req, _ = http.NewRequest("PUT", "/movie/"+insertedID, bytes.NewReader(updatedBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	req, _ = http.NewRequest("GET", "/movie/"+insertedID, nil)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var movieResponse models.Movie
// 	err = json.Unmarshal(w.Body.Bytes(), &movieResponse)
// 	assert.NoError(t, err)
// 	assert.Equal(t, updatedMovie.Title, movieResponse.Title)
// }

// func TestHandleDeleteMovieByID(t *testing.T) {
// 	r, err := setupMovieTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	movie := models.Movie{
// 		Title:       "Schindler's List",
// 		Genre:       "Biography, Drama, History",
// 		Description: "In German-occupied Poland during World War II, industrialist Oskar Schindler gradually becomes concerned for his Jewish workforce after witnessing their persecution by the Nazis.",
// 		PosterImage: "https://m.media-amazon.com/images/M/MV5BNjM1ZDQxYWUtMzQyZS00MTE1LWJmZGYtNGUyNTdlYjM3ZmVmXkEyXkFqcGc@._V1_QL75_UX140_CR0,1,140,207_.jpg",
// 		Duration:    195,
// 		Language:    "English, German, Hebrew, Polish",
// 		ReleaseDate: "1993",
// 		Rating:      "9.0",
// 		PGrating:    "R",
// 		Production:  "Universal Pictures",
// 		Producer:    "Steven Spielberg, Branko Lustig, Gerald R. Molen",
// 		Status:      "Released",
// 	}

// 	body, _ := json.Marshal(movie)
// 	req, _ := http.NewRequest("POST", "/movie", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var response map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	insertedID := response["inserted_id"].(string)

// 	req, _ = http.NewRequest("DELETE", "/movie/"+insertedID, nil)
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNoContent, w.Code)
// }

// func TestHandleDeleteAllMovie(t *testing.T) {
// 	r, err := setupMovieTestRouter()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req, _ := http.NewRequest("DELETE", "/movie", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNoContent, w.Code)
// }
