package models

import "time"

type Cinema struct {
	Name     string  `json:"name" bson:"name"`
	Address  string  `json:"address" bson:"address"`
	Rating   float64 `json:"rating" bson:"rating"`
	HallList []Hall  `json:"hall_list" bson:"hall_list"`
}

type Hall struct {
	Number      int         `json:"number" bson:"number"`
	RowCount    int         `json:"row_count" bson:"row_count"`
	ColumnCount int         `json:"column_count" bson:"column_count"`
	Seats       []Seat      `json:"seats" bson:"seats"`
	Screenings  []Screening `json:"screenings" bson:"screenings"`
}

type Seat struct {
	Row    string `json:"row" bson:"row"`
	Column string `json:"column" bson:"column"`
	Status string `json:"status" bson:"status"`
}

type Screening struct {
	MovieID        string    `json:"movie_id" bson:"movie_id"`
	Movie          Movie     `json:"movie" bson:"movie"`
	StartTime      time.Time `json:"start_time" bson:"start_time"`
	EndTime        time.Time `json:"end_time" bson:"end_time"`
	HallNumber     int       `json:"hall_number" bson:"hall_number"`
	AvailableSeats int       `json:"available_seats" bson:"available_seats"`
}
