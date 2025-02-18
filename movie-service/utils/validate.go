package utils

import (
	"strings"
	"time"

	"movie-service/internal/models"
)

func ValidateSeat(seat models.Seat) error {
	switch {
	case strings.TrimSpace(seat.Row) == "", strings.TrimSpace(seat.Column) == "", strings.TrimSpace(seat.Status) == "":
		return ErrSeatValidation
	}
	return nil
}

func ValidateHall(hall models.Hall) error {
	switch {
	case hall.Number < 1:
		return ErrHallNumberZero
	case len(hall.Seats) == 0:
		return ErrHallNoSeats
	}

	for _, seat := range hall.Seats {
		if err := ValidateSeat(seat); err != nil {
			return err
		}
	}
	return nil
}

func ValidateCinema(cinema models.Cinema) error {
	switch {
	case strings.TrimSpace(cinema.Name) == "",
		strings.TrimSpace(cinema.Address) == "",
		strings.TrimSpace(cinema.City) == "":
		return ErrCinemaNameAddress
	}
	if strings.TrimSpace(cinema.Name) == "" || strings.TrimSpace(cinema.Address) == "" || strings.TrimSpace(cinema.City) == "" {
		return ErrCinemaNameAddress
	}

	if cinema.Rating < 0 {
		return ErrValidRating
	}

	// if len(cinema.HallList) == 0 {
	// 	return ErrCinemaNoHalls
	// }

	// for _, hall := range cinema.HallList {
	// 	if err := ValidateHall(hall); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func ValidateMovie(movie models.Movie) error {
	pgratinglist := map[string]struct{}{
		"G":     {},
		"PG":    {},
		"PG-13": {},
		"R":     {},
		"NC-17": {},
	}

	_, ValidPGrating := pgratinglist[movie.PGrating]
	switch {
	case strings.TrimSpace(movie.Title) == "":
		return ErrMovieTitleEmpty
	case strings.TrimSpace(movie.Genre) == "":
		return ErrMovieGenreEmpty
	case strings.TrimSpace(movie.Description) == "":
		return ErrMovieDescriptionEmpty
	case strings.TrimSpace(movie.PosterImage) == "":
		return ErrMoviePosterEmpty
	case movie.Duration <= 0:
		return ErrMovieDurationInvalid
	case strings.TrimSpace(movie.Language) == "":
		return ErrMovieLanguageEmpty
	case strings.TrimSpace(movie.ReleaseDate) == "":
		return ErrMovieReleaseDateEmpty
	case strings.TrimSpace(movie.Rating) == "":
		return ErrMovieRatingEmpty
	case !ValidPGrating:
		return ErrInvalidMoviePGrating
	case strings.TrimSpace(movie.Production) == "":
		return ErrMovieProductionEmpty
	case strings.TrimSpace(movie.Producer) == "":
		return ErrMovieProducerEmpty
	case strings.TrimSpace(movie.Status) == "":
		return ErrMovieStatusEmpty
	}

	return nil
}

func ValidateSesesion(session models.Session) error {
	switch {
	case strings.TrimSpace(session.MovieID) == "":
		return ErrSessionMovieIDEmpty
	case strings.TrimSpace(session.CinemaID) == "":
		return ErrSessionCinemaIDEmpty
	case session.HallNumber < 0:
		return ErrSessionInvalidHallNumber
	case session.StartTime.Before(time.Now()):
		return ErrSessionStartTimeInvalid
	case session.EndTime.Before(session.StartTime):
		return ErrSessionEndTimeInvalid
	}

	return nil
}
