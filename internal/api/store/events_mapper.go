package store

import (
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/store/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

func EventDomainToDB(event domain.Event) dbmodels.CreateEventParams {
	return dbmodels.CreateEventParams{
		Uuid:     event.Uuid,
		Location: event.Location,
		StartAt:  event.StartAt.Unix(),
		EndAt:    event.EndAt.Unix(),
		IsAllDay: event.IsAllDay,
		Host:     event.Host,
	}
}

func UpdateEventDomainToDB(event domain.UpdateEvent) dbmodels.UpdateEventParams {
	return dbmodels.UpdateEventParams{
		Uuid:     event.Uuid,
		Location: stringToNullString(event.Location),
		StartAt:  timeToNullInt64(event.StartAt),
		EndAt:    timeToNullInt64(event.EndAt),
		IsAllDay: boolToNullBool(event.IsAllDay),
		Host:     stringToNullString(event.Host),
	}
}

func EventDBToDomain(event dbmodels.Event) domain.Event {
	return domain.Event{
		Uuid:     event.Uuid,
		Location: event.Location,
		StartAt:  time.Unix(event.StartAt, 0),
		EndAt:    time.Unix(event.EndAt, 0),
		IsAllDay: event.IsAllDay,
		Host:     event.Host,
	}
}
