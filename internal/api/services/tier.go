package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/repository"
)

type TierServicer interface {
	Service[domain.Tier, int64, domain.UpdateTier]
}

type TierService struct {
	tierRepo repository.TierRepository
}

var _ TierServicer = (*TierService)(nil)

// There used to be a dbmodels.DBTX var here but I don't think it was used?
func NewTierService(tierRepo repository.TierRepository) *TierService {
	return &TierService{tierRepo: tierRepo}
}

type TierFilter interface {
	Apply(tiers []domain.Tier) []domain.Tier
}

// Tier Methods
func (s *TierService) Get(ctx context.Context, tierName int64) (domain.Tier, error) {
	return s.tierRepo.GetByID(ctx, tierName)
}

func (s *TierService) List(ctx context.Context, filters ...any) ([]domain.Tier, error) {
	tiers, err := s.tierRepo.GetAll(ctx)
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

func (s *TierService) Create(ctx context.Context, params domain.Tier) error {
	err := s.tierRepo.Create(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *TierService) Update(ctx context.Context, tierName int64, params domain.UpdateTier) error {
	params.Tier = int(tierName)
	err := s.tierRepo.Update(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *TierService) Delete(ctx context.Context, tierName int64) error {
	err := s.tierRepo.Delete(ctx, tierName)
	if err != nil {
		return err
	}
	return nil
}
