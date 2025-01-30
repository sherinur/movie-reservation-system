package models

type Booking struct {
	ScreeningID string   `json:"screening_id"`
	UserID      string   `json:"user_id"`
	Tickets     []Ticket `json:"tickets"`
}
