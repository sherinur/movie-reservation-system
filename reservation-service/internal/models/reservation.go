package models

import "time"

type Reservation struct {
	ScreeningID string    `bson:"screening_id"`
	Email       string    `bson:"email"`
	PhoneNumber string    `bson:"phone_number"`
	Status      string    `bson:"status"`
	Tickets     []Ticket  `bson:"tickets"`
	TotalPrice  float64   `bson:"total_price"`
	QRCode      string    `bson:"qr_code"`
	BoughtTime  time.Time `bson:"bought_time"`
	ExpiringAt  time.Time `bson:"expiring_at"`
}
