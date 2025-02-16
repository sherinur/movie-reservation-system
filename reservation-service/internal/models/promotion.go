package models

import "time"

// Reservation Body
type Promotion struct {
	Code      string    `bson:"code"`
	Discount  int       `bson:"discount"`
	ValidFrom time.Time `bson:"valid_from"`
	ValidTo   time.Time `bson:"valid_to"`
	AppliesTo string    `bson:"applies_to"`
}
