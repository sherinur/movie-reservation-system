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
	GetByUserId(ctx context.Context, userId string) ([]models.Reservation, error)
	GetById(ctx context.Context, id string) (*models.Reservation, error)
	AddReservation(ctx context.Context, process models.Reservation) (*mongo.InsertOneResult, error)
	UpdateReservation(ctx context.Context, id string, reservation models.Reservation) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) error
}

type reservationRepository struct {
	db *mongo.Database
}

func NewReservationRepository(db *mongo.Database) ReservationRepository {
	return &reservationRepository{db: db}
}

func (r *reservationRepository) GetByUserId(ctx context.Context, userId string) ([]models.Reservation, error) {
	coll := r.db.Collection("reservations")
	//ObjID, err := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "user_id", Value: userId}}
	var reservations []models.Reservation

	result, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for result.Next(ctx) {
		var reservation models.Reservation
		err = result.Decode(&reservation)
		if err != nil {
			return nil, err
		}

		reservations = append(reservations, reservation)
	}

	return reservations, nil
}

func (r *reservationRepository) GetById(ctx context.Context, id string) (*models.Reservation, error) {
	coll := r.db.Collection("reservations")
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var reservation models.Reservation
	err = coll.FindOne(ctx, bson.M{"_id": ObjID}).Decode(&reservation)
	if err != nil {
		return nil, err
	}

	return &reservation, nil
}

func (r *reservationRepository) AddReservation(ctx context.Context, process models.Reservation) (*mongo.InsertOneResult, error) {
	coll := r.db.Collection("reservations")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "expiring_at", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	}

	_, err := coll.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return nil, err
	}

	result, err := coll.InsertOne(ctx, process)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *reservationRepository) UpdateReservation(ctx context.Context, id string, reservation models.Reservation) (*mongo.UpdateResult, error) {
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

	result, err := coll.ReplaceOne(ctx, bson.M{"_id": objID}, resDoc)
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, ErrNotFoundById
	}

	return result, nil
}

func (r *reservationRepository) Delete(ctx context.Context, id string) error {
	coll := r.db.Collection("reservations")
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	res, err := coll.DeleteOne(ctx, bson.M{"_id": ObjID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return ErrNotFoundById
	}
	return nil
}
