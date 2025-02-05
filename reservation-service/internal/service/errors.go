package service

import "errors"

var ErrEmptyData = errors.New("not provided all data")
var ErrNoId = errors.New("id is empty")
var ErrPaidReservation = errors.New("could not pay already paid reservation")
var ErrWrongUser = errors.New("attemping to interact with someone else's reservation")
