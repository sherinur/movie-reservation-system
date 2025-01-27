package service

import "errors"

var (
	ErrBadRequest = errors.New("bad request")
	ErrInvalidId  = errors.New("invalid id")

	// Cinema errors
	ErrSeatValidation       = errors.New("seat must have non-empty row, column, and status")
	ErrHallNumberZero       = errors.New("hall number cannot be zero")
	ErrValidRating          = errors.New("rating cannot be less than zero")
	ErrHallDimensionInvalid = errors.New("hall must have non-zero row and column count")
	ErrScreeningMovieID     = errors.New("screening movie ID cannot be empty")
	ErrHallNoSeats          = errors.New("hall must have seats")
	ErrCinemaNameAddress    = errors.New("cinema name or address cannot be empty")
	ErrCinemaNoHalls        = errors.New("cinema must have at least one hall")

	// Movie errors
	ErrMovieTitleEmpty       = errors.New("movie title cannot be empty")
	ErrMovieGenreEmpty       = errors.New("movie genre cannot be empty")
	ErrMovieDescriptionEmpty = errors.New("movie description cannot be empty")
	ErrMoviePosterEmpty      = errors.New("movie poster image cannot be empty")
	ErrMovieLanguageEmpty    = errors.New("movie language cannot be empty")
	ErrMovieReleaseDateEmpty = errors.New("movie release date cannot be empty")
	ErrMovieRatingEmpty      = errors.New("movie rating cannot be empty")
	ErrMoviePGEmpty          = errors.New("movie PG rating cannot be empty")
	ErrMovieProductionEmpty  = errors.New("movie production cannot be empty")
	ErrMovieProducerEmpty    = errors.New("movie producer cannot be empty")
	ErrMovieStatusEmpty      = errors.New("movie status cannot be empty")
	ErrMovieDurationInvalid  = errors.New("movie duration must be greater than zero")

	BadRequestMovieErrors = map[error]struct{}{
		ErrInvalidId:             {},
		ErrBadRequest:            {},
		ErrMovieTitleEmpty:       {},
		ErrMovieGenreEmpty:       {},
		ErrMovieDescriptionEmpty: {},
		ErrMoviePosterEmpty:      {},
		ErrMovieLanguageEmpty:    {},
		ErrMovieReleaseDateEmpty: {},
		ErrMovieRatingEmpty:      {},
		ErrMoviePGEmpty:          {},
		ErrMovieProductionEmpty:  {},
		ErrMovieProducerEmpty:    {},
		ErrMovieStatusEmpty:      {},
		ErrMovieDurationInvalid:  {},
	}

	BadRequestCinemaErrors = map[error]struct{}{
		ErrInvalidId:            {},
		ErrValidRating:          {},
		ErrBadRequest:           {},
		ErrSeatValidation:       {},
		ErrHallNumberZero:       {},
		ErrHallDimensionInvalid: {},
		ErrScreeningMovieID:     {},
		ErrHallNoSeats:          {},
		ErrCinemaNameAddress:    {},
		ErrCinemaNoHalls:        {},
	}
)
