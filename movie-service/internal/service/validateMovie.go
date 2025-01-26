package service

import (
	"errors"
	"movie-service/internal/models"
	"strings"
)

func ValidateMovie(movie models.Movie) error {
	if strings.TrimSpace(movie.Title) == "" {
		return errors.New("movie title cannot be empty")
	}
	if strings.TrimSpace(movie.Genre) == "" {
		return errors.New("movie genre cannot be empty")
	}
	if strings.TrimSpace(movie.Description) == "" {
		return errors.New("movie description cannot be empty")
	}
	if strings.TrimSpace(movie.PosterImage) == "" {
		return errors.New("movie poster image cannot be empty")
	}
	if strings.TrimSpace(movie.Language) == "" {
		return errors.New("movie language cannot be empty")
	}
	if strings.TrimSpace(movie.ReleaseDate) == "" {
		return errors.New("movie release date cannot be empty")
	}
	if strings.TrimSpace(movie.Rating) == "" {
		return errors.New("movie rating cannot be empty")
	}
	if strings.TrimSpace(movie.PGrating) == "" {
		return errors.New("movie PG rating cannot be empty")
	}
	if strings.TrimSpace(movie.Production) == "" {
		return errors.New("movie production cannot be empty")
	}
	if strings.TrimSpace(movie.Producer) == "" {
		return errors.New("movie producer cannot be empty")
	}
	if strings.TrimSpace(movie.Status) == "" {
		return errors.New("movie status cannot be empty")
	}

	if movie.Duration <= 0 {
		return errors.New("movie duration must be greater than zero")
	}

	return nil
}
