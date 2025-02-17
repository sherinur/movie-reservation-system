package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Reservation Body
type Reservation struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ScreeningID string             `bson:"screening_id"`
	UserID      string             `bson:"user_id"`
	Email       string             `bson:"email"`
	PhoneNumber string             `bson:"phone_number"`
	Status      string             `bson:"status"`
	Tickets     []Ticket           `bson:"tickets"`
	TotalPrice  float64            `bson:"total_price"`
	QRCode      string             `bson:"qr_code"`
	BoughtTime  time.Time          `bson:"bought_time"`
	ExpiringAt  time.Time          `bson:"expiring_at"`
}
