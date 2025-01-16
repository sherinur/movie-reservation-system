package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	ID           primitive.ObjectID `bson:"_id"`
	MovieName    string `bson:"movie_name"`
	Email        string `bson:"email"`
	Seat         string `bson:"seat"`
	PurchaseTime string `bson:"purchase_time"`
}

