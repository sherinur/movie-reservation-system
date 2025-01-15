package dal

import "go.mongodb.org/mongo-driver/mongo"

type ReservationRepository interface {
}

type reservationRepository struct {
	db *mongo.Database
}

func NewReservationRepository(db *mongo.Database) ReservationRepository {
	return &reservationRepository{db: db}
}
