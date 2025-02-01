package dal

import (
	"context"

	"reservation-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type ReservationRepository interface {
	GetByUserId(userId string) ([]models.Reservation, error)
	GetById(id string) (*models.Reservation, error)
	Add(process models.Process) (*mongo.InsertOneResult, error)
	Update(id string, reservation models.Reservation) (*mongo.UpdateResult, error)
	Delete(id string) error
}

type reservationRepository struct {
	db *mongo.Database
}

func NewReservationRepository(db *mongo.Database) ReservationRepository {
	return &reservationRepository{db: db}
}

func (r *reservationRepository) GetByUserId(userId string) ([]models.Reservation, error) {
	coll := r.db.Collection("reservations")
	//ObjID, err := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "user_id", Value: userId}}
	var reservations []models.Reservation

	result, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	for result.Next(context.TODO()) {
		var reservation models.Reservation
		err = result.Decode(&reservation)
		if err != nil {
			return nil, err
		}

		reservations = append(reservations, reservation)
	}

	return reservations, nil
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

func (r *reservationRepository) Add(process models.Process) (*mongo.InsertOneResult, error) {
	coll := r.db.Collection("reservations")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "expiring_at", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	}

	_, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return nil, err
	}

	result, err := coll.InsertOne(context.TODO(), process)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *reservationRepository) Update(id string, reservation models.Reservation) (*mongo.UpdateResult, error) {
	coll := r.db.Collection("reservations")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	resDoc := bson.M{
		"_id":          objID,
		"screening_id": reservation.ScreeningID,
		"user_id":      reservation.UserID,
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
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, ErrNotFoundById
	}

	return result, nil
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
		return ErrNotFoundById
	}
	return nil
}
