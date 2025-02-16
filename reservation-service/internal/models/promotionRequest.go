package models

import "time"

// Reservation Body
type PromotionRequest struct {
	Code      string    `json:"code"`
	Discount  int       `json:"discount"`
	ValidFrom time.Time `json:"valid_from"`
	ValidTo   time.Time `json:"valid_to"`
	AppliesTo string    `json:"applies_to"`
}
