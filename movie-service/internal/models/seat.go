package models

type Seat struct {
	Row    string `json:"row" bson:"row"`
	Column string `json:"column" bson:"column"`
	Status string `json:"status" bson:"status"`
}
