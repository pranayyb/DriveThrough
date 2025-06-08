package car

import (
	"context"
	"github.com/pranayyb/DriveThrough/models"
	"github.com/pranayyb/DriveThrough/store"
)

type CarService struct {
	store store.CarStoreInterface
}

func NewCarService(store store.CarStoreInterface) *CarService {
	return &CarService{
		store: store,
	}
}

func (s *CarService) GetCarById(id string, ctx context.Context) (*models.Car, error) {
	car, err := s.store.GetCarById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (s *CarService) GetCarByBrand(brand string, ctx context.Context, isEngine bool) ([]models.Car, error) {
	car, err := s.store.GetCarByBrand(ctx, brand, isEngine)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (s *CarService) CreateCar(carReq *models.CarRequest, ctx context.Context) (*models.Car, error) {
	if err := models.ValidateRequest(*carReq); err != nil {
		return nil, err
	}
	car, err := s.store.CreateCar(ctx, carReq)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (s *CarService) UpdateCar(id string, carReq *models.CarRequest, ctx context.Context) (*models.Car, error) {
	if err := models.ValidateRequest(*carReq); err != nil {
		return nil, err
	}
	car, err := s.store.UpdateCar(ctx, id, carReq)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (s *CarService) DeleteCar(id string, ctx context.Context) (*models.Car, error) {
	car, err := s.store.DeleteCar(ctx, id)
	if err != nil {
		return nil, err
	}
	return &car, nil
}
