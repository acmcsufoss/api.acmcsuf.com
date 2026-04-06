package services

import (
	"context"
	"log"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/mapper"
)

type AnnouncementServicer interface {
	Service[domain.Announcement, string, domain.Announcement,
		domain.UpdateAnnouncement]
}

type AnnouncementService struct {
	q *dbmodels.Queries
}

// compile time check
var _ AnnouncementServicer = (*AnnouncementService)(nil)

func NewAnnouncementService(q *dbmodels.Queries) *AnnouncementService {
	return &AnnouncementService{q: q}
}

func (s *AnnouncementService) Get(ctx context.Context, uuid string) (domain.Announcement, error) {
	announcement, err := s.q.GetAnnouncement(ctx, uuid)
	if err != nil {
		return domain.Announcement{}, err
	}

	domainA := mapper.ToDBAnnouncementToDomain(announcement)

	return domainA, nil
}

func (s *AnnouncementService) Create(ctx context.Context,
	params domain.Announcement,
) error {
	dbA := mapper.ToDomainToCreateDBAnnouncement(params)
	if err := s.q.CreateAnnouncement(ctx, dbA); err != nil {
		return err
	}
	return nil
}

type AnnouncementFilter interface {
	Apply(events []dbmodels.Announcement) []dbmodels.Announcement
}

func (s *AnnouncementService) List(ctx context.Context,
	filters ...any,
) ([]domain.Announcement, error) {
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

	domainAs := make([]domain.Announcement, len(result))
	for i, elm := range result {
		domainAs[i] = mapper.ToDBAnnouncementToDomain(elm)
	}
	return domainAs, nil
}

func (s *AnnouncementService) Update(ctx context.Context, uuid string,
	params domain.UpdateAnnouncement,
) error {
	dbA := mapper.ToDomainToUpdateDBAnnouncement(params)

	err := s.q.UpdateAnnouncement(ctx, dbA)
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
