package dal

import (
	"context"
	"movie-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SessionRepository interface {
	AddSession(session models.Session) (*mongo.InsertOneResult, error)
	GetAllSession() ([]models.Session, error)
	UpdateSessionByID(sessionID string, session models.Session) (*mongo.UpdateResult, error)
	DeleteSessionByID(sessionID string) (*mongo.DeleteResult, error)
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
	col := r.db.Collection("cinema")

	filter := bson.M{"_id": session.CinemaID, "hall_list.number": session.HallNumber}
	projection := bson.M{"hall_list.$": 1}

	var result struct {
		HallList []models.Hall `bson:"hall_list"`
	}

	err := col.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}

	if len(result.HallList) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	session.Seats = result.HallList[0].Seats

	sessionCol := r.db.Collection("session")
	insertRes, err := sessionCol.InsertOne(context.TODO(), session)
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

func (r *sessionRepository) UpdateSessionByID(sessionID string, session models.Session) (*mongo.UpdateResult, error) {
	col := r.db.Collection("session")

	objectID, err := primitive.ObjectIDFromHex(sessionID)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": session,
	}

	updateResult, err := col.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (r *sessionRepository) DeleteSessionByID(sessionID string) (*mongo.DeleteResult, error) {
	col := r.db.Collection("session")

	objectID, err := primitive.ObjectIDFromHex(sessionID)
	if err != nil {
		return nil, err
	}

	deleteResult, err := col.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}})
	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}

func (r *sessionRepository) DeleteAllSession() (*mongo.DeleteResult, error) {
	col := r.db.Collection("session")

	deletResult, err := col.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	return deletResult, nil
}
