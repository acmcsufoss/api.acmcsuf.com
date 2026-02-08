package domain

import "time"

type Announcement struct {
	Visibility       string    `json:"visibility"`
	AnnounceAt       time.Time `json:"announce_at"`
	DiscordChannelID string    `json:"discord_channel_id"`
	DiscordMessageID string    `json:"discord_message_id"`
}
