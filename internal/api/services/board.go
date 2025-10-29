package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)

type BoardServicer interface {
	// Officer methods
	GetOfficer(ctx context.Context, id string) (models.Officer, error)
	ListOfficers(ctx context.Context, filters ...any) ([]models.Officer, error)
	CreateOfficer(ctx context.Context, params models.CreateOfficerParams) error
	UpdateOfficer(ctx context.Context, id string, params models.UpdateOfficerParams) error
	DeleteOfficer(ctx context.Context, id string) error

	// Tier methods
	GetTier(ctx context.Context, tierName int64) (models.Tier, error)
	ListTiers(ctx context.Context, filters ...any) ([]models.Tier, error)
	CreateTier(ctx context.Context, params models.CreateTierParams) error
	UpdateTier(ctx context.Context, params models.UpdateTierParams) error
	DeleteTier(ctx context.Context, tierName int64) error

	// Position methods
	GetPosition(ctx context.Context, oid string) (models.Position, error)
	ListPositions(ctx context.Context, filters ...any) ([]models.Position, error)
	CreatePosition(ctx context.Context, params models.CreatePositionParams) error
	UpdatePosition(ctx context.Context, params models.UpdatePositionParams) error
	DeletePosition(ctx context.Context, arg models.DeletePositionParams) error
}

type BoardService struct {
	q  *models.Queries
	db models.DBTX
}

var _ BoardServicer = (*BoardService)(nil)

func NewBoardService(q *models.Queries, db models.DBTX) *BoardService {
	return &BoardService{
		q:  q,
		db: db,
	}
}

// Officer Methods
func (s *BoardService) GetOfficer(ctx context.Context, uuid string) (models.Officer, error) {
	row, err := s.q.GetOfficer(ctx, uuid)
	if err != nil {
		return models.Officer{}, err
	}

	return models.Officer{
		Uuid:     uuid,
		FullName: row.FullName,
		Picture:  row.Picture,
		Github:   row.Github,
		Discord:  row.Discord,
	}, nil
}

func (s *BoardService) ListOfficers(ctx context.Context, filters ...any) ([]models.Officer, error) {
	query := `SELECT uuid, full_name, picture, github, discord FROM officer`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var officers []models.Officer
	for rows.Next() {
		var o models.Officer
		if err := rows.Scan(&o.Uuid, &o.FullName, &o.Picture, &o.Github, &o.Discord); err != nil {
			return nil, err
		}
		officers = append(officers, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return officers, nil
}

func (s *BoardService) CreateOfficer(ctx context.Context, params models.CreateOfficerParams) error {
	_, err := s.q.CreateOfficer(ctx, params)
	return err
}

func (s *BoardService) UpdateOfficer(ctx context.Context, uuid string, params models.UpdateOfficerParams) error {
	params.Uuid = uuid
	return s.q.UpdateOfficer(ctx, params)
}

func (s *BoardService) DeleteOfficer(ctx context.Context, uuid string) error {
	return s.q.DeleteOfficer(ctx, uuid)
}

// Tier Methods
func (s *BoardService) GetTier(ctx context.Context, tierName int64) (models.Tier, error) {
	return s.q.GetTier(ctx, tierName)
}

func (s *BoardService) ListTiers(ctx context.Context, filters ...any) ([]models.Tier, error) {
	query := `SELECT tier, title, t_index, team FROM tier ORDER BY tier`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tiers []models.Tier
	for rows.Next() {
		var t models.Tier
		if err := rows.Scan(&t.Tier, &t.Title, &t.TIndex, &t.Team); err != nil {
			return nil, err
		}
		tiers = append(tiers, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tiers, nil
}

func (s *BoardService) CreateTier(ctx context.Context, params models.CreateTierParams) error {
	_, err := s.q.CreateTier(ctx, params)
	return err
}

func (s *BoardService) UpdateTier(ctx context.Context, params models.UpdateTierParams) error {
	return s.q.UpdateTier(ctx, params)
}

func (s *BoardService) DeleteTier(ctx context.Context, tierName int64) error {
	return s.q.DeleteTier(ctx, tierName)
}

// Position Methods
func (s *BoardService) GetPosition(ctx context.Context, oid string) (models.Position, error) {
	return s.q.GetPosition(ctx, oid)
}

func (s *BoardService) ListPositions(ctx context.Context, filters ...any) ([]models.Position, error) {
	query := `SELECT oid, semester, tier, full_name, title, team FROM position`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var positions []models.Position
	for rows.Next() {
		var p models.Position
		if err := rows.Scan(&p.Oid, &p.Semester, &p.Tier, &p.FullName, &p.Title, &p.Team); err != nil {
			return nil, err
		}
		positions = append(positions, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}

func (s *BoardService) CreatePosition(ctx context.Context, params models.CreatePositionParams) error {
	_, err := s.q.CreatePosition(ctx, params)
	return err
}

func (s *BoardService) UpdatePosition(ctx context.Context, params models.UpdatePositionParams) error {
	return s.q.UpdatePosition(ctx, params)
}

func (s *BoardService) DeletePosition(ctx context.Context, arg models.DeletePositionParams) error {
	return s.q.DeletePosition(ctx, arg)
}
