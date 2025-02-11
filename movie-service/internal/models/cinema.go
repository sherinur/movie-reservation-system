package models

type Cinema struct {
	ID       string  `json:"id" bson:"_id"`
	Name     string  `json:"name" bson:"name"`
	City     string  `json:"city" bson:"city"`
	Address  string  `json:"address" bson:"address"`
	Rating   float64 `json:"rating" bson:"rating"`
	HallList []Hall  `json:"hall_list" bson:"hall_list"`
}
