package service

import (
	"context"
	"time"

	"reservation-service/internal/dal"
	"reservation-service/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentService interface {
	GetPayments(ctx context.Context, userId string) ([]models.Payment, error)
	GetPayment(ctx context.Context, id string) (*models.Payment, error)
	AddPayment(ctx context.Context, booking models.PaymentRequest) (*mongo.InsertOneResult, error)
	UpdatePayment(ctx context.Context, id string, paying models.Payment) (*mongo.UpdateResult, error)
	DeletePayment(ctx context.Context, id, userID string) (*mongo.DeleteResult, error)
}

type paymentService struct {
	paymentRepository dal.PaymentRepository
}

func NewPaymentService(r dal.PaymentRepository) PaymentService {
	return &paymentService{
		paymentRepository: r,
	}
}

func (s *paymentService) GetPayments(ctx context.Context, userId string) ([]models.Payment, error) {
	return s.paymentRepository.GetByUserId(ctx, userId)
}

func (s *paymentService) GetPayment(ctx context.Context, id string) (*models.Payment, error) {
	if id == "" {
		return nil, ErrNoId
	}

	return s.paymentRepository.GetById(ctx, id)
}

func (s *paymentService) AddPayment(ctx context.Context, requestBody models.PaymentRequest) (*mongo.InsertOneResult, error) {
	if requestBody.UserId == "" || requestBody.ReservationId == "" || requestBody.PaymentPrice <= 0 || requestBody.PaymentMethod == "" {
		return nil, ErrEmptyData
	}

	payment := models.Payment{
		UserId:          requestBody.UserId,
		ReservationId:   requestBody.ReservationId,
		PaymentPrice:    requestBody.PaymentPrice,
		PaymentMethod:   requestBody.PaymentMethod,
		Status:          requestBody.Status,
		TransactionDate: time.Now(),
	}

	result, err := s.paymentRepository.Add(ctx, payment)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *paymentService) UpdatePayment(ctx context.Context, id string, requestBody models.Payment) (*mongo.UpdateResult, error) {
	if id == "" {
		return nil, ErrNoId
	}

	process, err := s.paymentRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if process.UserId != requestBody.UserId {
		return nil, ErrWrongUser
	}

	payment := models.Payment{
		UserId:          requestBody.UserId,
		ReservationId:   requestBody.ReservationId,
		PaymentPrice:    requestBody.PaymentPrice,
		PaymentMethod:   requestBody.PaymentMethod,
		Status:          requestBody.Status,
		TransactionDate: time.Now(),
	}

	result, err := s.paymentRepository.Update(ctx, id, payment)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *paymentService) DeletePayment(ctx context.Context, id, userID string) (*mongo.DeleteResult, error) {
	if id == "" {
		return nil, ErrNoId
	}

	payment, err := s.paymentRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if payment.UserId != userID {
		return nil, ErrWrongUser
	}

	return s.paymentRepository.Delete(ctx, id)
}
