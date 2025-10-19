package services

import (
	"context"
	"errors"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)

type BoardServicer interface {
	GetOfficer(ctx context.Context, id string) (models.Officer, error)
	ListOfficers(ctx context.Context, filters ...any) ([]models.Officer, error)
	CreateOfficer(ctx context.Context, params models.CreateOfficerParams) error
	UpdateOfficer(ctx context.Context, id string, params models.UpdateOfficerParams) error
	DeleteOfficer(ctx context.Context, id string) error

	GetTier(ctx context.Context, tierName int64) (models.Tier, error)
	ListTiers(ctx context.Context, filters ...any) ([]models.Tier, error)
	CreateTier(ctx context.Context, params models.CreateTierParams) error
	DeleteTier(ctx context.Context, tierName int64) error

	GetPosition(ctx context.Context, fullName string) (models.Position, error)
	CreatePosition(ctx context.Context, params models.CreatePositionParams) error
}

type BoardService struct {
	q *models.Queries
}

var _ BoardServicer = (*BoardService)(nil)

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

func (s *BoardService) ListOfficers(ctx context.Context, filters ...any) ([]models.Officer, error) {
	return nil, errors.New("not implemented")
}

func (s *BoardService) CreateOfficer(ctx context.Context, params models.CreateOfficerParams) error {
	err := s.q.CreateOfficer(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *BoardService) UpdateOfficer(ctx context.Context, uuid string, params models.UpdateOfficerParams) error {
	err := s.q.UpdateOfficer(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *BoardService) DeleteOfficer(ctx context.Context, uuid string) error {
	err := s.q.DeleteOfficer(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s *BoardService) GetTier(ctx context.Context, tierName int64) (models.Tier, error) {
	tier, err := s.q.GetTier(ctx, tierName)
	if err != nil {
		return models.Tier{}, err
	}
	return tier, nil
}

func (s *BoardService) CreateTier(ctx context.Context, params models.CreateTierParams) error {
	err := s.q.CreateTier(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *BoardService) ListTiers(ctx context.Context, filters ...any) ([]models.Tier, error) {
	return nil, errors.New("not implemented")
}

func (s *BoardService) DeleteTier(ctx context.Context, tierName int64) error {
	err := s.q.DeleteOfficer(ctx, tierName)
	if err != nil {
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
	err := s.q.CreatePosition(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *BoardService) UpdatePosition(ctx context.Context, arg models.UpdateOfficerParams) error {
	err := s.q.UpdateOfficer(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

func (s *BoardService) DeletePosition(ctx context.Context, arg models.DeletePositionParams) error {
	err := s.q.DeletePosition(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
