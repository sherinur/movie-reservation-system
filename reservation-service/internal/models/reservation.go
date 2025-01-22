package models

type Reservation struct {
	ScreeningID  string   `json:"screening_id"`
	Email        string   `json:"email"`
	PhoneNumber  string   `json:"phone_number"`
	Status       string   `json:"status"`
	Tickets      []Ticket `json:"tickets"`
	TotalPrice   float64  `json:"total_price"`
	QRCode       string   `json:"qr_code"`
	BoughtTime   string   `json:"bought_time"`
	ExpiringTime string   `json:"expiring_time"`
}
