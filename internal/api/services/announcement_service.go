package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)

type AnnouncementService struct {
	q *models.Queries
}

var dll string

func NewAnnouncementService(q *models.Queries) *AnnouncementService {
	return &AnnouncementService{q: q}
}

func (s *AnnouncementService) GetAnnouncement(ctx context.Context, uuid string) (models.Announcement, error) {
	announcement, err := s.q.GetAnnouncement(ctx, uuid)
	if err != nil {
		return models.Announcement{}, err
	}
	return announcement, nil
}

func (s *AnnouncementService) CreateAnnouncement(ctx context.Context, arg models.CreateAnnouncementParams) (models.Announcement, error) {
	if err := s.q.CreateAnnouncement(ctx, params); err != nil {
		return err
	}
	return nil
}
