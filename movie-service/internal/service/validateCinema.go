package service

import (
	"strings"

	"movie-service/internal/models"
)

func ValidateScreening(screening models.Screening) error {
	if strings.TrimSpace(screening.MovieID) == "" {
		return ErrScreeningMovieID
	}
	err := ValidateMovie(screening.Movie)
	if err != nil {
		return err
	}
	return nil
}

func ValidateSeat(seat models.Seat) error {
	if strings.TrimSpace(seat.Row) == "" || strings.TrimSpace(seat.Column) == "" || strings.TrimSpace(seat.Status) == "" {
		return ErrSeatValidation
	}
	return nil
}

func ValidateHall(hall models.Hall) error {
	if hall.Number == 0 {
		return ErrHallNumberZero
	}
	if hall.RowCount == 0 || hall.ColumnCount == 0 {
		return ErrHallDimensionInvalid
	}
	if len(hall.Seats) == 0 {
		return ErrHallNoSeats
	}

	for _, seat := range hall.Seats {
		if err := ValidateSeat(seat); err != nil {
			return err
		}
	}
	for _, screening := range hall.Screenings {
		if err := ValidateScreening(screening); err != nil {
			return err
		}
	}
	return nil
}

func ValidateCinema(cinema models.Cinema) error {
	if strings.TrimSpace(cinema.Name) == "" || strings.TrimSpace(cinema.Address) == "" {
		return ErrCinemaNameAddress
	}

	if cinema.Rating < 0 {
		return ErrValidRating
	}

	if len(cinema.HallList) == 0 {
		return ErrCinemaNoHalls
	}

	for _, hall := range cinema.HallList {
		if err := ValidateHall(hall); err != nil {
			return err
		}
	}
	return nil
}
