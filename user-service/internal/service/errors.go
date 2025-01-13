package service

import "errors"

var (
	ErrWrongPassword = errors.New("wrong password")
	ErrNoUser        = errors.New("user not found")
)
