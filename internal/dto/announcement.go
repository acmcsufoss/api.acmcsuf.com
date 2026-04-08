package dto

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

type Announcement struct {
	Uuid             string  `json:"uuid"`
	Visibility       string  `json:"visibility"`
	AnnounceAt       int64   `json:"announce_at"`
	DiscordChannelID *string `json:"discord_channel_id"`
	DiscordMessageID *string `json:"discord_message_id"`
}

func (a *Announcement) ToDomain() domain.Announcement {
	if a == nil {
		return domain.Announcement{}
	}

	return domain.Announcement{
		Uuid:             a.Uuid,
		Visibility:       a.Visibility,
		AnnounceAt:       utils.UnixToTime(a.AnnounceAt),
		DiscordChannelID: a.DiscordChannelID,
		DiscordMessageID: a.DiscordMessageID,
	}
}

func AnnouncementDomainToDto(a *domain.Announcement) Announcement {
	return Announcement{
		Uuid:             a.Uuid,
		Visibility:       a.Visibility,
		AnnounceAt:       a.AnnounceAt.Unix(),
		DiscordChannelID: a.DiscordChannelID,
		DiscordMessageID: a.DiscordMessageID,
	}
}

type UpdateAnnouncement struct {
	Uuid             string  `json:"uuid"`
	Visibility       *string `json:"visibility"`
	AnnounceAt       *int64  `json:"announce_at"`
	DiscordChannelID *string `json:"discord_channel_id"`
	DiscordMessageID *string `json:"discord_message_id"`
}
