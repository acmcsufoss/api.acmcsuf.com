package mapper

import (
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/dto"
)

func ToAnnouncementDomain(a *dto.Announcement) domain.Announcement {
	if a == nil {
		return domain.Announcement{}
	}

	return domain.Announcement{
		Uuid:             a.Uuid,
		Visibility:       a.Visibility,
		AnnounceAt:       time.Unix(a.AnnounceAt, 0),
		DiscordChannelID: a.DiscordChannelID,
		DiscordMessageID: a.DiscordMessageID,
	}
}

func ToUpdateAnnouncementDomain(a *dto.UpdateAnnouncement) domain.UpdateAnnouncement {
	if a == nil {
		return domain.UpdateAnnouncement{}
	}

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

func ToEventDomain(e *dto.Event) domain.Event {
	if e == nil {
		return domain.Event{}
	}

	return domain.Event{
		Uuid:     e.Uuid,
		Location: e.Location,
		StartAt:  time.Unix(e.StartAt, 0),
		EndAt:    time.Unix(e.EndAt, 0),
		IsAllDay: e.IsAllDay,
		Host:     e.Host,
	}
}

func ToUpdateEventDomain(e *dto.UpdateEvent) domain.UpdateEvent {
	if e == nil {
		return domain.UpdateEvent{}
	}

	startAt := time.Unix(*e.StartAt, 0)
	endAt := time.Unix(*e.EndAt, 0)
	return domain.UpdateEvent{
		Uuid:     e.Uuid,
		Location: e.Location,
		StartAt:  &startAt,
		EndAt:    &endAt,
		IsAllDay: e.IsAllDay,
		Host:     e.Host,
	}
}

func ToOfficerDomain(o *dto.Officer) domain.Officer {
	if o == nil {
		return domain.Officer{}
	}

	return domain.Officer{
		FullName: o.FullName,
		Picture:  o.Picture,
		Github:   o.Github,
		Discord:  o.Discord,
	}
}

func ToUpdateOfficerDomain(o *dto.UpdateOfficer) domain.UpdateOfficer {
	if o == nil {
		return domain.UpdateOfficer{}
	}

	return domain.UpdateOfficer{
		FullName: o.FullName,
		Picture:  o.Picture,
		Github:   o.Github,
		Discord:  o.Discord,
	}
}

func ToPositionDomain(p *dto.Position) domain.Position {
	if p == nil {
		return domain.Position{}
	}

	return domain.Position{
		Oid:      p.Oid,
		Semester: p.Semester,
		Tier:     p.Tier,
		FullName: p.FullName,
		Title:    p.Title,
		Team:     p.Team,
	}
}

func ToUpdatePositionDomain(p *dto.UpdatePosition) domain.UpdatePosition {
	if p == nil {
		return domain.UpdatePosition{}
	}

	return domain.UpdatePosition{
		Oid:      p.Oid,
		Semester: p.Semester,
		Tier:     p.Tier,
		FullName: p.FullName,
		Title:    p.Title,
		Team:     p.Team,
	}
}

func ToTierDomain(t *dto.Tier) domain.Tier {
	if t == nil {
		return domain.Tier{}
	}

	return domain.Tier{
		Tier:   t.Tier,
		Title:  t.Title,
		Tindex: t.Tindex,
		Team:   t.Team,
	}
}

func ToUpdateTierDomain(t *dto.UpdateTier) domain.UpdateTier {
	if t == nil {
		return domain.UpdateTier{}
	}

	return domain.UpdateTier{
		Tier:   t.Tier,
		Title:  t.Title,
		Tindex: t.Tindex,
		Team:   t.Team,
	}
}
