// Package events implements the bulk of the business logic for events
package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/repository"
)

type EventsServicer interface {
	Service[domain.Event, string, domain.Event, domain.Event]
}

type EventsService struct {
	eventRepo repository.EventRepository
}

// this checks that EventsService implements EventsServicers at compile time
var _ EventsServicer = (*EventsService)(nil)

func NewEventsService(eventRepo repository.EventRepository) *EventsService {
	return &EventsService{eventRepo: eventRepo}
}

func (s *EventsService) Get(ctx context.Context, uuid string) (domain.Event, error) {
	event, err := s.eventRepo.GetByID(ctx, uuid)
	if err != nil {
		return domain.Event{}, err
	}
	return event, nil
}

func (s *EventsService) Create(ctx context.Context, params domain.Event) error {
	if err := s.eventRepo.Create(ctx, params); err != nil {
		return err
	}
	return nil
}

// TODO: Move filters to their own file or module or something
type EventFilter interface {
	Apply(events []domain.Event) []domain.Event
}

type HostFilter struct {
	Host string
}

// Ensure HostFilter implements EventFilter
var _ EventFilter = (*HostFilter)(nil)

func (f *HostFilter) Apply(events []domain.Event) []domain.Event {
	if f.Host == "" {
		return events
	}

	filtered := make([]domain.Event, 0)
	for _, event := range events {
		if event.Host == f.Host {
			filtered = append(filtered, event)
		}
	}
	return filtered
}

func (s *EventsService) List(ctx context.Context, filters ...any) ([]domain.Event, error) {
	events, err := s.eventRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	result := events
	for _, filter := range filters {
		if eventFilter, ok := filter.(EventFilter); ok {
			result = eventFilter.Apply(result)
		}
	}

	return result, nil
}

func (s *EventsService) Update(ctx context.Context, uuid string, params domain.Event) error {
	params.Uuid = uuid
	if err := s.eventRepo.Update(ctx, params); err != nil {
		return err
	}
	return nil
}

func (s *EventsService) Delete(ctx context.Context, uuid string) error {
	if err := s.eventRepo.Delete(ctx, uuid); err != nil {
		return err
	}
	return nil
}
