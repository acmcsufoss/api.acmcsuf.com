package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)

type BoardService struct {
	q *models.Queries
}

func NewBoardService(q *models.Queries) *BoardService {
	return &BoardService{q: q}
}

func (s *BoardService) GetOfficer(ctx context.Context, uuid string) (models.Officer, error) {
	officer, err := s.q.GetOfficer(ctx, uuid)
	if err != nil {
		return models.Officer{}, err
	}
	return officer, nil
}

func (s *BoardService) CreateOfficer(ctx context.Context, params models.CreateOfficerParams) error {
	if err := s.q.CreateOfficer(ctx, params); err != nil {
		return err
	}
	return nil
}

func (s *BoardService) GetTier(ctx context.Context, tier string) (models.Tier, error) {
	tier, err := s.q.GetTier(ctx, tier)
	if err != nil {
		return models.Tier{}, err
	}
	return tier, nil
}

func (s *BoardService) CreateTier(ctx context.Context, params models.CreateTierParams) error {
	if err := s.q.CreateTier(ctx, params); err != nil {
		return err
	}
	return nil
}

func (s *BoardService) GetPosition(ctx context.Context, fullName string) (models.Position, error) {
	position, err := s.q.GetPosition(ctx, fullName)
	if err != nil {
		return models.Position{}, err
	}
	return position, nil
}

func (s *BoardService) CreatePosition(ctx context.Context, params models.CreatePositionParams) error {
	if err := s.q.CreatePosition(ctx, params); err != nil {
		return err
	}
	return nil
}
