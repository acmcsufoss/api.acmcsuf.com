// Package events implements the bulk of the business logic for events
package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)

type EventsServicer interface {
	GetEvent(ctx context.Context, uuid string) (models.Event, error)
	CreateEvent(ctx context.Context, params models.CreateEventParams) error
	GetEvents(ctx context.Context) ([]models.Event, error)
	UpdateEvent(ctx context.Context, uuid string) error
	DeleteEvent(ctx context.Context, uuid string) error
}

type EventsService struct {
	q *models.Queries
}

// this checks that EventsService implements EventsServicers at compile time
var _ EventsServicer = (*EventsService)(nil)

func NewEventsService(q *models.Queries) *EventsService {
	return &EventsService{q: q}
}

func (s *EventsService) GetEvent(ctx context.Context, uuid string) (models.Event, error) {
	event, err := s.q.GetEvent(ctx, uuid)
	if err != nil {
		return models.Event{}, err
	}
	return event, nil
}

func (s *EventsService) CreateEvent(ctx context.Context, params models.CreateEventParams) error {
	if err := s.q.CreateEvent(ctx, params); err != nil {
		return err
	}
	return nil
}

func (s *EventsService) GetEvents(ctx context.Context) ([]models.Event, error) {
	panic("implement me")
}

func (s *EventsService) UpdateEvent(ctx context.Context, uuid string) error {
	panic("implement me")
}

func (s *EventsService) DeleteEvent(ctx context.Context, uuid string) error {
	panic("implement me")
}
