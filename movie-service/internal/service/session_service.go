package service

import (
	"movie-service/internal/dal"
	"movie-service/internal/models"
	"movie-service/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type SessionService interface {
	AddSession(session models.Session) (*mongo.InsertOneResult, error)
	GetAllSession() ([]models.Session, error)
	GetSessionByID(sessionID string) (*models.Session, error)
	UpdateSessionByID(sessionID string, session models.Session) (*mongo.UpdateResult, error)
	DeleteSessionByID(sessionID string) (*mongo.DeleteResult, error)
	DeleteAllSession() (*mongo.DeleteResult, error)

	GetSeat(sessionID string) ([]models.Seat, error)
	GetSessionsByMovieID(movieID string) ([]models.Session, error)
	PostSeatClose(sessionID string, seat models.Seat) (*mongo.UpdateResult, error)
	PostSeatOpen(sessionID string, seat models.Seat) (*mongo.UpdateResult, error)
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

	insertResult, err := s.sessionRepository.AddSession(session)
	if err != nil {
		return nil, err
	}
	return insertResult, nil
}

func (s *sessionService) GetAllSession() ([]models.Session, error) {
	session, err := s.sessionRepository.GetAllSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sessionService) GetSessionByID(sessionID string) (*models.Session, error) {
	if sessionID == "" {
		return nil, utils.ErrInvalidId
	}

	session, err := s.sessionRepository.GetSessionByID(sessionID)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sessionService) GetSeat(sessionID string) ([]models.Seat, error) {
	if sessionID == "" {
		return nil, utils.ErrInvalidId
	}

	Seats, err := s.sessionRepository.GetSeats(sessionID)
	if err != nil {
		return nil, err
	}
	return Seats, nil
}

func (s *sessionService) UpdateSessionByID(sessionID string, session models.Session) (*mongo.UpdateResult, error) {
	if sessionID == "" {
		return nil, utils.ErrInvalidId
	}

	err := utils.ValidateSesesion(session)
	if err != nil {
		return nil, err
	}

	updateResult, err := s.sessionRepository.UpdateSessionByID(sessionID, session)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (s *sessionService) DeleteSessionByID(sessionID string) (*mongo.DeleteResult, error) {
	if sessionID == "" {
		return nil, utils.ErrInvalidId
	}

	deleteResult, err := s.sessionRepository.DeleteSessionByID(sessionID)
	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}

func (s *sessionService) DeleteAllSession() (*mongo.DeleteResult, error) {
	deleteResult, err := s.sessionRepository.DeleteAllSession()
	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}

func (s *sessionService) GetSessionsByMovieID(movieID string) ([]models.Session, error) {
	if movieID == "" {
		return nil, utils.ErrInvalidId
	}

	sessions, err := s.sessionRepository.GetSessionsByMovieID(movieID)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *sessionService) PostSeatClose(sessionID string, seat models.Seat) (*mongo.UpdateResult, error) {
	if sessionID == "" {
		return nil, utils.ErrInvalidId
	}

	updateResult, err := s.sessionRepository.PostSeatClose(sessionID, seat)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (s *sessionService) PostSeatOpen(sessionID string, seat models.Seat) (*mongo.UpdateResult, error) {
	if sessionID == "" {
		return nil, utils.ErrInvalidId
	}

	updateResult, err := s.sessionRepository.PostSeatOpen(sessionID, seat)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}
