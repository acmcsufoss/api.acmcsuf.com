package services

import (
	"context"
	"fmt"
	"log"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/repository"
)

type AnnouncementServicer interface {
	Service[domain.Announcement, string]
}

type AnnouncementService struct {
	announcementRepository repository.AnnouncementRepository
}

// compile time check
var _ AnnouncementServicer = (*AnnouncementService)(nil)

func NewAnnouncementService(announcementRepository repository.AnnouncementRepository) *AnnouncementService {
	return &AnnouncementService{announcementRepository: announcementRepository}
}

func (s *AnnouncementService) Get(ctx context.Context, uuid string) (domain.Announcement, error) {
	announcement, err := s.announcementRepository.GetByID(ctx, uuid)
	if err != nil {
		return domain.Announcement{}, err
	}
	return announcement, nil
}

func (s *AnnouncementService) Create(ctx context.Context,
	params domain.Announcement,
) error {
	if params.Uuid == "" {
		return fmt.Errorf("no unique identifier for announcement")
	}
	if err := s.announcementRepository.Create(ctx, params); err != nil {
		return err
	}
	return nil
}

type AnnouncementFilter interface {
	Apply(events []domain.Announcement) []domain.Announcement
}

func (s *AnnouncementService) List(ctx context.Context,
	filters ...any,
) ([]domain.Announcement, error) {
	announcements, err := s.announcementRepository.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
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
	params domain.Announcement,
) error {
	err := s.announcementRepository.Update(ctx, params)
	if err != nil {
		log.Printf("Error updating announcement with UUID %s: %v", uuid, err)
		return err
	}
	return nil
}

func (s *AnnouncementService) Delete(ctx context.Context, uuid string) error {
	err := s.announcementRepository.Delete(ctx, uuid)
	if err != nil {
		log.Printf("Error deleting announcement with UUID %s: %v", uuid, err)
		return err
	}
	return nil
}
