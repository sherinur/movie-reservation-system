package service

import "errors"

var (
	ErrBadRequest = errors.New("bad request")
	ErrInvalidId  = errors.New("invalid id")
)
