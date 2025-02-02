package service

import (
	"movie-service/internal/models"
	"movie-service/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type SessionService interface {
	AddSession(session models.Session) (*mongo.InsertOneResult, error)
}

type sessionService struct {
	db *mongo.Database
}

func NewSessionService(db *mongo.Database) SessionService {
	return &sessionService{
		db: db,
	}
}

func (s *sessionService) AddSession(session models.Session) (*mongo.InsertOneResult, error) {
	err := utils.ValidateSesesion(session)
	if err != nil {
		return nil, err
	}

	insertResult, err := s.AddSession(session)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}
