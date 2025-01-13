package service

import (
	"user-service/internal/dal"
	"user-service/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	Register() (*mongo.InsertOneResult, error)
}

type userService struct {
	userRepository dal.UserRepository
}

func NewUserService(r dal.UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}

func (s *userService) Register() (*mongo.InsertOneResult, error) {
	user := &models.User{
		Username: "sherinur",
		Email:    "sherinurislam@gmail.com",
		Password: "123",
	}

	result, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
