// Package events implements the bulk of the business logic for events
package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
)

type EventsServicer interface {
	Service[dbmodels.Event, string, dbmodels.CreateEventParams, dbmodels.UpdateEventParams]
}

type EventsService struct {
	q *dbmodels.Queries
}

// this checks that EventsService implements EventsServicers at compile time
var _ EventsServicer = (*EventsService)(nil)

func NewEventsService(q *dbmodels.Queries) *EventsService {
	return &EventsService{q: q}
}

func (s *EventsService) Get(ctx context.Context, uuid string) (dbmodels.Event, error) {
	event, err := s.q.GetEvent(ctx, uuid)
	if err != nil {
		return dbmodels.Event{}, err
	}
	return event, nil
}

func (s *EventsService) Create(ctx context.Context, params dbmodels.CreateEventParams) error {
	if err := s.q.CreateEvent(ctx, params); err != nil {
		return err
	}
	return nil
}

// TODO: Move filters to their own file or module or something
type EventFilter interface {
	Apply(events []dbmodels.Event) []dbmodels.Event
}

type HostFilter struct {
	Host string
}

// Ensure HostFilter implements EventFilter
var _ EventFilter = (*HostFilter)(nil)

func (f *HostFilter) Apply(events []dbmodels.Event) []dbmodels.Event {
	if f.Host == "" {
		return events
	}

	filtered := make([]dbmodels.Event, 0)
	for _, event := range events {
		if event.Host == f.Host {
			filtered = append(filtered, event)
		}
	}
	return filtered
}

func (s *EventsService) List(ctx context.Context, filters ...any) ([]dbmodels.Event, error) {
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

func (s *EventsService) Update(ctx context.Context, uuid string, params dbmodels.UpdateEventParams) error {
	params.Uuid = uuid
	if err := s.q.UpdateEvent(ctx, params); err != nil {
		return err
	}
	return nil
}

func (s *EventsService) Delete(ctx context.Context, uuid string) error {
	if err := s.q.DeleteEvent(ctx, uuid); err != nil {
		return err
	}
	return nil
}
