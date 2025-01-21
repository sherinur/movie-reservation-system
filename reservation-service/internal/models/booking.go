package models

type Booking struct {
	MovieTitle string   `json:"movie_title"`
	Email      string   `json:"email"`
	Tickets    []Ticket `json:"tickets"`
}
