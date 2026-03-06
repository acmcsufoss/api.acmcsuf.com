package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/repository"
)

type OfficerServicer interface {
	Service[domain.Officer, string, domain.UpdateOfficer]
}

type OfficerService struct {
	officerRepo repository.OfficerRepository
}

var _ OfficerServicer = (*OfficerService)(nil)

// There used to be a dbmodels.DBTX var here but I don't think it was used?
func NewOfficerService(officerRepo repository.OfficerRepository) *OfficerService {
	return &OfficerService{officerRepo: officerRepo}
}

type OfficerFilter interface {
	Apply(officers []domain.Officer) []domain.Officer
}

// Officer Methods
func (s *OfficerService) Get(ctx context.Context, uuid string) (domain.Officer, error) {
	row, err := s.officerRepo.GetByID(ctx, uuid)
	if err != nil {
		return domain.Officer{}, err
	}

	return domain.Officer{
		Uuid:     uuid,
		FullName: row.FullName,
		Picture:  row.Picture,
		Github:   row.Github,
		Discord:  row.Discord,
	}, nil
}

func (s *OfficerService) List(ctx context.Context, filters ...any) ([]domain.Officer, error) {
	officers, err := s.officerRepo.GetAll(ctx)
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

func (s *OfficerService) Create(ctx context.Context, params domain.Officer) error {
	err := s.officerRepo.Create(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *OfficerService) Update(ctx context.Context, uuid string, params domain.UpdateOfficer) error {
	params.Uuid = uuid
	err := s.officerRepo.Update(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *OfficerService) Delete(ctx context.Context, uuid string) error {
	err := s.officerRepo.Delete(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}
