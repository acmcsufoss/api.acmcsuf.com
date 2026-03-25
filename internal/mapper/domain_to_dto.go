package mapper

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/dto"
)

// File deticated to mapping domain models to dto response models

// --- announcements ---
func ToAnnouncementDTO(a *domain.Announcement) dto.Announcement {
	return dto.Announcement{
		Uuid:             a.Uuid,
		Visibility:       a.Visibility,
		AnnounceAt:       a.AnnounceAt.Unix(),
		DiscordChannelID: a.DiscordChannelID,
		DiscordMessageID: a.DiscordMessageID,
	}
}

// --- event ---
func ToEventEventDTO(e *domain.Event) dto.Event {
	return dto.Event{
		Uuid:     e.Uuid,
		Location: e.Location,
		StartAt:  e.StartAt.Unix(),
		EndAt:    e.EndAt.Unix(),
		IsAllDay: e.IsAllDay,
		Host:     e.Host,
	}
}

// --- offcer ---
func ToOfficerDTO(o domain.Officer) dto.Officer {
	return dto.Officer{
		Uuid:     o.Uuid,
		FullName: o.FullName,
		Picture:  o.Picture,
		Github:   o.Github,
		Discord:  o.Discord,
	}
}

// --- postition ---
func ToPositionDTO(p *domain.Position) dto.Position {
	return dto.Position{
		Oid:      p.Oid,
		Semester: p.Semester,
		Tier:     p.Tier,
		FullName: p.FullName,
		Title:    p.Title,
		Team:     p.Team,
	}
}

// --- tier ---
func ToTierDTO(t *domain.Tier) dto.Tier {
	return dto.Tier{
		Tier:   t.Tier,
		Title:  t.Title,
		Tindex: t.Tindex,
		Team:   t.Team,
	}
}
