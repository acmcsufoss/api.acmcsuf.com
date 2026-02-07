package repository

import (
	"context"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

type EventRepository interface {
	GetAllEvents(ctx context.Context) ([]*domain.Event, error)

	GetEventByID(ctx context.Context, id string) (*domain.Event, error)
	Delete(ctx context.Context, id string) error

	Create(ctx context.Context, args domain.Event) error
	Update(ctx context.Context, args domain.Event) error
}

type eventRepository struct {
	db *dbmodels.Queries
}

func NewEventRepository(db *dbmodels.Queries) EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) GetEventByID(ctx context.Context, id string) (*domain.Event, error) {
	dbEvent, err := r.db.GetEvent(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertDBEventToDomain(&dbEvent), nil
}

func (r *eventRepository) GetAllEvents(ctx context.Context) ([]*domain.Event, error) {
	dbEvent, err := r.db.GetEvents(ctx)
	if err != nil {
		return nil, err
	}

	var eventSlice []*domain.Event
	for _, elm := range dbEvent {
		eventSlice = append(eventSlice, convertDBEventToDomain(&elm))
	}
	return eventSlice, nil
}

func (r *eventRepository) Delete(ctx context.Context, id string) error {
	err := r.db.DeleteEvent(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *eventRepository) Create(ctx context.Context, args domain.Event) error {
	err := r.db.CreateEvent(ctx, *convertDomainToCreateDBEvent(&args))
	if err != nil {
		return err
	}
	return nil
}

func (r *eventRepository) Update(ctx context.Context, args domain.Event) error {
	err := r.db.UpdateEvent(ctx, *convertDomainToUpdateDBEvent(&args))
	if err != nil {
		return err
	}
	return nil
}
