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


type EngineServiceInterface interface{
	GetEngineById(ctx context.Context, id string) (*models.Engine, error)
	CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (*models.Engine, error)
	UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (*models.Engine, error)
	DeleteEngine(ctx context.Context, id string) (*models.Engine, error)
}