package models

type Hall struct {
	Number      int    `json:"number" bson:"number"`
	RowCount    int    `json:"row_count" bson:"row_count"`
	ColumnCount int    `json:"column_count" bson:"column_count"`
	Seats       []Seat `json:"seats" bson:"seats"`
}
