// Package events implements the bulk of the business logic for events
package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/store"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/store/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

type EventsServicer interface {
	Service[domain.Event, string, domain.Event, domain.UpdateEvent]
}

type EventsService struct {
	q *dbmodels.Queries
}

// this checks that EventsService implements EventsServicers at compile time
var _ EventsServicer = (*EventsService)(nil)

func NewEventsService(q *dbmodels.Queries) *EventsService {
	return &EventsService{q: q}
}

func (s *EventsService) Get(ctx context.Context, uuid string) (domain.Event, error) {
	event, err := s.q.GetEvent(ctx, uuid)
	if err != nil {
		return domain.Event{}, err
	}
	return store.EventDBToDomain(event), nil
}

func (s *EventsService) Create(ctx context.Context, params domain.Event) error {
	dbParams := store.EventDomainToDB(params)
	if err := s.q.CreateEvent(ctx, dbParams); err != nil {
		return err
	}
	return nil
}

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
	events, err := s.q.GetEvents(ctx)
	if err != nil {
		return nil, err
	}
	domainEvents := make([]domain.Event, len(events))
	for i, event := range events {
		domainEvents[i] = store.EventDBToDomain(event)
	}

	result := domainEvents
	for _, filter := range filters {
		if eventFilter, ok := filter.(EventFilter); ok {
			result = eventFilter.Apply(result)
		}
	}

	return result, nil
}

func (s *EventsService) Update(ctx context.Context, uuid string, params domain.UpdateEvent) error {
	dbParams := store.UpdateEventDomainToDB(params)
	dbParams.Uuid = uuid
	if err := s.q.UpdateEvent(ctx, dbParams); err != nil {
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
