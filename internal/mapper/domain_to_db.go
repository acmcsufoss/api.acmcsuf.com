package mapper

import (
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

// File for converting Doamin models into Database models

// ---- Event Converter ----

func ConvertDomainToCreateDBEvent(dEvent domain.Event) dbmodels.CreateEventParams {
	return dbmodels.CreateEventParams{
		Uuid:     dEvent.Uuid,
		Location: dEvent.Location,
		StartAt:  dEvent.StartAt.Unix(),
		EndAt:    dEvent.EndAt.Unix(),
		IsAllDay: dEvent.IsAllDay,
		Host:     dEvent.Host,
	}
}

func ConvertDomainToUpdateDBEvent(dEvent domain.UpdateEvent) dbmodels.UpdateEventParams {
	// -- sql null values --
	loc := stringToNullString(dEvent.Location)

	start := timeToNullInt64(dEvent.StartAt)

	end := timeToNullInt64(dEvent.EndAt)

	allDay := boolToNullBool(dEvent.IsAllDay)

	host := stringToNullString(dEvent.Host)

	return dbmodels.UpdateEventParams{
		Uuid:     dEvent.Uuid,
		Location: loc,
		StartAt:  start,
		EndAt:    end,
		IsAllDay: allDay,
		Host:     host,
	}
}

// ---- Officer Converter ----
func ConvertDomainToCreateDBOfficer(dOfficer domain.Officer) dbmodels.CreateOfficerParams {
	// -- sql null values --
	pic := stringToNullString(dOfficer.Picture)

	github := stringToNullString(dOfficer.Github)

	discord := stringToNullString(dOfficer.Discord)

	return dbmodels.CreateOfficerParams{
		Uuid:     dOfficer.Uuid,
		FullName: dOfficer.FullName,
		Picture:  pic,
		Github:   github,
		Discord:  discord,
	}
}

func ConvertDomainToUpdateDBOfficer(dOfficer domain.UpdateOfficer) dbmodels.UpdateOfficerParams {
	// -- sql null values --
	pic := stringToNullString(dOfficer.Picture)

	github := stringToNullString(dOfficer.Github)

	discord := stringToNullString(dOfficer.Discord)

	return dbmodels.UpdateOfficerParams{
		Uuid:     dOfficer.Uuid,
		FullName: *dOfficer.FullName,
		Picture:  pic,
		Github:   github,
		Discord:  discord,
	}
}

// ---- Announcement Converter ----
func ConvertDomainToCreateDBAnnouncement(dAnnouncement domain.Announcement) dbmodels.CreateAnnouncementParams {
	// -- sql null values --
	chanID := stringToNullString(dAnnouncement.DiscordChannelID)

	msgID := stringToNullString(dAnnouncement.DiscordMessageID)
	return dbmodels.CreateAnnouncementParams{
		Uuid:             dAnnouncement.Uuid,
		Visibility:       dAnnouncement.Visibility,
		AnnounceAt:       dAnnouncement.AnnounceAt.Unix(),
		DiscordChannelID: chanID,
		DiscordMessageID: msgID,
	}
}

func ConvertDomainToUpdateDBAnnouncement(dAnnouncement domain.UpdateAnnouncement) dbmodels.UpdateAnnouncementParams {
	// -- sql null values --
	vis := stringToNullString(dAnnouncement.Visibility)

	announceAt := timeToNullInt64(dAnnouncement.AnnounceAt)

	chanID := stringToNullString(dAnnouncement.DiscordChannelID)

	msgID := stringToNullString(dAnnouncement.DiscordMessageID)

	return dbmodels.UpdateAnnouncementParams{
		Uuid:             dAnnouncement.Uuid,
		Visibility:       vis,
		AnnounceAt:       announceAt,
		DiscordChannelID: chanID,
		DiscordMessageID: msgID,
	}
}

// ---- Tier Converter ----
func ConvertDomainToCreateDBTier(dTier domain.Tier) dbmodels.CreateTierParams {
	// -- sql null values --
	title := stringToNullString(dTier.Title)

	tIdx := intToNullInt64(dTier.Tindex)

	team := stringToNullString(dTier.Team)

	return dbmodels.CreateTierParams{
		Tier:   int64(dTier.Tier),
		Title:  title,
		TIndex: tIdx,
		Team:   team,
	}
}

func ConvertDomainToUpdateDBTier(dTier domain.UpdateTier) dbmodels.UpdateTierParams {
	// -- sql null values --
	title := stringToNullString(dTier.Title)

	tIdx := intToNullInt64(dTier.Tindex)

	team := stringToNullString(dTier.Team)

	return dbmodels.UpdateTierParams{
		Tier:   int64(dTier.Tier),
		Title:  title,
		TIndex: tIdx,
		Team:   team,
	}
}

// ---- Position Converter ----
func ConvertDomainToCreateDBPosition(dPosition domain.Position) dbmodels.CreatePositionParams {
	// -- sql null types --
	title := stringToNullString(dPosition.Title)

	team := stringToNullString(dPosition.Team)

	return dbmodels.CreatePositionParams{
		Oid:      dPosition.Oid,
		Semester: dPosition.Semester,
		Tier:     int64(dPosition.Tier),
		FullName: dPosition.FullName,
		Title:    title,
		Team:     team,
	}
}

func ConvertDomainToUpdateDBPosition(dPosition domain.UpdatePosition) dbmodels.UpdatePositionParams {
	// -- sql null types --
	title := stringToNullString(dPosition.Title)

	team := stringToNullString(dPosition.Team)

	return dbmodels.UpdatePositionParams{
		Oid:      dPosition.Oid,
		Semester: dPosition.Semester,
		Tier:     int64(dPosition.Tier),
		FullName: dPosition.FullName,
		Title:    title,
		Team:     team,
	}
}

func ConvertDomainToDeleteDBPosition(dPosition domain.Position) dbmodels.DeletePositionParams {
	return dbmodels.DeletePositionParams{
		Oid:      dPosition.Oid,
		Semester: dPosition.Semester,
		Tier:     int64(dPosition.Tier),
	}
}
