package dal

import (
	"context"
	"encoding/json"
	"movie-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieRepository interface {
	AddMovie(movielist []models.Movie) (*mongo.InsertManyResult, error)
	GetAllMovie() ([]byte, error)
	UpdateMovieById(id string, movie *models.Movie) (*mongo.UpdateResult, error)
	DeleteMovieById(id string) (*mongo.DeleteResult, error)
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

	doc := []interface{}{}
	for _, movie := range movielist {
		bsonDoc, err := ConvertToBsonD(movie)
		if err != nil {
			return nil, err
		}

		doc = append(doc, bsonDoc)
	}

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

func (r *movieRepository) UpdateMovieById(id string, movie *models.Movie) (*mongo.UpdateResult, error) {
	col := r.db.Collection("movie")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update, err := ConvertToBsonD(movie)
	if err != nil {
		return nil, err
	}

	res, err := col.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}}, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *movieRepository) DeleteMovieById(id string) (*mongo.DeleteResult, error) {
	col := r.db.Collection("movie")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res, err := col.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}})
	if err != nil {
		return nil, err
	}

	return res, nil

}

// Convert struct to bson.D format
func ConvertToBsonD(movie interface{}) (interface{}, error) {
	bsonData, err := bson.Marshal(movie)
	if err != nil {
		return nil, err
	}

	var bsonDoc bson.D
	if err := bson.Unmarshal(bsonData, &bsonDoc); err != nil {
		return nil, err
	}

	return bsonDoc, nil
}
