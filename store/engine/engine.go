package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/pranayyb/DriveThrough/models"
)

type EngineStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EngineStore {
	return &EngineStore{
		db: db,
	}
}

func (e EngineStore) GetEngineById(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Println("failed to rollback transaction: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Println("could not commit properly: %v", cmErr)
			}
		}
	}()
	err = tx.QueryRowContext(ctx, "SELECT id, displacement, no_of_cylinders, car_range FROM engines WHERE id=$1", id).Scan(
		&engine.EngineID,
		&engine.Displacement,
		&engine.NoOfCylinders,
		&engine.CarRange,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Failed to get engine:", err)
			return engine, nil
		}
	}
	return engine, err
}

func (e EngineStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {
	var engine models.Engine
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Println("failed to rollback transaction: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Println("could not commit properly: %v", cmErr)
			}
		}
	}()

	engineID := uuid.New()

	_, err = tx.ExecContext(ctx, "INSERT INTO engines(id, displacement, no_of_cylinders, car_range) VALUES ($1, $2, $3, $4)", engineID, engineReq.Displacement, engineReq.NoOfCylinders, engineReq.CarRange)
	if err != nil {
		return models.Engine{}, err
	}
	engine = models.Engine{
		EngineID:      engineID,
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}
	return engine, nil
}

func (e EngineStore) UpdateEngine(ctx context.Context, id string, engine *models.EngineRequest) (models.Engine, error) {
	engineID, err := uuid.Parse(id)
	if err != nil {
		return models.Engine{}, fmt.Errorf("invalid engine id: %w", err)
	}
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Errorf("failed to rollback transaction: %w", rbErr)
			}
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Errorf("could not commit properly: %w", cmErr)
			}
		}
	}()

	results, err := tx.ExecContext(ctx, "UPDATE engine SET displacement=$2, no_of_cylinders=$3, car_range=$4 WHERE id=$1", engineID, engine.Displacement, engine.NoOfCylinders, engine.CarRange)

	if err != nil {
		return models.Engine{}, err
	}

	rowAffected, err := results.RowsAffected()
	if err != nil {
		return models.Engine{}, err
	}

	if rowAffected == 0 {
		return models.Engine{}, fmt.Errorf("no rows were updated")
	}
	engineUpdated := models.Engine{
		EngineID:      engineID,
		Displacement:  engine.Displacement,
		NoOfCylinders: engine.NoOfCylinders,
		CarRange:      engine.CarRange,
	}
	return engineUpdated, nil
}
func (e EngineStore) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Println("failed to rollback transaction: %v", rbErr)
			}
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Println("could not commit properly: %v", cmErr)
			}
		}
	}()
	err = tx.QueryRowContext(ctx, "SELECT id, displacement, no_of_cylinders, car_range FROM engines WHERE id=$1", id).Scan(
		&engine.EngineID,
		&engine.Displacement,
		&engine.NoOfCylinders,
		&engine.CarRange,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Failed to get engine:", err)
			return engine, nil
		}
	}

	result, err := tx.ExecContext(ctx, "DELETE FROM engines WHERE id=$1", id)
	if err != nil {
		return models.Engine{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Engine{}, err
	}
	if rowsAffected == 0 {
		return models.Engine{}, fmt.Errorf("no rows were deleted")
	}
	
	return engine, nil
}
