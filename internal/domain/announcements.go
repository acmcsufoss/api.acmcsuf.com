package domain

import "time"

type Announcements struct {
	Uuid             string
	Visibility       string
	AnnounceAt       time.Time
	DiscordChannelID string
	DiscordMessageID string
}
