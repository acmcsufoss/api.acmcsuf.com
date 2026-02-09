package domain

import (
	"time"

	dto_response "github.com/acmcsufoss/api.acmcsuf.com/internal/dto/response"
)

type Event struct {
	Uuid     string
	Location string
	StartAt  time.Time
	EndAt    time.Time
	IsAllDay bool
	Host     string
}

func (e *Event) ToDTO() dto_response.Event {
	return dto_response.Event{
		Uuid:     e.Uuid,
		Location: e.Location,
		StartAt:  e.StartAt.Unix(),
		EndAt:    e.EndAt.Unix(),
		IsAllDay: e.IsAllDay,
		Host:     e.Host,
	}
}
