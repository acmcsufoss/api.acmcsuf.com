package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)

type BoardServicer interface {
	//Service[models.Officer, models.Tier, models.Position]
}

type BoardService struct {
	q *models.Queries
}

var _ BoardServicer = (*BoardService)(nil)

func NewBoardService(q *models.Queries) *BoardService {
	return &BoardService{q: q}
}

func (s *BoardService) GetOfficer(ctx context.Context, uuid interface{}) (models.Officer, error) {
	officer, err := s.q.GetOfficer(ctx, uuid)
	if err != nil {
		return models.Officer{}, err
	}
	return officer, nil
}

func (s *BoardService) CreateOfficer(ctx context.Context, params models.CreateOfficerParams) error {
	err := s.q.CreateOfficer(ctx, params)
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
