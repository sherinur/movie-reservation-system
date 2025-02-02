package dal

import (
	"context"
	"movie-service/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepository interface{}

type sessionRepository struct {
	db *mongo.Database
}

func NewSessionRepository(db *mongo.Database) SessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (r *sessionRepository) AddSession(session models.Session) (*mongo.InsertOneResult, error) {
	col := r.db.Collection("session")

	insertRes, err := col.InsertOne(context.TODO(), session)
	if err != nil {
		return nil, err
	}

	return insertRes, nil
}
