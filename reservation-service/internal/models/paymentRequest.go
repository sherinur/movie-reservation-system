package models

import "time"

// Reservation Body
type PaymentRequest struct {
	UserId          string    `json:"user_id"`
	ReservationId   string    `json:"reservation_id"`
	PaymentPrice    int       `json:"payment_price"`
	PaymentMethod   string    `json:"payment_method"`
	Status          string    `json:"status"`
	TransactionDate time.Time `json:"transaction_date"`
}
