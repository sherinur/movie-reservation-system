package service

import (
	"context"

	"reservation-service/internal/dal"
	"reservation-service/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type PromotionService interface {
	GetPromotions(ctx context.Context) ([]models.Promotion, error)
	GetPromotion(ctx context.Context, id string) (*models.Promotion, error)
	AddPromotion(ctx context.Context, requestBody models.PromotionRequest) (*mongo.InsertOneResult, error)
	UpdatePromotion(ctx context.Context, id string, requestBody models.Promotion) (*mongo.UpdateResult, error)
	DeletePromotion(ctx context.Context, id string) (*mongo.DeleteResult, error)
}

type promotionService struct {
	promotionRepository dal.PromotionRepository
}

func NewPromotionService(r dal.PromotionRepository) PromotionService {
	return &promotionService{
		promotionRepository: r,
	}
}

func (s *promotionService) GetPromotions(ctx context.Context) ([]models.Promotion, error) {
	return s.promotionRepository.GetAll(ctx)
}

func (s *promotionService) GetPromotion(ctx context.Context, id string) (*models.Promotion, error) {
	if id == "" {
		return nil, ErrNoId
	}

	return s.promotionRepository.GetById(ctx, id)
}

func (s *promotionService) AddPromotion(ctx context.Context, requestBody models.PromotionRequest) (*mongo.InsertOneResult, error) {
	if requestBody.Code == "" || requestBody.Discount <= 0 || requestBody.AppliesTo == "" {
		return nil, ErrEmptyData
	}

	promotion := models.Promotion{
		Code:      requestBody.Code,
		Discount:  requestBody.Discount,
		ValidFrom: requestBody.ValidFrom,
		ValidTo:   requestBody.ValidTo,
		AppliesTo: requestBody.AppliesTo,
	}

	result, err := s.promotionRepository.Add(ctx, promotion)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *promotionService) UpdatePromotion(ctx context.Context, id string, requestBody models.Promotion) (*mongo.UpdateResult, error) {
	if id == "" {
		return nil, ErrNoId
	}

	promotion := models.Promotion{
		Code:      requestBody.Code,
		Discount:  requestBody.Discount,
		ValidFrom: requestBody.ValidFrom,
		ValidTo:   requestBody.ValidTo,
		AppliesTo: requestBody.AppliesTo,
	}

	result, err := s.promotionRepository.Update(ctx, id, promotion)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *promotionService) DeletePromotion(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	if id == "" {
		return nil, ErrNoId
	}

	return s.promotionRepository.Delete(ctx, id)
}
