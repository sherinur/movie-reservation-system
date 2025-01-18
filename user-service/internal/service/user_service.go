package service

import (
	"user-service/internal/dal"
	"user-service/internal/models"
	"user-service/internal/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	Register(req *models.RegisterRequest) (*mongo.InsertOneResult, error)
	Authorize(req *models.LoginRequest) (string, error)
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
	// login validation
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

	jwtToken, err := utils.GenerateJWT(user, []byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (s *userService) Register(req *models.RegisterRequest) (*mongo.InsertOneResult, error) {
	// TODO: Validate the user id

	// check for uniqueness
	existingUser, err := s.userRepository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	} else if existingUser != nil {
		return nil, ErrUserExists
	}

	// password validation
	if !utils.ValidatePassword(req.Password) {
		return nil, ErrInvalidPassword
	}

	// username validation
	if !utils.ValidateUsername(req.Username) {
		return nil, ErrInvalidUsername
	}

	// create a new user
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	newUser, err := s.userRepository.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
