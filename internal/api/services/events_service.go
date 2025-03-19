// Package services implements the bulk of the business logic
package services

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)


type EventsService struct {
	q *models.Queries
}

var ddl string

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

// TODO: implement the following services
// NOTE: these are just copy-pasted from GetEvent and need to have their interfaces modified
func (s *EventsService) GetEvents(ctx context.Context, uuid string) (models.Event, error) {
	panic("implement me")
}

func (s *EventsService) CreateEvent(ctx context.Context, uuid string) (models.Event, error) {
	panic("implement me")
}

func (s *EventsService) UpdateEvent(ctx context.Context, uuid string) (models.Event, error) {
	panic("implement me")
}

func (s *EventsService) DeleteEvent(ctx context.Context, uuid string) (models.Event, error) {
	panic("implement me")
}

