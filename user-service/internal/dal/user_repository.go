package dal

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository interface{}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}
