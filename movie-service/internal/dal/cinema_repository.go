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
	AddCinema(cinemalist models.Cinema) (*mongo.InsertOneResult, error)
	AddHall(id string, hall models.Hall) (*mongo.UpdateResult, error)
	GetAllCinema() ([]byte, error)
	GetCinemaById(id string) ([]byte, error)
	UpdateCinemaById(id string, movie *models.Cinema) (*mongo.UpdateResult, error)
	DeleteCinemaById(id string) (*mongo.DeleteResult, error)
	DeleteAllCinema() (*mongo.DeleteResult, error)
}

type cinemaRepository struct {
	db *mongo.Database
}

func NewCinemaRepository(db *mongo.Database) CinemaRepository {
	return &cinemaRepository{
		db: db,
	}
}

func (r *cinemaRepository) AddCinema(cinema models.Cinema) (*mongo.InsertOneResult, error) {
	col := r.db.Collection("cinema")

	bsonDoc, err := utils.ConvertToBsonD(cinema)
	if err != nil {
		return nil, err
	}

	res, err := col.InsertOne(context.TODO(), bsonDoc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *cinemaRepository) AddHall(id string, hall models.Hall) (*mongo.UpdateResult, error) {
	col := r.db.Collection("cinema")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	newHall, err := utils.ConvertToBsonD(hall)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$push": bson.M{
			"hall_list": newHall,
		},
	}

	updateResult, err := col.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}}, update)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
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

func (r *cinemaRepository) GetCinemaById(id string) ([]byte, error) {
	var result bson.M
	col := r.db.Collection("cinema")

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

	deleteres, err := col.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}})
	if err != nil {
		return nil, err
	}

	return deleteres, nil
}

func (r *cinemaRepository) DeleteAllCinema() (*mongo.DeleteResult, error) {
	col := r.db.Collection("cinema")

	deleteres, err := col.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	return deleteres, nil
}
