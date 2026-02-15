package mapper

import (
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
	dto_request "github.com/acmcsufoss/api.acmcsuf.com/internal/dto/request"
)

func ToAnnouncementDomain(a *dto_request.Announcement) domain.Announcement {
	return domain.Announcement{
		Uuid:             a.Uuid,
		Visibility:       a.Visibility,
		AnnounceAt:       time.Unix(a.AnnounceAt, 0),
		DiscordChannelID: &a.DiscordChannelID,
		DiscordMessageID: &a.DiscordMessageID,
	}
}

func ToEventDomain(e *dto_request.Event) domain.Event {
	return domain.Event{
		Location: e.Location,
		StartAt:  time.Unix(e.StartAt, 0),
		EndAt:    time.Unix(e.EndAt, 0),
		IsAllDay: e.IsAllDay,
		Host:     e.Host,
	}
}

func ToOfficerDomain(o *dto_request.Officer) domain.Officer {
	return domain.Officer{
		FullName: o.FullName,
		Picture:  &o.Picture,
		Github:   &o.Github,
		Discord:  &o.Discord,
	}
}

func ToPositionDomain(p *dto_request.Position) domain.Position {
	return domain.Position{
		Oid:      p.Oid,
		Semester: p.Semester,
		Tier:     p.Tier,
		FullName: p.FullName,
		Title:    p.Title,
		Team:     p.Team,
	}
}

func ToTierDomain(t *dto_request.Tier) domain.Tier {
	return domain.Tier{
		Tier:   t.Tier,
		Title:  &t.Title,
		Tindex: &t.Tindex,
		Team:   &t.Team,
	}
}
