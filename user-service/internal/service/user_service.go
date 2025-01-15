package service

import (
	"user-service/internal/dal"
	"user-service/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	Register(user *models.User) (*mongo.InsertOneResult, error)
	Authorize(loginRequest *models.LoginRequest) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	userRepository dal.UserRepository
}

func NewUserService(r dal.UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) Authorize(req *models.LoginRequest) (*models.User, error) {
	user, err := s.userRepository.FindUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrNoUser
	}

	if req.Password != user.Password {
		return nil, ErrWrongPassword
	}

	return user, nil
}

func (s *userService) Register(user *models.User) (*mongo.InsertOneResult, error) {
	// TODO: Validate user data

	result, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
