package models

import "time"

// Reservation Body
type Payment struct {
	UserId          string    `bson:"user_id"`
	ReservationId   string    `bson:"reservation_id"`
	PaymentPrice    int       `bson:"payment_price"`
	PaymentMethod   string    `bson:"payment_method"`
	Status          string    `bson:"status"`
	TransactionDate time.Time `bson:"transaction_date"`
}
