package mapper

import (
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

// File for converting Database models into Doamin models

// ---- Event Converter ----
func ConvertDBEventToDomain(dbEvent dbmodels.Event) domain.Event {
	return domain.Event{
		Uuid:     dbEvent.Uuid,
		Location: dbEvent.Location,
		StartAt:  time.Unix(dbEvent.StartAt, 0),
		EndAt:    time.Unix(dbEvent.EndAt, 0),
		IsAllDay: dbEvent.IsAllDay,
		Host:     dbEvent.Host,
	}
}

// ---- Officer Converter ----
func ConvertDBOfficerToDomain(dbOfficer dbmodels.Officer) domain.Officer {
	return domain.Officer{
		Uuid:     dbOfficer.Uuid,
		FullName: dbOfficer.FullName,
		Picture:  &dbOfficer.Picture.String,
		Github:   &dbOfficer.Github.String,
		Discord:  &dbOfficer.Discord.String,
	}
}

// ---- Announcement Converter ----
func ConvertDBAnnouncementToDomain(dbAnnouncement dbmodels.Announcement) domain.Announcement {
	return domain.Announcement{
		Uuid:             dbAnnouncement.Uuid,
		Visibility:       dbAnnouncement.Visibility,
		AnnounceAt:       time.Unix(dbAnnouncement.AnnounceAt, 0),
		DiscordChannelID: &dbAnnouncement.DiscordChannelID.String,
		DiscordMessageID: &dbAnnouncement.DiscordMessageID.String,
	}
}

// ---- Tier Converter ----
func ConvertDBTierToDomain(dbTier dbmodels.Tier) domain.Tier {
	// note: &int(exp) / &(int)(exp) is illegal, so it is split into v and then &v
	v := int(dbTier.TIndex.Int64)
	return domain.Tier{
		Tier:   int(dbTier.Tier),
		Title:  &dbTier.Title.String,
		Tindex: &v,
		Team:   &dbTier.Team.String,
	}
}

// ---- Position Converter ----
func ConvertDBPositionToDomain(dbPosition dbmodels.Position) domain.Position {
	return domain.Position{
		Oid:      dbPosition.Oid,
		Semester: dbPosition.Semester,
		Tier:     int(dbPosition.Tier),
		FullName: dbPosition.FullName,
		Title:    &dbPosition.Team.String,
		Team:     &dbPosition.Team.String,
	}
}
