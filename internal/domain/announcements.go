package domain

import (
	"time"

	dto_response "github.com/acmcsufoss/api.acmcsuf.com/internal/dto/response"
)

type Announcement struct {
	Uuid             string
	Visibility       string
	AnnounceAt       time.Time
	DiscordChannelID string
	DiscordMessageID string
}

func (a *Announcement) ToDTO() dto_response.Announcement {
	return dto_response.Announcement{
		Uuid:             a.Uuid,
		Visibility:       a.Visibility,
		AnnounceAt:       a.AnnounceAt.Unix(),
		DiscordChannelID: a.DiscordChannelID,
		DiscordMessageID: a.DiscordMessageID,
	}
}
