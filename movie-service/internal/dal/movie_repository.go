package dal

import (
	"context"
	"encoding/json"
	"movie-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieRepository interface {
	AddMovie(movielist []models.Movie) (*mongo.InsertManyResult, error)
	GetAllMovie() ([]byte, error)
	UpdateMovie() error
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

func (r *movieRepository) AddMovie(movielist []models.Movie) (*mongo.InsertManyResult, error) {
	collection := r.db.Collection("movie")

	doc := ConvertToDoc(movielist)
	res, err := collection.InsertMany(context.TODO(), doc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *movieRepository) GetAllMovie() ([]byte, error) {
	col := r.db.Collection("movie")
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var result []bson.M
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		return nil, err
	}

	data, err := json.MarshalIndent(result, "", "")
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *movieRepository) UpdateMovie() error {
	return nil
}

func (r *movieRepository) DeleteMovie() {

}

func ConvertToDoc(movielist []models.Movie) []interface{} {
	documents := []interface{}{}
	for _, movie := range movielist {
		documents = append(documents, bson.D{
			{Key: "Title", Value: movie.Title},
		})
	}

	return documents
}
