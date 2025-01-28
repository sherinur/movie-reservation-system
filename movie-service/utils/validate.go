package utils

import (
	"strings"

	"movie-service/internal/models"
)

func ValidateScreening(screening models.Session) error {
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
	for _, screening := range hall.Session {
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

func ValidateMovie(movie models.Movie) error {
	if strings.TrimSpace(movie.Title) == "" {
		return ErrMovieTitleEmpty
	}
	if strings.TrimSpace(movie.Genre) == "" {
		return ErrMovieGenreEmpty
	}
	if strings.TrimSpace(movie.Description) == "" {
		return ErrMovieDescriptionEmpty
	}
	if strings.TrimSpace(movie.PosterImage) == "" {
		return ErrMoviePosterEmpty
	}
	if movie.Duration <= 0 {
		return ErrMovieDurationInvalid
	}
	if strings.TrimSpace(movie.Language) == "" {
		return ErrMovieLanguageEmpty
	}
	if strings.TrimSpace(movie.ReleaseDate) == "" {
		return ErrMovieReleaseDateEmpty
	}
	if strings.TrimSpace(movie.Rating) == "" {
		return ErrMovieRatingEmpty
	}
	if strings.TrimSpace(movie.PGrating) == "" {
		return ErrMoviePGEmpty
	}
	if strings.TrimSpace(movie.Production) == "" {
		return ErrMovieProductionEmpty
	}
	if strings.TrimSpace(movie.Producer) == "" {
		return ErrMovieProducerEmpty
	}
	if strings.TrimSpace(movie.Status) == "" {
		return ErrMovieStatusEmpty
	}

	return nil
}
