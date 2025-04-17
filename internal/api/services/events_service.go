// Package events implements the bulk of the business logic for events
package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)

type EventsServicer interface {
	Service[models.Event, string, models.CreateEventParams, models.UpdateEventParams]
}

type EventsService struct {
	q *models.Queries
}

// this checks that EventsService implements EventsServicers at compile time
var _ EventsServicer = (*EventsService)(nil)

func NewEventsService(q *models.Queries) *EventsService {
	return &EventsService{q: q}
}

func (s *EventsService) Get(ctx context.Context, uuid string) (models.Event, error) {
	event, err := s.q.GetEvent(ctx, uuid)
	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}

func (s *EventsService) Create(ctx context.Context, params models.CreateEventParams) error {
	if err := s.q.CreateEvent(ctx, params); err != nil {
		return err
	}
	return nil
}

// TODO: Move filters to their own file or module or something
type EventFilter interface {
	Apply(events []models.Event) []models.Event
}

type HostFilter struct {
	Host string
}

// Ensure HostFilter implements EventFilter
var _ EventFilter = (*HostFilter)(nil)

func (f *HostFilter) Apply(events []models.Event) []models.Event {
	if f.Host == "" {
		return events
	}

	filtered := make([]models.Event, 0)
	for _, event := range events {
		if event.Host == f.Host {
			filtered = append(filtered, event)
		}
	}
	return filtered
}

func (s *EventsService) List(ctx context.Context, filters ...any) ([]models.Event, error) {
	events, err := s.q.GetEvents(ctx)
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

func (s *EventsService) Update(ctx context.Context, uuid string, params models.UpdateEventParams) error {
	panic("implement me (EventsService Update)")
}

func (s *EventsService) Delete(ctx context.Context, uuid string) error {
	panic("implement me (EventsService Delete)")
}
