package dal

import (
	"context"
	"errors"

	"reservation-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type ReservationRepository interface {
	Add(process models.Process) (*mongo.InsertOneResult, error)
	Update(id string, reservation models.Reservation) (*mongo.UpdateResult, error)
	Delete(id string) error
	GetById(id string) (*models.Reservation, error)
}

type reservationRepository struct {
	db *mongo.Database
}

func NewReservationRepository(db *mongo.Database) ReservationRepository {
	return &reservationRepository{db: db}
}

func (r *reservationRepository) Add(process models.Process) (*mongo.InsertOneResult, error) {
	coll := r.db.Collection("reservations")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "expiring_at", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	}

	_, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return nil, errors.New("failed to add expiring after time index to collection")
	}

	result, err := coll.InsertOne(context.TODO(), process)

	if err != nil {
		return nil, errors.New("error adding new process to repository")
	}

	return result, nil
}

func (r *reservationRepository) Update(id string, reservation models.Reservation) (*mongo.UpdateResult, error) {
	coll := r.db.Collection("reservations")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("error getting objID")
	}

	resDoc := bson.M{
		"_id":          objID,
		"screening_id": reservation.ScreeningID,
		"email":        reservation.Email,
		"phone_number": reservation.PhoneNumber,
		"status":       reservation.Status,
		"tickets":      reservation.Tickets,
		"total_price":  reservation.TotalPrice,
		"qr_code":      reservation.QRCode,
		"bought_time":  reservation.BoughtTime,
	}

	result, err := coll.ReplaceOne(context.TODO(), bson.M{"_id": objID}, resDoc)
	if err != nil {
		return nil, errors.New("could not update in repository by id: " + err.Error())
	}
	if result.ModifiedCount == 0 {
		return nil, errors.New("no process found with the given ID")
	}

	return result, nil
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
		return nil, err
	}

	return &reservation, nil
}
