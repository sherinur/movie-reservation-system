package models

// Processing Request Body for creating new Processing
type ProcessingRequest struct {
	ScreeningID string   `json:"screening_id"`
	UserID      string   `json:"user_id"`
	MovieTitle  string   `json:"movie_title"`
	Tickets     []Ticket `json:"tickets"`
}
