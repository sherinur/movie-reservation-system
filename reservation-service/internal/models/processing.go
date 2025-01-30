package models

import "time"

type Process struct {
	ScreeningID string    `bson:"screening_id"`
	UserID      string    `bson:"user_id"`
	Status      string    `bson:"status"`
	Tickets     []Ticket  `bson:"tickets"`
	TotalPrice  float64   `bson:"total_price"`
	ExpiringAt  time.Time `bson:"expiring_at"`
}
