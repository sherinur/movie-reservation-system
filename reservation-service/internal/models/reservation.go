package models

type Reservation struct {
	MovieTitle string   `json:"movie_title"`
	Email      string   `json:"email"`
	Status     string   `json:"status"`
	BoughtTime string   `json:"bought_time"`
	Tickets    []Ticket `json:"tickets"`
}

type Ticket struct {
	Seat  string  `json:"seat"`
	Price float64 `json:"price"`
	Type  string  `json:"type"`
}
