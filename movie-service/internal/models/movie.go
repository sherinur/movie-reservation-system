package models

type Movie struct {
	ID          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Genre       string `json:"genre" bson:"genre"`
	Description string `json:"description" bson:"description"`
	PosterImage string `json:"posterimage" bson:"posterimage"`
	Duration    int    `json:"duration" bson:"duration"`
	Language    string `json:"language" bson:"language"`
	ReleaseDate string `json:"releasedate" bson:"releasedate"`
	Rating      string `json:"rating" bson:"rating"`
	PGrating    string `json:"pgrating" bson:"pgrating"`
	Production  string `json:"production" bson:"production"`
	Producer    string `json:"producer" bson:"producer"`
	Status      string `json:"status" bson:"status"`
}
