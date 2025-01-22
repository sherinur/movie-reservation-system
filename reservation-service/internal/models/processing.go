package models

type Process struct {
	ScreeningID string   `json:"screening_id"`
	Status      string   `json:"status"`
	Tickets     []Ticket `json:"tickets"`
	TotalPrice  float64  `json:"total_price"`
	CreatedTime string   `json:"created_time"`
}
