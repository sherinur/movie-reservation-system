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

type CinemaRepository interface {
	AddCinema(movielist []models.Cinema) (*mongo.InsertManyResult, error)
	GetAllCinema() ([]byte, error)
	UpdateCinemaById(id string, movie *models.Cinema) (*mongo.UpdateResult, error)
	DeleteCinemaById(id string) (*mongo.DeleteResult, error)
}

type cinemaRepository struct {
	db *mongo.Database
}

func NewCinemaRepository(db *mongo.Database) CinemaRepository {
	return &cinemaRepository{
		db: db,
	}
}

func (r *cinemaRepository) AddCinema(cinemalist []models.Cinema) (*mongo.InsertManyResult, error) {
	col := r.db.Collection("cinema")

	doc := []interface{}{}
	for _, cinema := range cinemalist {
		bsonDoc, err := utils.ConvertToBsonD(cinema)
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

func (r *cinemaRepository) GetAllCinema() ([]byte, error) {
	col := r.db.Collection("cinema")
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

func (r *cinemaRepository) UpdateCinemaById(id string, cinema *models.Cinema) (*mongo.UpdateResult, error) {
	col := r.db.Collection("cinema")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update, err := utils.ConvertToBsonD(cinema)
	if err != nil {
		return nil, err
	}

	res, err := col.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}}, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *cinemaRepository) DeleteCinemaById(id string) (*mongo.DeleteResult, error) {
	col := r.db.Collection("cinema")

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
