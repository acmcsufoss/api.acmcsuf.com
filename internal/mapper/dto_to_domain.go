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
		DiscordChannelID: a.DiscordChannelID,
		DiscordMessageID: a.DiscordMessageID,
	}
}

func ToUpdateAnnouncementDomain(a *dto_request.UpdateAnnouncement) domain.UpdateAnnouncement {
	var announceAt time.Time
	if a.AnnounceAt != nil {
		announceAt = time.Unix(*a.AnnounceAt, 0)
	}
	return domain.UpdateAnnouncement{
		Uuid:             a.Uuid,
		Visibility:       a.Visibility,
		AnnounceAt:       &announceAt,
		DiscordChannelID: a.DiscordChannelID,
		DiscordMessageID: a.DiscordMessageID,
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

func ToUpdateEventDomain(e *dto_request.UpdateEvent) domain.UpdateEvent {
	startAt := time.Unix(*e.StartAt, 0)
	endAt := time.Unix(*e.EndAt, 0)
	return domain.UpdateEvent{
		Location: e.Location,
		StartAt:  &startAt,
		EndAt:    &endAt,
		IsAllDay: e.IsAllDay,
		Host:     e.Host,
	}
}

func ToOfficerDomain(o *dto_request.Officer) domain.Officer {
	return domain.Officer{
		FullName: o.FullName,
		Picture:  o.Picture,
		Github:   o.Github,
		Discord:  o.Discord,
	}
}

func ToUpdateOfficerDomain(o *dto_request.UpdateOfficer) domain.UpdateOfficer {
	return domain.UpdateOfficer{
		FullName: o.FullName,
		Picture:  o.Picture,
		Github:   o.Github,
		Discord:  o.Discord,
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

func ToUpdatePositionDomain(p *dto_request.UpdatePosition) domain.UpdatePosition {
	return domain.UpdatePosition{
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
		Title:  t.Title,
		Tindex: t.Tindex,
		Team:   t.Team,
	}
}

func ToUpdateTierDomain(t *dto_request.UpdateTier) domain.UpdateTier {
	return domain.UpdateTier{
		Tier:   t.Tier,
		Title:  t.Title,
		Tindex: t.Tindex,
		Team:   t.Team,
	}
}
