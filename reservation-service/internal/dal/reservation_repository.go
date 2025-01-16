package dal

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"reservation-service/reservation-service/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ReservationRepository interface {
	AddTicket(ticket *models.Ticket) (*mongo.InsertOneResult, error)
	RemoveTicket(id string) (*mongo.DeleteResult, error)
}

type reservationRepository struct {
	db *mongo.Database
}

func NewReservationRepository(db *mongo.Database) ReservationRepository {
	return &reservationRepository{db: db}
}

func (r *reservationRepository) AddTicket(ticket *models.Ticket) (*mongo.InsertOneResult, error) {
	coll := r.db.Collection("reservations")
	result, err := coll.InsertOne(context.Background(), ticket)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *reservationRepository) RemoveTicket(id string) (*mongo.DeleteResult, error) {
	coll := r.db.Collection("reservations")
	result, err := coll.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return nil, err
	}

	if result.DeletedCount == 0 {
		return nil, errors.New("reservation not found")
	}

	return result, nil
}
