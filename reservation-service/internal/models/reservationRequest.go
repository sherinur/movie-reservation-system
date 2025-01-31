package models

// Reservation Request Body to pay the processing
type ReservationRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
