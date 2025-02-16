package dal

import (
	"context"

	"reservation-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepository interface {
	GetByUserId(ctx context.Context, userId string) ([]models.Payment, error)
	GetById(ctx context.Context, id string) (*models.Payment, error)
	Add(ctx context.Context, payment models.Payment) (*mongo.InsertOneResult, error)
	Update(ctx context.Context, id string, payment models.Payment) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (*mongo.DeleteResult, error)
}

type paymentRepository struct {
	db *mongo.Database
}

func NewPaymentRepository(db *mongo.Database) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) GetByUserId(ctx context.Context, userId string) ([]models.Payment, error) {
	coll := r.db.Collection("payments")
	//ObjID, err := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "user_id", Value: userId}}
	var payments []models.Payment

	result, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for result.Next(ctx) {
		var payment models.Payment
		err = result.Decode(&payment)
		if err != nil {
			return nil, err
		}

		payments = append(payments, payment)
	}

	return payments, nil
}

func (r *paymentRepository) GetById(ctx context.Context, id string) (*models.Payment, error) {
	coll := r.db.Collection("payments")
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var payment models.Payment
	err = coll.FindOne(ctx, bson.M{"_id": ObjID}).Decode(&payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *paymentRepository) Add(ctx context.Context, process models.Payment) (*mongo.InsertOneResult, error) {
	coll := r.db.Collection("payments")

	result, err := coll.InsertOne(ctx, process)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *paymentRepository) Update(ctx context.Context, id string, payment models.Payment) (*mongo.UpdateResult, error) {
	coll := r.db.Collection("payments")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	resDoc := bson.M{
		"_id":             objID,
		"user_id":         payment.UserId,
		"reservation_id":  payment.ReservationId,
		"payment_price":   payment.PaymentPrice,
		"payment_method":  payment.PaymentMethod,
		"status":          payment.Status,
		"trasaction_date": payment.TransactionDate,
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

func (r *paymentRepository) Delete(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	coll := r.db.Collection("reservations")
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
