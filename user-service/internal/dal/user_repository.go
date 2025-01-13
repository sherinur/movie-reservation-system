package dal

import (
	"context"

	"user-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUsers() *mongo.Collection
	CreateUser(user *models.User) (*mongo.InsertOneResult, error)
	FindUserByEmail(email string) (*models.User, error)
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
	// cur, err := r.db.Collection("users").Find(context.Background())
	// return cur
	return nil
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

	var user *models.User
	err := coll.FindOne(context.TODO(), filter).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
