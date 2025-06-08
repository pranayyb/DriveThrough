package service

import (
	"context"

	"github.com/pranayyb/DriveThrough/models"
)

type CarServiceInterface interface {
	GetCarById(id string, ctx context.Context) (*models.Car, error)
	GetCarByBrand(brand string, ctx context.Context, isEngine bool) (*[]models.Car, error)
	CreateCar(carReq *models.CarRequest, ctx context.Context) (*models.Car, error)
	UpdateCar(id string, carReq *models.CarRequest, ctx context.Context) (*models.Car, error)
	DeleteCar(id string, ctx context.Context) (*models.Car, error)
}
