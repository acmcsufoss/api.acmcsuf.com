package services

import (
	"context"

	"log"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/repository"
)

type AnnouncementServicer interface {
	Service[dbmodels.Announcement, string, dbmodels.CreateAnnouncementParams,
		dbmodels.UpdateAnnouncementParams]
}

type AnnouncementService struct {
	q *dbmodels.Queries
}

// compile time check
var _ AnnouncementServicer = (*AnnouncementService)(nil)

func NewAnnouncementService(q *dbmodels.Queries) *repository.AnnouncementService {
	return &AnnouncementService{q: q}
}

func (s *AnnouncementService) Get(ctx context.Context, uuid string) (domain.Announcement, error) {
	announcement, err := s.q.GetAnnouncement(ctx, uuid)
	if err != nil {
		return dbmodels.Announcement{}, err
	}
	return announcement, nil
}

func (s *AnnouncementService) Create(ctx context.Context,
	params dbmodels.CreateAnnouncementParams) error {

	if err := s.q.CreateAnnouncement(ctx, params); err != nil {
		return err
	}
	return nil
}

type AnnouncementFilter interface {
	Apply(events []dbmodels.Announcement) []dbmodels.Announcement
}

func (s *AnnouncementService) List(ctx context.Context,
	filters ...any) ([]dbmodels.Announcement, error) {

	announcements, err := s.q.GetAnnouncements(ctx)
	if err != nil {
		return nil, err
	}

	result := announcements
	for _, filter := range filters {
		if announcementFilter, ok := filter.(AnnouncementFilter); ok {
			result = announcementFilter.Apply(result)
		}
	}
	return result, nil
}

func (s *AnnouncementService) Update(ctx context.Context, uuid string,
	params dbmodels.UpdateAnnouncementParams) error {

	err := s.q.UpdateAnnouncement(ctx, params)
	if err != nil {
		log.Printf("Error updating announcement with UUID %s: %v", uuid, err)
		return err
	}
	return nil
}

func (s *AnnouncementService) Delete(ctx context.Context, uuid string) error {
	err := s.q.DeleteAnnouncement(ctx, uuid)
	if err != nil {
		log.Printf("Error deleting announcement with UUID %s: %v", uuid, err)
		return err
	}
	return nil
}
