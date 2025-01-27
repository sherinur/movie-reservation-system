package service

import (
	"strings"

	"movie-service/internal/models"
)

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
