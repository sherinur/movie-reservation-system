package models

type Hall struct {
	// ID          string `json:"id" bson:"_id"`
	Number      int    `json:"number" bson:"number"`
	RowCount    int    `json:"row_count" bson:"row_count"`
	ColumnCount int    `json:"column_count" bson:"column_count"`
	Seats       []Seat `json:"seats" bson:"seats"`
}

type Hall_list struct {
	Hall_list []Hall `json:"hall_list" bson:"hall_list"`
}
