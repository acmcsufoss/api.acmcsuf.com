package domain

import "time"

type Announcement struct {
	Uuid             string
	Visibility       string
	AnnounceAt       time.Time
	DiscordChannelID string
	DiscordMessageID string
}
