package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)

type AnnouncementServicer interface {
	Service[models.Announcement, string, models.CreateAnnouncementParams, models.UpdateAnnouncementParams]
}

type AnnouncementService struct {
	q *models.Queries
}

// compile time check
var _ AnnouncementServicer = (*AnnouncementService)(nil)

func NewAnnouncementService(q *models.Queries) *AnnouncementService {
	return &AnnouncementService{q: q}
}

func (s *AnnouncementService) Get(ctx context.Context, uuid string) (models.Announcement, error) {
	announcement, err := s.q.GetAnnouncement(ctx, uuid)
	if err != nil {
		return models.Announcement{}, err
	}
	return announcement, nil
}

func (s *AnnouncementService) Create(ctx context.Context, params models.CreateAnnouncementParams) error {
	if err := s.q.CreateAnnouncement(ctx, params); err != nil {
		return err
	}
	return nil
}

type AnnouncementFilter interface {
	Apply(events []models.Announcement) []models.Announcement
}

func (s *AnnouncementService) List(ctx context.Context, filters ...any) ([]models.Announcement, error) {
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

func (s *AnnouncementService) Update(ctx context.Context, uuid string, params models.UpdateAnnouncementParams) error {
	err := s.q.UpdateAnnouncement(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *AnnouncementService) Delete(ctx context.Context, uuid string) error {
	err := s.q.DeleteAnnouncement(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}
