package repository

import (
	"database/sql"
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

// File for converting Database models into Doamin models

// ---- Event Converter ----
func convertDBEventToDomain(dbEvent *dbmodels.Event) *domain.Event {
	return &domain.Event{
		Uuid:     dbEvent.Uuid,
		Location: dbEvent.Location,
		StartAt:  time.Unix(dbEvent.StartAt, 0),
		EndAt:    time.Unix(dbEvent.EndAt, 0),
		IsAllDay: dbEvent.IsAllDay,
		Host:     dbEvent.Host,
	}
}

func convertDomainToCreateDBEvent(dEvent *domain.Event) *dbmodels.CreateEventParams {
	return &dbmodels.CreateEventParams{
		Uuid:     dEvent.Uuid,
		Location: dEvent.Location,
		StartAt:  dEvent.StartAt.Unix(),
		EndAt:    dEvent.EndAt.Unix(),
		IsAllDay: dEvent.IsAllDay,
		Host:     dEvent.Host,
	}
}

func convertDomainToUpdateDBEvent(dEvent *domain.Event) *dbmodels.UpdateEventParams {
	return &dbmodels.UpdateEventParams{
		Uuid:     dEvent.Uuid,
		Location: sql.NullString{String: dEvent.Location, Valid: true},
		StartAt:  sql.NullInt64{Int64: dEvent.StartAt.Unix(), Valid: true},
		EndAt:    sql.NullInt64{Int64: dEvent.EndAt.Unix(), Valid: true},
		IsAllDay: sql.NullBool{Bool: dEvent.IsAllDay, Valid: true},
		Host:     sql.NullString{String: dEvent.Host, Valid: true},
	}
}

// ---- Officer Converter ----
func convertDBOfficerToDomain(dbOfficer *dbmodels.Officer) *domain.Officer {
	return &domain.Officer{
		Uuid:     dbOfficer.Uuid,
		FullName: dbOfficer.FullName,
		Picture:  dbOfficer.Picture.String,
		Github:   dbOfficer.Github.String,
		Discord:  dbOfficer.Discord.String,
	}
}

func convertDomainToCreateDBOfficer(dOfficer *domain.Officer) *dbmodels.CreateOfficerParams {
	return &dbmodels.CreateOfficerParams{
		Uuid:     dOfficer.Uuid,
		FullName: dOfficer.FullName,
		Picture:  sql.NullString{String: dOfficer.Picture, Valid: true},
		Github:   sql.NullString{String: dOfficer.Github, Valid: true},
		Discord:  sql.NullString{String: dOfficer.Discord, Valid: true},
	}
}

func convertDomainToUpdateDBOfficer(dOfficer *domain.Officer) *dbmodels.UpdateOfficerParams {
	return &dbmodels.UpdateOfficerParams{
		Uuid:     dOfficer.Uuid,
		FullName: dOfficer.FullName,
		Picture:  sql.NullString{String: dOfficer.Picture, Valid: true},
		Github:   sql.NullString{String: dOfficer.Github, Valid: true},
		Discord:  sql.NullString{String: dOfficer.Discord, Valid: true},
	}
}

// ---- Announcement Converter ----
func convertDBAnnouncementToDomain(dbAnnouncement *dbmodels.Announcement) *domain.Announcement {
	return &domain.Announcement{
		Uuid:             dbAnnouncement.Uuid,
		Visibility:       dbAnnouncement.Visibility,
		AnnounceAt:       time.Unix(dbAnnouncement.AnnounceAt, 0),
		DiscordChannelID: dbAnnouncement.DiscordChannelID.String,
		DiscordMessageID: dbAnnouncement.DiscordMessageID.String,
	}
}

func convertDomainToCreateDBAnnouncement(dAnnouncement *domain.Announcement) *dbmodels.CreateAnnouncementParams {
	return &dbmodels.CreateAnnouncementParams{
		Uuid:             dAnnouncement.Uuid,
		Visibility:       dAnnouncement.Visibility,
		AnnounceAt:       dAnnouncement.AnnounceAt.Unix(),
		DiscordChannelID: sql.NullString{String: dAnnouncement.DiscordChannelID, Valid: true},
		DiscordMessageID: sql.NullString{String: dAnnouncement.DiscordMessageID, Valid: true},
	}
}

func convertDomainToUpdateDBAnnouncement(dAnnouncement *domain.Announcement) *dbmodels.UpdateAnnouncementParams {
	return &dbmodels.UpdateAnnouncementParams{
		Uuid:             dAnnouncement.Uuid,
		Visibility:       sql.NullString{String: dAnnouncement.Visibility, Valid: true},
		AnnounceAt:       sql.NullInt64{Int64: dAnnouncement.AnnounceAt.Unix(), Valid: true},
		DiscordChannelID: sql.NullString{String: dAnnouncement.DiscordChannelID, Valid: true},
		DiscordMessageID: sql.NullString{String: dAnnouncement.DiscordMessageID, Valid: true},
	}
}

// ---- Tier Converter ----
func convertDBTierToDomain(dbTier *dbmodels.Tier) *domain.Tier {
	return &domain.Tier{
		Tier:   int(dbTier.Tier),
		Title:  dbTier.Title.String,
		Tindex: int(dbTier.TIndex.Int64),
		Team:   dbTier.Team.String,
	}
}

func convertDomainToCreateDBTier(dTier *domain.Tier) *dbmodels.CreateTierParams {
	return &dbmodels.CreateTierParams{
		Tier:   int64(dTier.Tier),
		Title:  sql.NullString{String: dTier.Title, Valid: true},
		TIndex: sql.NullInt64{Int64: int64(dTier.Tindex), Valid: true},
		Team:   sql.NullString{String: dTier.Team, Valid: true},
	}
}

func convertDomainToUpdateDBTier(dTier *domain.Tier) *dbmodels.UpdateTierParams {
	return &dbmodels.UpdateTierParams{
		Tier:   int64(dTier.Tier),
		Title:  sql.NullString{String: dTier.Title, Valid: true},
		TIndex: sql.NullInt64{Int64: int64(dTier.Tindex), Valid: true},
		Team:   sql.NullString{String: dTier.Team, Valid: true},
	}
}

// ---- Position Converter ----
func convertDBPositionToDomain(dbPosition *dbmodels.Position) *domain.Position {
	return &domain.Position{
		Oid:      dbPosition.Oid,
		Semester: dbPosition.Semester,
		Tier:     int(dbPosition.Tier),
	}
}

func convertDomainToCreateDBPosition(dPositon *domain.Position) *dbmodels.CreatePositionParams {
	return &dbmodels.CreatePositionParams{
		Oid:      dPositon.Oid,
		Semester: dPositon.Semester,
		Tier:     int64(dPositon.Tier),
	}
}

func convertDomainToUpdateDBPosition(dPositon *domain.Position) *dbmodels.UpdatePositionParams {
	return &dbmodels.UpdatePositionParams{
		Oid:      dPositon.Oid,
		Semester: dPositon.Semester,
		Tier:     int64(dPositon.Tier),
	}
}

func convertDomainToDeleteDBPosition(dPositon *domain.Position) *dbmodels.DeletePositionParams {
	return &dbmodels.DeletePositionParams{
		Oid:      dPositon.Oid,
		Semester: dPositon.Semester,
		Tier:     int64(dPositon.Tier),
	}
}
