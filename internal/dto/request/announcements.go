package dto_request

import (
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

type Announcement struct {
	Uuid             string `json:"uuid,omitempty"`
	Visibility       string `json:"visibility"`
	AnnounceAt       int64  `json:"announce_at"`
	DiscordChannelID string `json:"discord_channel_id"`
	DiscordMessageID string `json:"discord_message_id"`
}

func (a *Announcement) ToDomain() domain.Announcement {
	return domain.Announcement{
		Uuid:             a.Uuid,
		Visibility:       a.Visibility,
		AnnounceAt:       time.Unix(a.AnnounceAt, 0),
		DiscordChannelID: a.DiscordChannelID,
		DiscordMessageID: a.DiscordMessageID,
	}
}
