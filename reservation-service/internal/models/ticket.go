package models

type Ticket struct {
	Seat     string  `json:"seat"`
	Price    float64 `json:"price"`
	SeatType string  `json:"seat_type"`
	UserType string  `json:"user_type"`
}
