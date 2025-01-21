package models

import "time"

// TODO!: think about struct cinema and add special data
type Cinema struct {
	Name     string  `bson:"name"`
	Address  string  `bson:"address"`
	Rating   float64 `bson:"rating"`
	HallList []Hall  `bson:"hall_list"`
}

type Hall struct {
	Number      int         `bson:"number"`
	RowCount    int         `bson:"row_count"`
	ColumnCount int         `bson:"column_count"`
	Seats       []Seat      `bson:"seats"`
	Screenings  []Screening `bson:"screenings"`
}

type Seat struct {
	Row    string `bson:"row"`
	Column string `bson:"column"`
	Status string `bson:"status"`
}

type Screening struct {
	MovieID        string    `bson:"movie_id"`
	Movie          Movie     `bson:"movie"`
	StartTime      time.Time `bson:"start_time"`
	EndTime        time.Time `bson:"end_time"`
	HallNumber     int       `bson:"hall_number"`
	AvailableSeats int       `bson:"available_seats"`
}
