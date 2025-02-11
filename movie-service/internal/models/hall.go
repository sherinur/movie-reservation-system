package models

type Hall struct {
	Number         int    `json:"number" bson:"number"`
	AvailableSeats int    `json:"available_seats" bson:"available_seats"`
	Seats          []Seat `json:"seats" bson:"seats"`
}

type Hall_list struct {
	Hall_list []Hall `json:"hall_list" bson:"hall_list"`
}
