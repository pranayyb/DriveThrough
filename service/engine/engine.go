package engine

import (
	"context"
	"github.com/pranayyb/DriveThrough/models"
	"github.com/pranayyb/DriveThrough/store"
)

type EngineService struct {
	store store.EngineStoreInterface
}

func NewEngineService(store store.EngineStoreInterface) *EngineService {
	return &EngineService{
		store: store,
	}
}
func (s *EngineService) GetEngineById(ctx context.Context, id string) (*models.Engine, error) {
	engine, err := s.store.GetEngineById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &engine, nil
}

func (s *EngineService) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (*models.Engine, error) {
	if err := models.ValidateEngineRequest(*engineReq); err != nil {
		return nil, err
	}
	createdEngine, err := s.store.CreateEngine(ctx, engineReq)
	if err != nil {
		return nil, err
	}
	return &createdEngine, err
}

func (s *EngineService) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (*models.Engine, error) {
	if err := models.ValidateEngineRequest(*engineReq); err != nil {
		return nil, err
	}
	updatedEngine, err := s.store.UpdateEngine(ctx, id, engineReq)
	if err != nil {
		return nil, err
	}
	return &updatedEngine, nil
}

func (s *EngineService) DeleteEngine(ctx context.Context, id string) (*models.Engine, error) {
	deletedEngine, err := s.store.DeleteEngine(ctx, id)
	if err != nil {
		return nil, err
	}
	return &deletedEngine, nil
}
