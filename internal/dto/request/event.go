package domain

import "time"

type Event struct {
	Uuid     string    `json:"uuid"`
	Location string    `json:"location"`
	StartAt  time.Time `json:"start_at"`
	EndAt    time.Time `json:"end_at"`
	IsAllDay bool      `json:"is_all_day"`
	Host     string    `json:"host"`
}
