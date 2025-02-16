package dal

import (
	"context"

	"reservation-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type PromotionRepository interface {
	GetAll(ctx context.Context) ([]models.Promotion, error)
	GetById(ctx context.Context, id string) (*models.Promotion, error)
	Add(ctx context.Context, promotion models.Promotion) (*mongo.InsertOneResult, error)
	Update(ctx context.Context, id string, promotion models.Promotion) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (*mongo.DeleteResult, error)
}

type promotionRepository struct {
	db *mongo.Database
}

func NewPromotionRepository(db *mongo.Database) PromotionRepository {
	return &promotionRepository{db: db}
}

func (r *promotionRepository) GetAll(ctx context.Context) ([]models.Promotion, error) {
	coll := r.db.Collection("promotions")
	filter := bson.D{{}}
	var promotions []models.Promotion

	result, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for result.Next(ctx) {
		var promotion models.Promotion
		err = result.Decode(&promotion)
		if err != nil {
			return nil, err
		}

		promotions = append(promotions, promotion)
	}

	return promotions, nil
}

func (r *promotionRepository) GetById(ctx context.Context, id string) (*models.Promotion, error) {
	coll := r.db.Collection("promotions")
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var promotion models.Promotion
	err = coll.FindOne(ctx, bson.M{"_id": ObjID}).Decode(&promotion)
	if err != nil {
		return nil, err
	}

	return &promotion, nil
}

func (r *promotionRepository) Add(ctx context.Context, promotion models.Promotion) (*mongo.InsertOneResult, error) {
	coll := r.db.Collection("promotions")

	result, err := coll.InsertOne(ctx, promotion)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *promotionRepository) Update(ctx context.Context, id string, promotion models.Promotion) (*mongo.UpdateResult, error) {
	coll := r.db.Collection("promotions")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	resDoc := bson.M{
		"_id":        objID,
		"code":       promotion.Code,
		"discount":   promotion.Discount,
		"valid_from": promotion.ValidFrom,
		"valid_to":   promotion.ValidTo,
		"applies_to": promotion.AppliesTo,
	}

	result, err := coll.ReplaceOne(ctx, bson.M{"_id": objID}, resDoc)
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, ErrNotFoundById
	}

	return result, nil
}

func (r *promotionRepository) Delete(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	coll := r.db.Collection("promotions")
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res, err := coll.DeleteOne(ctx, bson.M{"_id": ObjID})
	if err != nil {
		return nil, err
	}
	if res.DeletedCount == 0 {
		return nil, ErrNotFoundById
	}
	return res, nil
}
