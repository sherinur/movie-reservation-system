package dal

import (
	"context"
	"errors"

	"reservation-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
		return errors.New("error getting objID")
	}

	update := bson.M{"$set": bson.M{"status": "Paid"}}

	_, err = coll.UpdateByID(context.TODO(), objID, update)
	if err != nil {
		return errors.New("could not update in repository by id")
	}

	return nil
}

func (r *reservationRepository) Delete(id string) error {
	coll := r.db.Collection("reservations")
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("error getting objID")
	}

	res, err := coll.DeleteOne(context.TODO(), bson.M{"_id": ObjID})
	if err != nil {
		return errors.New("could not delete from repository by id")
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
		return nil, errors.New("error getting objID")
	}

	var reservation models.Reservation
	err = coll.FindOne(context.TODO(), bson.M{"_id": ObjID}).Decode(&reservation)
	if err != nil {
		return nil, errors.New("no reservation found with the given ID")
	}

	return &reservation, nil
}
