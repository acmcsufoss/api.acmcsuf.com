package services

import (
	"context"
	"fmt"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/repository"
)

type PositionServicer interface {
	Service[domain.Position, string, domain.UpdatePosition]

	DeletePosition(ctx context.Context, arg domain.Position) error
}

type PositionService struct {
	positionRepo repository.PositionRepository
}

var _ PositionServicer = (*PositionService)(nil)

// There used to be a dbmodels.DBTX var here but I don't think it was used?
func NewPositionService(positionRepo repository.PositionRepository) *PositionService {
	return &PositionService{positionRepo: positionRepo}
}

type PositionFilter interface {
	Apply(positions []domain.Position) []domain.Position
}

// Position Methods
func (s *PositionService) Get(ctx context.Context, oid string) (domain.Position, error) {
	position, err := s.positionRepo.GetByID(ctx, oid)
	if err != nil {
		return domain.Position{}, err
	}
	return position, nil
}

func (s *PositionService) List(ctx context.Context, filters ...any) ([]domain.Position, error) {
	positions, err := s.positionRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	result := positions
	for _, filter := range filters {
		if positionFilter, ok := filter.(PositionFilter); ok {
			result = positionFilter.Apply(result)
		}
	}

	return result, nil
}

func (s *PositionService) Create(ctx context.Context, params domain.Position) error {
	err := s.positionRepo.Create(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *PositionService) Update(ctx context.Context, uuid string, params domain.UpdatePosition) error {
	err := s.positionRepo.Update(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (s *PositionService) DeletePosition(ctx context.Context, arg domain.Position) error {
	err := s.positionRepo.DeletePosition(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

func (s *PositionService) Delete(ctx context.Context, arg string) error {
	// This function is here only to satisfy Service interface, use DeletePosition to delete a position
	return fmt.Errorf("error: the developer is using the wrong delete function")
}
