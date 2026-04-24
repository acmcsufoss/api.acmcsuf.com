package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/store"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/store/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

type BoardServicer interface {
	// Officer methods
	GetOfficer(ctx context.Context, id string) (domain.Officer, error)
	ListOfficers(ctx context.Context, filters ...any) ([]domain.Officer, error)
	CreateOfficer(ctx context.Context, params domain.Officer) error
	UpdateOfficer(ctx context.Context, id string, params domain.UpdateOfficer) error
	DeleteOfficer(ctx context.Context, id string) error

	// Tier methods
	GetTier(ctx context.Context, tierName int64) (domain.Tier, error)
	ListTiers(ctx context.Context, filters ...any) ([]domain.Tier, error)
	CreateTier(ctx context.Context, params domain.Tier) (domain.Tier, error)
	UpdateTier(ctx context.Context, tierName int64, params domain.UpdateTier) error
	DeleteTier(ctx context.Context, tierName int64) error

	// Position methods
	GetPosition(ctx context.Context, oid string) (domain.Position, error)
	ListPositions(ctx context.Context, filters ...any) ([]domain.Position, error)
	CreatePosition(ctx context.Context, params domain.Position) (domain.Position, error)
	UpdatePosition(ctx context.Context, params domain.UpdatePosition) error
	DeletePosition(ctx context.Context, arg domain.DeletePosition) error
}

type BoardService struct {
	q *dbmodels.Queries
}

var _ BoardServicer = (*BoardService)(nil)

func NewBoardService(q *dbmodels.Queries, db dbmodels.DBTX) *BoardService {
	return &BoardService{
		q: q,
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

// ==== Officer Methods ========================================================

func (s *BoardService) GetOfficer(ctx context.Context, uuid string) (domain.Officer, error) {
	row, err := s.q.GetOfficer(ctx, uuid)
	if err != nil {
		return domain.Officer{}, err
	}

	domainModel := store.GetOfficerDBToDomain(row)
	domainModel.Uuid = uuid
	return domainModel, nil
}

func (s *BoardService) ListOfficers(ctx context.Context, filters ...any) ([]domain.Officer, error) {
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

	domainAs := make([]domain.Officer, len(result))
	for i, elm := range result {
		domainAs[i] = store.OfficerDBToDomain(elm)
	}

	return domainAs, nil
}

func (s *BoardService) CreateOfficer(ctx context.Context, officer domain.Officer) error {
	dbParams := store.OfficerDomainToDB(officer)
	_, err := s.q.CreateOfficer(ctx, dbParams)
	return err
}

func (s *BoardService) UpdateOfficer(ctx context.Context, uuid string,
	updateOfficer domain.UpdateOfficer) error {
	dbParams := store.UpdateOfficerDomainToDB(updateOfficer)
	dbParams.Uuid = uuid
	return s.q.UpdateOfficer(ctx, dbParams)
}

func (s *BoardService) DeleteOfficer(ctx context.Context, uuid string) error {
	return s.q.DeleteOfficer(ctx, uuid)
}

// ==== Tier Methods ===========================================================

func (s *BoardService) GetTier(ctx context.Context, tierName int64) (domain.Tier, error) {
	dbTier, err := s.q.GetTier(ctx, tierName)
	if err != nil {
		return domain.Tier{}, err
	}
	return store.TierDBToDomain(dbTier), nil
}

func (s *BoardService) ListTiers(ctx context.Context, filters ...any) ([]domain.Tier, error) {
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

	domainTiers := make([]domain.Tier, len(result))
	for i, tier := range result {
		domainTiers[i] = store.TierDBToDomain(tier)
	}
	return domainTiers, nil
}

func (s *BoardService) CreateTier(ctx context.Context, params domain.Tier) (domain.Tier, error) {
	dbTier, err := s.q.CreateTier(ctx, store.TierDomainToDB(params))
	if err != nil {
		return domain.Tier{}, nil
	}
	return store.TierDBToDomain(dbTier), nil
}

func (s *BoardService) UpdateTier(ctx context.Context, tierName int64,
	params domain.UpdateTier) error {
	dbParams := store.UpdateTierDomainToDB(params)
	dbParams.Tier = tierName
	return s.q.UpdateTier(ctx, dbParams)
}

func (s *BoardService) DeleteTier(ctx context.Context, tierName int64) error {
	return s.q.DeleteTier(ctx, tierName)
}

// ==== Position Methods =======================================================

func (s *BoardService) GetPosition(ctx context.Context, oid string) (domain.Position, error) {
	dbPosition, err := s.q.GetPosition(ctx, oid)
	if err != nil {
		return domain.Position{}, nil
	}
	return store.PositionDBToDomain(dbPosition), nil
}

func (s *BoardService) ListPositions(ctx context.Context, filters ...any) ([]domain.Position, error) {
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

	domainPositions := make([]domain.Position, len(result))
	for i, pos := range positions {
		domainPositions[i] = store.PositionDBToDomain(pos)
	}
	return domainPositions, nil
}

func (s *BoardService) CreatePosition(ctx context.Context,
	params domain.Position) (domain.Position, error) {
	dbPosition, err := s.q.CreatePosition(ctx, store.PositionDomainToDB(params))
	if err != nil {
		return domain.Position{}, err
	}
	return store.PositionDBToDomain(dbPosition), nil
}

func (s *BoardService) UpdatePosition(ctx context.Context, params domain.UpdatePosition) error {
	return s.q.UpdatePosition(ctx, store.UpdatePositionDomainToDB(params))
}

func (s *BoardService) DeletePosition(ctx context.Context, arg domain.DeletePosition) error {
	return s.q.DeletePosition(ctx, dbmodels.DeletePositionParams{
		OfficerID: arg.OfficerID,
		Semester:  arg.Semester,
		Tier:      arg.Tier,
	})
}
