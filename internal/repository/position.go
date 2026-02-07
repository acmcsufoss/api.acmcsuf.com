/*
package repository

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

type OfficerRepository interface {
	GetAllOfficers(ctx context.Context) ([]*domain.Officer, error)
	GetOfficerByID(ctx context.Context, id string) (*domain.Officer, error)
	Create(ctx context.Context, args domain.Officer) error
	Update(ctx context.Context, args domain.Officer) error
	Delete(ctx context.Context, id string) error
}

type officerRepository struct {
	db *dbmodels.Queries
}

func NewOfficerRepository(db *dbmodels.Queries) OfficerRepository {
	return &officerRepository{db: db}
}

func (r *officerRepository) GetAllOfficers(ctx context.Context) ([]*domain.Officer, error) {
	dbOfficers, err := r.db.GetOfficers(ctx)
	if err != nil {
		return nil, err
	}

	var officers []*domain.Officer
	for _, dbOfficer := range dbOfficers {
		officers = append(officers, convertDBOfficerToDomain(&dbOfficer))
	}
	return officers, nil
}

func (r *officerRepository) GetOfficerByID(ctx context.Context, id string) (*domain.Officer, error) {
	row, err := r.db.GetOfficer(ctx, id) // Get officers and get officers return completly different things?

	if err != nil {
		return nil, err
	}

	dbOfficer := dbmodels.Officer{
		Uuid:     id,
		FullName: row.FullName,
		Picture:  row.Picture,
		Github:   row.Github,
		Discord:  row.Discord,
	}

	return convertDBOfficerToDomain(&dbOfficer), nil
}

func (r *officerRepository) Delete(ctx context.Context, id string) error {
	err := r.db.DeleteOfficer(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *officerRepository) Create(ctx context.Context, args domain.Officer) error {
	_, err := r.db.CreateOfficer(ctx, *convertDomaintoCreateDBOfficer(&args))
	if err != nil {
		return err
	}
	return nil
}

func (r *officerRepository) Update(ctx context.Context, args domain.Officer) error {
	err := r.db.UpdateOfficer(ctx, *convertDomaintoUpdateDBOfficer(&args))
	if err != nil {
		return err
	}
	return nil
}
*/