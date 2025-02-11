package dal

import (
	"context"
	"strings"

	"movie-service/internal/models"
	"movie-service/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CinemaRepository interface {
	AddCinema(cinema models.Cinema) (*mongo.InsertOneResult, error)
	GetAllCinema() ([]models.Cinema, error)
	GetCinemaById(id string) (*models.Cinema, error)
	UpdateCinemaById(id string, movie *models.Cinema) (*mongo.UpdateResult, error)
	DeleteCinemaById(id string) (*mongo.DeleteResult, error)
	DeleteAllCinema() (*mongo.DeleteResult, error)

	AddHall(id string, hall models.Hall) (*mongo.UpdateResult, error)
	GetHall(cinemaID string, hallNumber int) (*models.Hall, error)
	GetAllHall(cinemaID string) ([]models.Hall, error)
	DeleteHall(cinemaID string, hallNumber int) (*mongo.UpdateResult, error)
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

	if len(strings.TrimSpace(cinema.ID)) == 0 {
		cinema.ID = utils.GenerateID()
	}

	for i := range cinema.HallList {
		cinema.HallList[i].AvailableSeats = len(cinema.HallList[i].Seats)
	}

	res, err := col.InsertOne(context.TODO(), cinema)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *cinemaRepository) GetAllCinema() ([]models.Cinema, error) {
	col := r.db.Collection("cinema")
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var cinema_list []models.Cinema
	err = cursor.All(context.TODO(), &cinema_list)
	if err != nil {
		return nil, err
	}

	return cinema_list, nil
}

func (r *cinemaRepository) GetCinemaById(id string) (*models.Cinema, error) {
	col := r.db.Collection("cinema")

	var cinema models.Cinema
	err := col.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&cinema)
	if err != nil {
		return nil, err
	}

	return &cinema, nil
}

func (r *cinemaRepository) UpdateCinemaById(id string, cinema *models.Cinema) (*mongo.UpdateResult, error) {
	col := r.db.Collection("cinema")

	if len(strings.TrimSpace(cinema.ID)) == 0 {
		cinema.ID = id
	}

	for i := range cinema.HallList {
		cinema.HallList[i].AvailableSeats = len(cinema.HallList[i].Seats)
	}

	update, err := utils.ConvertToBsonD(cinema)
	if err != nil {
		return nil, err
	}

	res, err := col.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: id}}, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *cinemaRepository) DeleteCinemaById(id string) (*mongo.DeleteResult, error) {
	col := r.db.Collection("cinema")

	deleteres, err := col.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
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

func (r *cinemaRepository) AddHall(cinemaID string, hall models.Hall) (*mongo.UpdateResult, error) {
	col := r.db.Collection("cinema")

	hall.AvailableSeats = len(hall.Seats)
	update := bson.M{
		"$push": bson.M{
			"hall_list": hall,
		},
	}

	updateResult, err := col.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: cinemaID}}, update)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (r *cinemaRepository) GetHall(cinemaID string, hallNumber int) (*models.Hall, error) {
	col := r.db.Collection("cinema")

	filter := bson.M{"hall_list.number": hallNumber, "_id": cinemaID}
	projection := bson.M{"hall_list.$": 1}

	hall_list := models.Hall_list{}
	err := col.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&hall_list)
	if err != nil {
		return nil, err
	}

	if len(hall_list.Hall_list) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	hall := hall_list.Hall_list[0]
	return &hall, nil
}

func (r *cinemaRepository) GetAllHall(cinemaID string) ([]models.Hall, error) {
	var halls []models.Hall
	col := r.db.Collection("cinema")

	filter := bson.M{"_id": cinemaID}
	projection := bson.M{"hall_list": 1, "_id": 0}

	var result models.Hall_list
	err := col.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return halls, err
	}

	return result.Hall_list, nil
}

func (r *cinemaRepository) DeleteHall(cinemaID string, hallNumber int) (*mongo.UpdateResult, error) {
	col := r.db.Collection("cinema")

	filter := bson.M{"_id": cinemaID}
	update := bson.M{
		"$pull": bson.M{
			"hall_list": bson.M{"number": hallNumber},
		},
	}

	updateResult, err := col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}
