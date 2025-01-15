package dal

import (
	"context"

	"user-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) (*mongo.InsertOneResult, error)
	FindUserByEmail(email string) (*models.User, error)
	IsEmailExists(email string) (bool, error)
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	// getting cursor of users from mongo
	cur, err := r.db.Collection("users").Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var users []models.User

	// cursor iteration
	for cur.Next(context.Background()) {
		user := models.User{}

		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	// handling cursor err
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	coll := r.db.Collection("users")

	result, err := coll.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userRepository) FindUserByEmail(email string) (*models.User, error) {
	coll := r.db.Collection("users")

	filter := bson.D{{Key: "email", Value: email}}

	var user models.User
	err := coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) IsEmailExists(email string) (bool, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var result bson.M
	err := r.db.Collection("users").FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
