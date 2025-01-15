package service

import (
	"user-service/internal/dal"
	"user-service/internal/models"
	"user-service/internal/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	Register(registerRequest *models.RegisterRequest) (*mongo.InsertOneResult, error)
	Authorize(loginRequest *models.LoginRequest) (string, error)
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	userRepository dal.UserRepository
	secretKey      string
}

func NewUserService(r dal.UserRepository, secretKey string) UserService {
	return &userService{
		userRepository: r,
		secretKey:      secretKey,
	}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) Authorize(req *models.LoginRequest) (string, error) {
	// TODO: Generate and return JWT token using secret key

	user, err := s.userRepository.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", ErrNoUser
	}

	if req.Password != user.Password {
		return "", ErrWrongPassword
	}

	return "", nil
}

func (s *userService) Register(registerRequest *models.RegisterRequest) (*mongo.InsertOneResult, error) {
	// TODO: Validate the user id

	// password validation
	if !utils.ValidatePassword(registerRequest.Password) {
		return nil, ErrInvalidPassword
	}

	// username validation
	if !utils.ValidateUsername(registerRequest.Username) {
		return nil, ErrInvalidUsername
	}

	// check for uniqueness
	existingUser, err := s.userRepository.GetUserByEmail(registerRequest.Email)
	if err != nil {
		return nil, err
	} else if existingUser != nil {
		return nil, ErrUserExists
	}

	// create a new user
	user := models.User{
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}

	newUser, err := s.userRepository.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
