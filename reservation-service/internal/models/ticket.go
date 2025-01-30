package models

type Ticket struct {
	SeatRow    string  `json:"seat_row"`
	SeatColumn string  `json:"seat_column"`
	Price      float64 `json:"price"`
	SeatType   string  `json:"seat_type"`
	UserType   string  `json:"user_type"`
}
