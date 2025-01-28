package dal

import "errors"

var ErrNotFoundById = errors.New("no process/reservation found with the given ID")
