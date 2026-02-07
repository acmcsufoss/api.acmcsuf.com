package domain

import "time"

type Event struct {
	Uuid     string
	Location string
	StartAt  time.Time
	EndAt    time.Time
	IsAllDay bool
	Host     string
}
