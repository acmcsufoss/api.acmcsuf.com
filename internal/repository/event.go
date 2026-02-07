package repository

import (
	"context"
	"database/sql"
	"time"

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
	err := r.db.CreateEvent(ctx, *convertDomaintoCreateDBEvent(&args))
	if err != nil {
		return err
	}
	return nil
}

func (r *eventRepository) Update(ctx context.Context, args domain.Event) error {
	err := r.db.UpdateEvent(ctx, *convertDomaintoUpdateDBEvent(&args))
	if err != nil {
		return err
	}
	return nil
}

// ---- Helper func ----
func convertDBEventToDomain(dbEvent *dbmodels.Event) *domain.Event {
	return &domain.Event{
		Uuid:     dbEvent.Uuid,
		Location: dbEvent.Location,
		StartAt:  time.Unix(dbEvent.StartAt, 0),
		EndAt:    time.Unix(dbEvent.EndAt, 0),
		IsAllDay: dbEvent.IsAllDay,
		Host:     dbEvent.Host,
	}
}

func convertDomaintoCreateDBEvent(dEvent *domain.Event) *dbmodels.CreateEventParams {
	return &dbmodels.CreateEventParams{
		Uuid:     dEvent.Uuid,
		Location: dEvent.Location,
		StartAt:  dEvent.StartAt.Unix(),
		EndAt:    dEvent.EndAt.Unix(),
		IsAllDay: dEvent.IsAllDay,
		Host:     dEvent.Host,
	}
}

func convertDomaintoUpdateDBEvent(dEvent *domain.Event) *dbmodels.UpdateEventParams {
	return &dbmodels.UpdateEventParams{
		Uuid:     dEvent.Uuid,
		Location: sql.NullString{String: dEvent.Location, Valid: true},
		StartAt:  sql.NullInt64{Int64: dEvent.StartAt.Unix(), Valid: true},
		EndAt:    sql.NullInt64{Int64: dEvent.EndAt.Unix(), Valid: true},
		IsAllDay: sql.NullBool{Bool: dEvent.IsAllDay, Valid: true},
		Host:     sql.NullString{String: dEvent.Host, Valid: true},
	}
}
