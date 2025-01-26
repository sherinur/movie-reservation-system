package models

type Ticket struct {
	SeatRow    string  `json:"seat_row"`
	SeatColumn string  `json:"seat_column"`
	Price      float64 `json:"price"`
	Type       string  `json:"type"`
}
