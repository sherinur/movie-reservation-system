package models

type Booking struct {
	ScreeningID string   `json:"screening_id"`
	Tickets     []Ticket `json:"tickets"`
}
