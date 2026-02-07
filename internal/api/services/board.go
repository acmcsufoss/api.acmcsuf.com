package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
)

type BoardServicer interface {
	// Officer methods
	GetOfficer(ctx context.Context, id string) (dbmodels.Officer, error)
	ListOfficers(ctx context.Context, filters ...any) ([]dbmodels.Officer, error)
	CreateOfficer(ctx context.Context, params dbmodels.CreateOfficerParams) error
	UpdateOfficer(ctx context.Context, id string, params dbmodels.UpdateOfficerParams) error
	DeleteOfficer(ctx context.Context, id string) error

	// Tier methods
	GetTier(ctx context.Context, tierName int64) (dbmodels.Tier, error)
	ListTiers(ctx context.Context, filters ...any) ([]dbmodels.Tier, error)
	CreateTier(ctx context.Context, params dbmodels.CreateTierParams) error
	UpdateTier(ctx context.Context, params dbmodels.UpdateTierParams) error
	DeleteTier(ctx context.Context, tierName int64) error

	// Position methods
	GetPosition(ctx context.Context, oid string) (dbmodels.Position, error)
	ListPositions(ctx context.Context, filters ...any) ([]dbmodels.Position, error)
	CreatePosition(ctx context.Context, params dbmodels.CreatePositionParams) error
	UpdatePosition(ctx context.Context, params dbmodels.UpdatePositionParams) error
	DeletePosition(ctx context.Context, arg dbmodels.DeletePositionParams) error
}

type BoardService struct {
	q  *dbmodels.Queries
	db dbmodels.DBTX
}

var _ BoardServicer = (*BoardService)(nil)

func NewBoardService(q *dbmodels.Queries, db dbmodels.DBTX) *BoardService {
	return &BoardService{
		q:  q,
		db: db,
	}
}

type OfficerFilter interface {
	Apply(officers []dbmodels.Officer) []dbmodels.Officer
}

type TierFilter interface {
	Apply(tiers []dbmodels.Tier) []dbmodels.Tier
}

type PositionFilter interface {
	Apply(positions []dbmodels.Position) []dbmodels.Position
}

// Officer Methods
func (s *BoardService) GetOfficer(ctx context.Context, uuid string) (dbmodels.Officer, error) {
	row, err := s.q.GetOfficer(ctx, uuid)
	if err != nil {
		return dbmodels.Officer{}, err
	}

	return dbmodels.Officer{
		Uuid:     uuid,
		FullName: row.FullName,
		Picture:  row.Picture,
		Github:   row.Github,
		Discord:  row.Discord,
	}, nil
}

func (s *BoardService) ListOfficers(ctx context.Context, filters ...any) ([]dbmodels.Officer, error) {
	officers, err := s.q.GetOfficers(ctx)
	if err != nil {
		return nil, err
	}

	result := officers
	for _, filter := range filters {
		if officerFilter, ok := filter.(OfficerFilter); ok {
			result = officerFilter.Apply(result)
		}
	}

	return result, nil
}

func (s *BoardService) CreateOfficer(ctx context.Context, params dbmodels.CreateOfficerParams) error {
	_, err := s.q.CreateOfficer(ctx, params)
	return err
}

func (s *BoardService) UpdateOfficer(ctx context.Context, uuid string, params dbmodels.UpdateOfficerParams) error {
	params.Uuid = uuid
	return s.q.UpdateOfficer(ctx, params)
}

func (s *BoardService) DeleteOfficer(ctx context.Context, uuid string) error {
	return s.q.DeleteOfficer(ctx, uuid)
}

// Tier Methods
func (s *BoardService) GetTier(ctx context.Context, tierName int64) (dbmodels.Tier, error) {
	return s.q.GetTier(ctx, tierName)
}

func (s *BoardService) ListTiers(ctx context.Context, filters ...any) ([]dbmodels.Tier, error) {
	tiers, err := s.q.GetTiers(ctx)
	if err != nil {
		return nil, err
	}

	result := tiers
	for _, filter := range filters {
		if tierFilter, ok := filter.(TierFilter); ok {
			result = tierFilter.Apply(result)
		}
	}

	return result, nil
}

func (s *BoardService) CreateTier(ctx context.Context, params dbmodels.CreateTierParams) error {
	_, err := s.q.CreateTier(ctx, params)
	return err
}

func (s *BoardService) UpdateTier(ctx context.Context, params dbmodels.UpdateTierParams) error {
	return s.q.UpdateTier(ctx, params)
}

func (s *BoardService) DeleteTier(ctx context.Context, tierName int64) error {
	return s.q.DeleteTier(ctx, tierName)
}

// Position Methods
func (s *BoardService) GetPosition(ctx context.Context, oid string) (dbmodels.Position, error) {
	return s.q.GetPosition(ctx, oid)
}

func (s *BoardService) ListPositions(ctx context.Context, filters ...any) ([]dbmodels.Position, error) {
	positions, err := s.q.GetPositions(ctx)
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

func (s *BoardService) CreatePosition(ctx context.Context, params dbmodels.CreatePositionParams) error {
	_, err := s.q.CreatePosition(ctx, params)
	return err
}

func (s *BoardService) UpdatePosition(ctx context.Context, params dbmodels.UpdatePositionParams) error {
	return s.q.UpdatePosition(ctx, params)
}

func (s *BoardService) DeletePosition(ctx context.Context, arg dbmodels.DeletePositionParams) error {
	return s.q.DeletePosition(ctx, arg)
}
