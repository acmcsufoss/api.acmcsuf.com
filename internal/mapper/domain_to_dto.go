package mapper

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"

	dto_response "github.com/acmcsufoss/api.acmcsuf.com/internal/dto/response"
)

// File deticated to mapping domain models to dto response models

func ToAnnouncementDTO(a *domain.Announcement) dto_response.Announcement {
	return dto_response.Announcement{
		Uuid:             a.Uuid,
		Visibility:       a.Visibility,
		AnnounceAt:       a.AnnounceAt.Unix(),
		DiscordChannelID: *a.DiscordChannelID,
		DiscordMessageID: *a.DiscordMessageID,
	}
}

func ToEventEventDTO(e *domain.Event) dto_response.Event {
	return dto_response.Event{
		Uuid:     e.Uuid,
		Location: e.Location,
		StartAt:  e.StartAt.Unix(),
		EndAt:    e.EndAt.Unix(),
		IsAllDay: e.IsAllDay,
		Host:     e.Host,
	}
}

func ToOfficerDTO(o domain.Officer) dto_response.Officer {
	return dto_response.Officer{
		Uuid:     o.Uuid,
		FullName: o.FullName,
		Picture:  *o.Picture,
		Github:   *o.Github,
		Discord:  *o.Discord,
	}
}

func ToPositionDTO(p *domain.Position) dto_response.Position {
	return dto_response.Position{
		Oid:      p.Oid,
		Semester: p.Semester,
		Tier:     p.Tier,
		FullName: p.FullName,
		Title:    p.Title,
		Team:     p.Team,
	}
}

func ToTierDTO(t *domain.Tier) dto_response.Tier {
	return dto_response.Tier{
		Tier:   t.Tier,
		Title:  *t.Title,
		Tindex: *t.Tindex,
		Team:   *t.Team,
	}
}
