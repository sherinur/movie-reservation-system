package service

import "errors"

var (
	ErrWrongPassword   = errors.New("wrong password")
	ErrNoUser          = errors.New("user not found")
	ErrInvalidEmail    = errors.New("email is not valid")
	ErrInvalidPassword = errors.New("password is not valid")
	ErrInvalidUsername = errors.New("username is not valid")
	ErrUserExists      = errors.New("the user is already exists")
)
