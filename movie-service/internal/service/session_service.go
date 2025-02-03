package service

import (
	"movie-service/internal/dal"
	"movie-service/internal/models"
	"movie-service/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type SessionService interface {
	AddSession(session models.Session) (*mongo.InsertOneResult, error)
	DeleteAllSession() (*mongo.DeleteResult, error)
}

type sessionService struct {
	sessionRepository dal.SessionRepository
}

func NewSessionService(r dal.SessionRepository) SessionService {
	return &sessionService{
		sessionRepository: r,
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

func (s *sessionService) DeleteAllSession() (*mongo.DeleteResult, error) {
	deleteResult, err := s.sessionRepository.DeleteAllSession()
	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}
