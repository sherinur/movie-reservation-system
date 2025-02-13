package models

import "time"

type Session struct {
	ID             string    `json:"id" bson:"_id"`
	MovieID        string    `json:"movie_id" bson:"movie_id"`
	CinemaID       string    `json:"cinema_id" bson:"cinema_id"`
	CinemaAddres   string    `json:"address" bson:"address"`
	HallNumber     int       `json:"hall_number" bson:"hall_number"`
	Date           time.Time `json:"date" bson:"date"`
	StartTime      string    `json:"start_time" bson:"start_time"`
	EndTime        string    `json:"end_time" bson:"end_time"`
	AvailableSeats int       `json:"available_seats" bson:"available_seats"`
	Seats          []Seat    `json:"seats" bson:"seats"`
}
