package dto

import (
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

type Event struct {
	Uuid     string `json:"uuid"`
	Location string `json:"location"`
	StartAt  int64  `json:"start_at"`
	EndAt    int64  `json:"end_at"`
	IsAllDay bool   `json:"is_all_day"`
	Host     string `json:"host"`
}

func (e Event) ToDomain() domain.Event {
	return domain.Event{
		Uuid:     e.Uuid,
		Location: e.Location,
		StartAt:  utils.UnixToTime(e.StartAt),
		EndAt:    utils.UnixToTime(e.EndAt),
		IsAllDay: e.IsAllDay,
		Host:     e.Host,
	}
}

func EventDomainToDto(e *domain.Event) Event {
	return Event{
		Uuid:     e.Uuid,
		Location: e.Location,
		StartAt:  e.StartAt.Unix(),
		EndAt:    e.EndAt.Unix(),
		IsAllDay: e.IsAllDay,
		Host:     e.Host,
	}
}

type UpdateEvent struct {
	Location *string `json:"location"`
	StartAt  *int64  `json:"start_at"`
	EndAt    *int64  `json:"end_at"`
	IsAllDay *bool   `json:"is_all_day"`
	Host     *string `json:"host"`
}

func (e UpdateEvent) ToDomain() domain.UpdateEvent {
	d := domain.UpdateEvent{
		Location: e.Location,
		IsAllDay: e.IsAllDay,
		Host:     e.Host,
	}
	if e.StartAt != nil {
		*d.StartAt = utils.UnixToTime(*e.StartAt)
	}
	if e.EndAt != nil {
		*d.EndAt = utils.UnixToTime(*e.EndAt)
	}
	return d
}
