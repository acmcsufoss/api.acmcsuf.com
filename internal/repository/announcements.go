package repository

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

type AnnouncementRepository interface {
	GetAllAnnouncements(ctx context.Context) ([]*domain.Announcement, error)

	GetAnnouncementByID(ctx context.Context, id string) (*domain.Announcement, error)
	Delete(ctx context.Context, id string) error

	Create(ctx context.Context, args domain.Announcement) error
	Update(ctx context.Context, args domain.Announcement) error
}

type announcementRepository struct {
	db *dbmodels.Queries
}

func NewAnnouncementRepository(db *dbmodels.Queries) AnnouncementRepository {
	return &announcementRepository{db: db}
}

func (r *announcementRepository) GetAnnouncementByID(ctx context.Context, id string) (*domain.Announcement, error) {
	dbAnnouncement, err := r.db.GetAnnouncement(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertDBAnnouncementToDomain(&dbAnnouncement), nil
}

func (r *announcementRepository) GetAllAnnouncements(ctx context.Context) ([]*domain.Announcement, error) {
	dbAnnouncement, err := r.db.GetAnnouncements(ctx)
	if err != nil {
		return nil, err
	}

	var eventSlice []*domain.Announcement
	for _, elm := range dbAnnouncement {
		eventSlice = append(eventSlice, convertDBAnnouncementToDomain(&elm))
	}
	return eventSlice, nil
}

func (r *announcementRepository) Delete(ctx context.Context, id string) error {
	err := r.db.DeleteAnnouncement(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *announcementRepository) Create(ctx context.Context, args domain.Announcement) error {
	err := r.db.CreateAnnouncement(ctx, *convertDomaintoCreateDBAnnouncement(&args))
	if err != nil {
		return err
	}
	return nil
}

func (r *announcementRepository) Update(ctx context.Context, args domain.Announcement) error {
	err := r.db.UpdateAnnouncement(ctx, *convertDomaintoUpdateDBAnnouncement(&args))
	if err != nil {
		return err
	}
	return nil
}
