package dto_request

import (
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

type Event struct {
	Uuid     string `json:"uuid"`
	Location string `json:"location"`
	StartAt  int64  `json:"start_at"`
	EndAt    int64  `json:"end_at"`
	IsAllDay bool   `json:"is_all_day"`
	Host     string `json:"host"`
}

func (e *Event) ToDomain() domain.Event {
	return domain.Event{
		Location: e.Location,
		StartAt:  time.Unix(e.StartAt, 0),
		EndAt:    time.Unix(e.EndAt, 0),
		IsAllDay: e.IsAllDay,
		Host:     e.Host,
	}
}
