package dal

import (
	"context"
	"strings"

	"movie-service/internal/models"
	"movie-service/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieRepository interface {
	AddBatchOfMovie(movielist []models.Movie) (*mongo.InsertManyResult, error)
	AddMovie(movie models.Movie) (*mongo.InsertOneResult, error)
	GetAllMovie() ([]models.Movie, error)
	GetMovieById(id string) (*models.Movie, error)
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
		if len(strings.TrimSpace(movie.ID)) == 0 {
			movie.ID = utils.GenerateID()
		}
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

	if len(strings.TrimSpace(movie.ID)) == 0 {
		movie.ID = utils.GenerateID()
	}

	insertRes, err := col.InsertOne(context.TODO(), movie)
	if err != nil {
		return nil, err
	}

	return insertRes, nil
}

func (r *movieRepository) GetAllMovie() ([]models.Movie, error) {
	var movies []models.Movie
	col := r.db.Collection("movie")

	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *movieRepository) GetMovieById(id string) (*models.Movie, error) {
	var movie models.Movie
	col := r.db.Collection("movie")

	err := col.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&movie)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *movieRepository) UpdateMovieById(id string, movie *models.Movie) (*mongo.UpdateResult, error) {
	col := r.db.Collection("movie")

	if len(strings.TrimSpace(movie.ID)) == 0 {
		movie.ID = id
	}

	res, err := col.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: id}}, bson.D{{Key: "$set", Value: movie}})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *movieRepository) DeleteMovieById(id string) (*mongo.DeleteResult, error) {
	col := r.db.Collection("movie")

	deleteres, err := col.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
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
