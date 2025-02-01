package handler

import "errors"

var ErrMethodNotPost = errors.New("only POST method is allowed")
var ErrMethodNotPut = errors.New("only PUT method is supported")
var ErrMethodNotDelete = errors.New("only DELETE method is supported")
var ErrEmptyData = errors.New("incoming data must be entered")
var ErrNoId = errors.New("missing update ID")
var ErrNotAutorized = errors.New("not autorized")
