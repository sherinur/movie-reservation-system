package models

type Movie struct {
	Title       string `bson:"title"`
	Genre       string `bson:"genre"`
	Description string `bson:"description"`
	PosterImage string `bson:"posterimage"`
	Duration    int    `bson:"duration"`
	Language    string `bson:"language"`
	ReleaseDate string `bson:"releasedate"`
	Rating      string `bson:"rating"`
	PGrating    string `bson:"pgrating"`
	Production  string `bson:"production"`
	Producer    string `bson:"producer"`
	Status      string `bson:"status"`
}
