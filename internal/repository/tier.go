package repository

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

type TierRepository interface {
	Repository[domain.Tier, int64]
}

type tierRepository struct {
	db *dbmodels.Queries
}

func NewTierRepository(db *dbmodels.Queries) TierRepository {
	return &tierRepository{db: db}
}

func (r *tierRepository) GetAll(ctx context.Context) ([]domain.Tier, error) {
	dbTiers, err := r.db.GetTiers(ctx)
	if err != nil {
		return nil, err
	}

	var tiers []domain.Tier
	for _, dbTier := range dbTiers {
		tiers = append(tiers, convertDBTierToDomain(dbTier))
	}
	return tiers, nil
}

func (r *tierRepository) GetByID(ctx context.Context, id int64) (domain.Tier, error) {
	dbTier, err := r.db.GetTier(ctx, id)
	if err != nil {
		return domain.Tier{}, err
	}

	return convertDBTierToDomain(dbTier), nil
}

func (r *tierRepository) Delete(ctx context.Context, id int64) error {
	err := r.db.DeleteTier(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *tierRepository) Create(ctx context.Context, args domain.Tier) error {
	_, err := r.db.CreateTier(ctx, convertDomainToCreateDBTier(args))
	if err != nil {
		return err
	}
	return nil
}

func (r *tierRepository) Update(ctx context.Context, args domain.Tier) error {
	err := r.db.UpdateTier(ctx, convertDomainToUpdateDBTier(args))
	if err != nil {
		return err
	}
	return nil
}
