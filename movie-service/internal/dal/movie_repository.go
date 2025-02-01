package dal

import (
	"context"
	"encoding/json"

	"movie-service/internal/models"
	"movie-service/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieRepository interface {
	AddBatchOfMovie(movielist []models.Movie) (*mongo.InsertManyResult, error)
	AddMovie(movie models.Movie) (*mongo.InsertOneResult, error)
	GetAllMovie() ([]byte, error)
	GetMovieById(id string) ([]byte, error)
	UpdateMovieById(id string, movie *models.Movie) (*mongo.UpdateResult, error)
	DeleteMovieById(id string) (*mongo.DeleteResult, error)
	DeleteAllMovie() (*mongo.DeleteResult, error)
}

type movieRepository struct {
	db *mongo.Database
}

func NewMovieRepository(db *mongo.Database) MovieRepository {
	return &movieRepository{
		db: db,
	}
}

func (r *movieRepository) AddBatchOfMovie(movielist []models.Movie) (*mongo.InsertManyResult, error) {
	col := r.db.Collection("movie")

	doc := []interface{}{}
	for _, movie := range movielist {
		bsonDoc, err := utils.ConvertToBsonD(movie)
		if err != nil {
			return nil, err
		}

		doc = append(doc, bsonDoc)
	}

	res, err := col.InsertMany(context.TODO(), doc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *movieRepository) AddMovie(movie models.Movie) (*mongo.InsertOneResult, error) {
	col := r.db.Collection("movie")

	insertRes, err := col.InsertOne(context.TODO(), movie)
	if err != nil {
		return nil, err
	}

	return insertRes, nil
}

func (r *movieRepository) GetAllMovie() ([]byte, error) {
	var result []bson.M
	col := r.db.Collection("movie")

	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &result)
	if err != nil {
		return nil, err
	}

	data, err := json.MarshalIndent(result, "", "")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *movieRepository) GetMovieById(id string) ([]byte, error) {
	var result bson.M
	col := r.db.Collection("movie")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = col.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	data, err := json.MarshalIndent(result, "", "")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *movieRepository) UpdateMovieById(id string, movie *models.Movie) (*mongo.UpdateResult, error) {
	col := r.db.Collection("movie")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update, err := utils.ConvertToBsonD(movie)
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

	deleteres, err := col.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}})
	if err != nil {
		return nil, err
	}

	return deleteres, nil
}

func (r *movieRepository) DeleteAllMovie() (*mongo.DeleteResult, error) {
	col := r.db.Collection("movie")

	res, err := col.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	return res, nil
}
