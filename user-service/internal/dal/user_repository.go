package dal

import (
	"context"

	"user-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	IsEmailExists(ctx context.Context, email string) (bool, error)
	UpdatePasswordById(ctx context.Context, id string, password string) error
	DeleteUserById(ctx context.Context, id string) error
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	// getting a cursor of users from mongo
	cur, err := r.db.Collection("users").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var users []models.User

	// cursor iteration
	for cur.Next(ctx) {
		user := models.User{}

		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	// handling a cursor err
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	coll := r.db.Collection("users")

	result, err := coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: objID}}

	var user models.User
	err = r.db.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var user models.User
	err := r.db.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) IsEmailExists(ctx context.Context, email string) (bool, error) {
	filter := bson.D{{Key: "email", Value: email}}

	var result bson.M
	err := r.db.Collection("users").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *userRepository) DeleteUserById(ctx context.Context, id string) error {
	filter := bson.D{{Key: "email", Value: id}}

	_, err := r.db.Collection("users").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdatePasswordById(ctx context.Context, id string, password string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	set := bson.D{{Key: "password", Value: password}}

	update := bson.D{{Key: "$set", Value: set}}
	_, err = r.db.Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
