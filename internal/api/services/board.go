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

func (s *BoardService) GetTier(ctx context.Context, tier string) (models.Officer, error) {
	officer, err := s.q.GetOfficer(ctx, uuid)
	if err != nil {
		return models.Officer{}, err
	}
	return officer, nil
}

func (s *BoardService) CreateTier(ctx context.Context, params models.CreateOfficerParams) error {
	if err := s.q.CreateOfficer(ctx, params); err != nil {
		return err
	}
	return nil
}

func (s *BoardService) GetPosition(ctx context.Context, uuid string) (models.Officer, error) {
	officer, err := s.q.GetOfficer(ctx, uuid)
	if err != nil {
		return models.Officer{}, err
	}
	return officer, nil
}

func (s *BoardService) CreatePosition(ctx context.Context, params models.CreateOfficerParams) error {
	if err := s.q.CreateOfficer(ctx, params); err != nil {
		return err
	}
	return nil
}
