package models

type Reservation struct {
	MovieTitle string `json:"movie_title"`
	Email      string `json:"email"`
	Status     string `json:"status"`
	DateTime   string `json:"date_time"`
}
