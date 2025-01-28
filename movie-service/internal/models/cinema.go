package models

type Cinema struct {
	Name     string  `json:"name" bson:"name"`
	Address  string  `json:"address" bson:"address"`
	Rating   float64 `json:"rating" bson:"rating"`
	HallList []Hall  `json:"hall_list" bson:"hall_list"`
}
