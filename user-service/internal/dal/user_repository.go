package dal

import (
	"context"

	"user-service/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUsers() *mongo.Collection
	CreateUser(user *models.User) (*mongo.InsertOneResult, error)
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUsers() *mongo.Collection {
	return r.db.Collection("users")
}

func (r *userRepository) CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
	collection := r.db.Collection("users")

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
