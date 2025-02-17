package dal

import (
	"context"
	"strings"
	"time"

	"movie-service/internal/models"
	"movie-service/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepository interface {
	AddSession(session models.Session) (*mongo.InsertOneResult, error)
	GetAllSession() ([]models.Session, error)
	GetSessionByID(sessionID string) (*models.Session, error)
	UpdateSessionByID(sessionID string, session models.Session) (*mongo.UpdateResult, error)
	DeleteSessionByID(sessionID string) (*mongo.DeleteResult, error)
	DeleteAllSession() (*mongo.DeleteResult, error)

	GetSeats(sessionID string) ([]models.Seat, error)
	GetSessionsByMovieID(movieID string) ([]models.Session, error)
	PostSeatClose(sessionID string, seat models.Seat) (*mongo.UpdateResult, error)
	PostSeatOpen(sessionID string, seat models.Seat) (*mongo.UpdateResult, error)
}

type sessionRepository struct {
	db               *mongo.Database
	cinemaRepository CinemaRepository
}

func NewSessionRepository(db *mongo.Database, r CinemaRepository) SessionRepository {
	return &sessionRepository{
		db:               db,
		cinemaRepository: r,
	}
}

func (r *sessionRepository) AddSession(session models.Session) (*mongo.InsertOneResult, error) {
	hall, err := r.cinemaRepository.GetHall(session.CinemaID, session.HallNumber)
	if err != nil {
		return nil, err
	}

	if len(strings.TrimSpace(session.ID)) == 0 {
		session.ID = utils.GenerateID()
	}

	session.Seats = hall.Seats
	session.AvailableSeats = hall.AvailableSeats

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

func (r *sessionRepository) GetSessionByID(sessionID string) (*models.Session, error) {
	col := r.db.Collection("session")

	var session models.Session
	err := col.FindOne(context.TODO(), bson.M{"_id": sessionID}).Decode(&session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *sessionRepository) GetSeats(sessionID string) ([]models.Seat, error) {
	col := r.db.Collection("session")

	var session models.Session
	err := col.FindOne(context.TODO(), bson.M{"_id": sessionID}).Decode(&session)
	if err != nil {
		return nil, err
	}

	return session.Seats, nil
}

func (r *sessionRepository) UpdateSessionByID(sessionID string, session models.Session) (*mongo.UpdateResult, error) {
	col := r.db.Collection("session")

	update := bson.D{
		{Key: "movie_id", Value: session.MovieID},
		{Key: "cinema_id", Value: session.CinemaID},
		{Key: "hall_number", Value: session.HallNumber},
		{Key: "start_time", Value: session.StartTime},
		{Key: "end_time", Value: session.EndTime},
		{Key: "seats", Value: session.Seats},
		{Key: "available_seats", Value: session.AvailableSeats},
	}

	updateResult, err := col.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: sessionID}}, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (r *sessionRepository) DeleteSessionByID(sessionID string) (*mongo.DeleteResult, error) {
	col := r.db.Collection("session")

	deleteResult, err := col.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: sessionID}})
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

func (r *sessionRepository) GetSessionsByMovieID(movieID string) ([]models.Session, error) {
	col := r.db.Collection("session")

	currentTime := time.Now()

	filter := bson.M{
		"movie_id":   movieID,
		"start_time": bson.M{"$gte": currentTime},
	}

	cursor, err := col.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var sessions []models.Session
	err = cursor.All(context.TODO(), &sessions)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (r *sessionRepository) PostSeatClose(sessionID string, seat models.Seat) (*mongo.UpdateResult, error) {
	col := r.db.Collection("session")
	seat.Status = "Busy"

	filter := bson.M{"_id": sessionID, "seats.row": seat.Row, "seats.column": seat.Column}
	update := bson.M{"$set": bson.M{"seats.$.status": seat.Status}}

	updateResult, err := col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (r *sessionRepository) PostSeatOpen(sessionID string, seat models.Seat) (*mongo.UpdateResult, error) {
	col := r.db.Collection("session")
	seat.Status = "Available"

	filter := bson.M{"_id": sessionID, "seats.row": seat.Row, "seats.column": seat.Column}
	update := bson.M{"$set": bson.M{"seats.$.status": seat.Status}}

	updateResult, err := col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}
