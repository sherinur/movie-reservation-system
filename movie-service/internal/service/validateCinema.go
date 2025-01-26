package service

import (
	"errors"
	"movie-service/internal/models"
	"strings"
)

func ValidateScreening(screening models.Screening) error {
	if strings.TrimSpace(screening.MovieID) == "" {
		return errors.New("screening movie ID cannot be empty")
	}
	err := ValidateMovie(screening.Movie)
	if err != nil {
		return err
	}
	return nil
}

func ValidateSeat(seat models.Seat) error {
	if strings.TrimSpace(seat.Row) == "" || strings.TrimSpace(seat.Column) == "" || strings.TrimSpace(seat.Status) == "" {
		return errors.New("seat must have non-empty row, column, and status")
	}
	return nil
}

func ValidateHall(hall models.Hall) error {
	if hall.Number == 0 {
		return errors.New("hall number cannot be zero")
	}
	if hall.RowCount == 0 || hall.ColumnCount == 0 {
		return errors.New("hall must have non-zero row and column count")
	}
	if len(hall.Seats) == 0 {
		return errors.New("hall must have seats")
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
		return errors.New("cinema name or address cannot be empty")
	}
	if len(cinema.HallList) == 0 {
		return errors.New("cinema must have at least one hall")
	}

	for _, hall := range cinema.HallList {
		if err := ValidateHall(hall); err != nil {
			return err
		}
	}
	return nil
}
