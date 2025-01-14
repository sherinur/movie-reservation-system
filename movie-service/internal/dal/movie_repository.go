package dal

import "go.mongodb.org/mongo-driver/mongo"

type MovieRepository interface {
}

type movieRepository struct {
	db *mongo.Database
}

func NewMovieRepository(db *mongo.Database) MovieRepository {
	return &movieRepository{
		db: db,
	}
}
