package models

// Reservation Request Body to pay the processing
type ReservationRequest struct {
	UserID      string  `json:"user_id"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	TotalPrice  float64 `json:"total_price"`
}
