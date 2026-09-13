package domain

import (
	"time"
)

type Announcement struct {
	Uuid             string
	Visibility       string
	AnnounceAt       time.Time
	DiscordChannelID *string
	DiscordMessageID *string
}

type UpdateAnnouncement struct {
	Uuid             string
	Visibility       *string
	AnnounceAt       *time.Time
	DiscordChannelID *string
	DiscordMessageID *string
}

func (a Announcement) IsZero() bool {
	return a.Uuid == ""
}
