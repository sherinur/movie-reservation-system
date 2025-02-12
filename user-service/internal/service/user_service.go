package service

import (
	"context"

	"user-service/internal/dal"
	"user-service/internal/models"
	"user-service/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *models.RegisterRequest) error
	Authorize(ctx context.Context, req *models.LoginRequest) (*models.User, error)
	Refresh(ctx context.Context, refreshToken string) error
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	UpdatePasswordById(ctx context.Context, id string, password string) error
	DeleteUser(ctx context.Context, id string) error
}

type userService struct {
	userRepository dal.UserRepository
	hashCost       int
}

func NewUserService(r dal.UserRepository) UserService {
	return &userService{
		userRepository: r,
		hashCost:       bcrypt.DefaultCost,
	}
}

func (s *userService) Authorize(ctx context.Context, req *models.LoginRequest) (*models.User, error) {
	// login validation
	user, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrNoUser
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrWrongPassword
		}
		return nil, err
	}

	return user, nil
}

func (s *userService) Register(ctx context.Context, req *models.RegisterRequest) error {
	// email validation
	if !utils.ValidateEmail(req.Email) {
		return ErrInvalidEmail
	}

	// check for uniqueness
	existingUser, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	} else if existingUser != nil {
		return ErrUserExists
	}

	// password validation
	if !utils.ValidatePassword(req.Password) {
		return ErrInvalidPassword
	}

	// username validation
	if !utils.ValidateUsername(req.Username) {
		return ErrInvalidUsername
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), s.hashCost)
	if err != nil {
		return err
	}

	// create a new user
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashPassword),
		Role:     "User",
	}

	_, err = s.userRepository.CreateUser(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) Refresh(ctx context.Context, refreshToken string) error {
	return nil
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

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), s.hashCost)
	if err != nil {
		return err
	}

	return s.userRepository.UpdatePasswordById(ctx, id, string(hashPassword))
}
