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

// TODO: implement the following services
// NOTE: these are just copy-pasted from GetEvent and need to have their interfaces modified
func (s *AnnouncementService) CreateAnnouncement(ctx context.Context, arg models.CreateAnnouncementParams) error {
	panic("implement me")
}
