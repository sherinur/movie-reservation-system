package dal

import (
	"context"
	"movie-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieRepository interface {
	AddMovie(movielist []interface{}) (*mongo.InsertManyResult, error)
	GetMovie() (*models.MovieList, error)
	UpdateMovie()
	DeleteMovie()
}

type movieRepository struct {
	db *mongo.Database
}

func NewMovieRepository(db *mongo.Database) MovieRepository {
	return &movieRepository{
		db: db,
	}
}

func (r *movieRepository) AddMovie(movielist []interface{}) (*mongo.InsertManyResult, error) {
	collection := r.db.Collection("movie")

	res, err := collection.InsertMany(context.TODO(), movielist)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *movieRepository) GetMovie() (*models.MovieList, error) {
	collection := r.db.Collection("movie")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	data := models.MovieList{List: []models.Movie{}}
	err = cursor.All(context.TODO(), data.List)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *movieRepository) UpdateMovie() {

}

func (r *movieRepository) DeleteMovie() {

}
