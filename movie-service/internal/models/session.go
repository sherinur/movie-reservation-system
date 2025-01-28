package models

import "time"

type Session struct {
	MovieID        string    `json:"movie_id" bson:"movie_id"`
	Movie          Movie     `json:"movie" bson:"movie"`
	StartTime      time.Time `json:"start_time" bson:"start_time"`
	EndTime        time.Time `json:"end_time" bson:"end_time"`
	HallNumber     int       `json:"hall_number" bson:"hall_number"`
	AvailableSeats int       `json:"available_seats" bson:"available_seats"`
}
