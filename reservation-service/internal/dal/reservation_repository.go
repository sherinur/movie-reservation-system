package dal

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reservation-service/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ReservationRepository interface {
	Add(reservation models.Reservation) error
	Update(id string) error
	Delete(id string) error
	GetById(id string) (*models.Reservation, error)
}

type reservationRepository struct {
	db *mongo.Database
}

func NewReservationRepository(db *mongo.Database) ReservationRepository {
	return &reservationRepository{db: db}
}

func (r *reservationRepository) Add(reservation models.Reservation) error {
	coll := r.db.Collection("reservations")
	_, err := coll.InsertOne(context.TODO(), reservation)
	return err
}

func (r *reservationRepository) Update(id string) error {
	coll := r.db.Collection("reservations")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"status": "Paid"}}

	_, err = coll.UpdateByID(context.TODO(), objID, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *reservationRepository) Delete(id string) error {
	coll := r.db.Collection("reservations")
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	res, err := coll.DeleteOne(context.TODO(), bson.M{"_id": ObjID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no reservation found with the given ID")
	}
	return nil
}

func (r *reservationRepository) GetById(id string) (*models.Reservation, error) {
	coll := r.db.Collection("reservations")
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var reservation models.Reservation
	err = coll.FindOne(context.TODO(), bson.M{"_id": ObjID}).Decode(&reservation)
	if err != nil {
		return nil, err
	}

	return &reservation, nil
}
