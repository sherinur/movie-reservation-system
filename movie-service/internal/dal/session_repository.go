package dal

import (
	"context"
	"movie-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepository interface {
	AddSession(session models.Session) (*mongo.InsertOneResult, error)
	DeleteAllSession() (*mongo.DeleteResult, error)
}

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

func (r *sessionRepository) GetAllSession() ([]models.Session, error) {
	col := r.db.Collection("session")

	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var session []models.Session
	err = cursor.All(context.TODO(), &session)
	if err != nil {
		return nil, err
	}

	return session, nil

}

func (r *sessionRepository) DeleteAllSession() (*mongo.DeleteResult, error) {
	col := r.db.Collection("session")

	deletResult, err := col.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	return deletResult, nil
}
