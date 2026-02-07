package repository

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

type PositionRepository interface {
	GetAllpositions(ctx context.Context) ([]*domain.Officer, error)
	GetpositionByID(ctx context.Context, id string) (*domain.Officer, error)
	Create(ctx context.Context, args domain.Position) error
	Update(ctx context.Context, args domain.Position) error
	Delete(ctx context.Context, args domain.Position) error
}

type positionRepository struct {
	db *dbmodels.Queries
}

func NewPositionRepository(db *dbmodels.Queries) OfficerRepository {
	return &positionRepository{db: db}
}

func (r *positionRepository) GetAllPositions(ctx context.Context) ([]*domain.Position, error) {
	dbPositions, err := r.db.GetPositions(ctx)
	if err != nil {
		return nil, err
	}

	var positions []*domain.Position
	for _, dbPosition := range dbPositions {
		positions = append(positions, convertDBPositionToDomain(&dbPosition))
	}
	return positions, nil
}

func (r *positionRepository) GetPositionByID(ctx context.Context, id string) (*domain.Position, error) {
	dbPosition, err := r.db.GetPosition(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertDBPositionToDomain(&dbPosition), nil
}

func (r *positionRepository) Delete(ctx context.Context, args domain.Position) error {
	err := r.db.DeletePosition(ctx, *convertDomainToDeleteDBPosition(&args))
	if err != nil {
		return err
	}
	return nil
}

func (r *positionRepository) Create(ctx context.Context, args domain.Position) error {
	_, err := r.db.CreatePosition(ctx, *convertDomainToCreateDBPosition(&args))
	if err != nil {
		return err
	}
	return nil
}

func (r *positionRepository) Update(ctx context.Context, args domain.Position) error {
	err := r.db.UpdatePosition(ctx, *convertDomainToUpdateDBPosition(&args))
	if err != nil {
		return err
	}
	return nil
}
