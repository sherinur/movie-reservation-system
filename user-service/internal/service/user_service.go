package service

import "user-service/internal/dal"

type UserService interface{}

type userService struct {
	userRepository dal.UserRepository
}

func NewUserService(r *dal.UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}
