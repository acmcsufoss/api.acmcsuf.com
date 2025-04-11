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
	if announcement, err := s.q.CreateAnnouncement(ctx, arg); err != nil {
		return announcement, err
	}
	return models.Announcement{}, nil
}
