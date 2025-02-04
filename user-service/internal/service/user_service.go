package service

import (
	"context"

	"user-service/internal/dal"
	"user-service/internal/models"
	"user-service/internal/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	Register(ctx context.Context, req *models.RegisterRequest) (*mongo.InsertOneResult, error)
	Authorize(ctx context.Context, req *models.LoginRequest) (string, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	UpdatePasswordById(ctx context.Context, id string, password string) error
	DeleteUser(ctx context.Context, id string) error
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

func (s *userService) Authorize(ctx context.Context, req *models.LoginRequest) (string, error) {
	// login validation
	user, err := s.userRepository.GetUserByEmail(ctx, req.Email)
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

func (s *userService) Register(ctx context.Context, req *models.RegisterRequest) (*mongo.InsertOneResult, error) {
	// email validation
	if !utils.ValidateEmail(req.Email) {
		return nil, ErrInvalidEmail
	}

	// check for uniqueness
	existingUser, err := s.userRepository.GetUserByEmail(ctx, req.Email)
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

	newUser, err := s.userRepository.CreateUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *userService) GetUser(ctx context.Context, id string) (*models.User, error) {
	user, err := s.userRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	users, err := s.userRepository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	return s.userRepository.DeleteUserById(ctx, id)
}

func (s *userService) UpdatePasswordById(ctx context.Context, id string, password string) error {
	// password validation
	if !utils.ValidatePassword(password) {
		return ErrInvalidPassword
	}

	return s.userRepository.UpdatePasswordById(ctx, id, password)
}
