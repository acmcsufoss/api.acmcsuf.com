package mapper

import (
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/dto"
)

// --- announcement ---
func ToAnnouncementDomain(a *dto.Announcement) domain.Announcement {
	if a == nil {
		return domain.Announcement{}
	}

	return domain.Announcement{
		Uuid:             a.Uuid,
		Visibility:       a.Visibility,
		AnnounceAt:       unixToTime(a.AnnounceAt),
		DiscordChannelID: a.DiscordChannelID,
		DiscordMessageID: a.DiscordMessageID,
	}
}

func ToUpdateAnnouncementDomain(a *dto.UpdateAnnouncement) domain.UpdateAnnouncement {
	if a == nil {
		return domain.UpdateAnnouncement{}
	}

	return domain.UpdateAnnouncement{
		Uuid:             a.Uuid,
		Visibility:       a.Visibility,
		AnnounceAt:       unixToTimePtr(a.AnnounceAt),
		DiscordChannelID: a.DiscordChannelID,
		DiscordMessageID: a.DiscordMessageID,
	}
}

// --- event ---
func ToEventDomain(e *dto.Event) domain.Event {
	if e == nil {
		return domain.Event{}
	}

	return domain.Event{
		Uuid:     e.Uuid,
		Location: e.Location,
		StartAt:  unixToTime(e.StartAt),
		EndAt:    unixToTime(e.EndAt),
		IsAllDay: e.IsAllDay,
		Host:     e.Host,
	}
}

func ToUpdateEventDomain(e *dto.UpdateEvent) domain.UpdateEvent {
	if e == nil {
		return domain.UpdateEvent{}
	}

	return domain.UpdateEvent{
		Uuid:     e.Uuid,
		Location: e.Location,
		StartAt:  unixToTimePtr(e.StartAt),
		EndAt:    unixToTimePtr(e.EndAt),
		IsAllDay: e.IsAllDay,
		Host:     e.Host,
	}
}

// --- officer ---
func ToOfficerDomain(o *dto.Officer) domain.Officer {
	if o == nil {
		return domain.Officer{}
	}

	return domain.Officer{
		Uuid:     o.Uuid,
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
		Uuid:     o.Uuid,
		FullName: o.FullName,
		Picture:  o.Picture,
		Github:   o.Github,
		Discord:  o.Discord,
	}
}

// --- position ---
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

// --- tier ---
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

// --- helpers ---
func unixToTime(v int64) time.Time {
	return time.Unix(v, 0)
}

func unixToTimePtr(v *int64) *time.Time {
	if v == nil {
		return nil
	}
	t := time.Unix(*v, 0)
	return &t
}
