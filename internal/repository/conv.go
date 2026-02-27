package repository

import (
	"database/sql"
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

// File for converting Database models into Doamin models

// ---- Event Converter ----
func convertDBEventToDomain(dbEvent dbmodels.Event) domain.Event {
	return domain.Event{
		Uuid:     dbEvent.Uuid,
		Location: dbEvent.Location,
		StartAt:  time.Unix(dbEvent.StartAt, 0),
		EndAt:    time.Unix(dbEvent.EndAt, 0),
		IsAllDay: dbEvent.IsAllDay,
		Host:     dbEvent.Host,
	}
}

func convertDomainToCreateDBEvent(dEvent domain.Event) dbmodels.CreateEventParams {
	return dbmodels.CreateEventParams{
		Uuid:     dEvent.Uuid,
		Location: dEvent.Location,
		StartAt:  dEvent.StartAt.Unix(),
		EndAt:    dEvent.EndAt.Unix(),
		IsAllDay: dEvent.IsAllDay,
		Host:     dEvent.Host,
	}
}

func convertDomainToUpdateDBEvent(dEvent domain.UpdateEvent) dbmodels.UpdateEventParams {
	// -- sql null values --
	var loc string
	if dEvent.Location != nil {
		loc = *dEvent.Location
	}

	var start int64
	if dEvent.StartAt != nil {
		start = dEvent.StartAt.Unix()
	}

	var end int64
	if dEvent.EndAt != nil {
		end = dEvent.StartAt.Unix()
	}

	var allDay bool
	if dEvent.IsAllDay != nil {
		allDay = *dEvent.IsAllDay
	}

	var host string
	if dEvent.Host != nil {
		host = *dEvent.Host
	}

	return dbmodels.UpdateEventParams{
		Uuid:     dEvent.Uuid,
		Location: sql.NullString{String: loc, Valid: validString(dEvent.Location)},
		StartAt:  sql.NullInt64{Int64: start, Valid: validTime(dEvent.StartAt)},
		EndAt:    sql.NullInt64{Int64: end, Valid: validTime(dEvent.EndAt)},
		IsAllDay: sql.NullBool{Bool: allDay, Valid: validBool(dEvent.IsAllDay)},
		Host:     sql.NullString{String: host, Valid: validString(dEvent.Host)},
	}
}

// ---- Officer Converter ----
func convertDBOfficerToDomain(dbOfficer dbmodels.Officer) domain.Officer {
	return domain.Officer{
		Uuid:     dbOfficer.Uuid,
		FullName: dbOfficer.FullName,
		Picture:  &dbOfficer.Picture.String,
		Github:   &dbOfficer.Github.String,
		Discord:  &dbOfficer.Discord.String,
	}
}

func convertDomainToCreateDBOfficer(dOfficer domain.Officer) dbmodels.CreateOfficerParams {
	// -- sql null values --
	pic := dOfficer.Picture
	github := dOfficer.Github
	discord := dOfficer.Discord

	return dbmodels.CreateOfficerParams{
		Uuid:     dOfficer.Uuid,
		FullName: dOfficer.FullName,
		Picture:  sql.NullString{String: *pic, Valid: validString(pic)},
		Github:   sql.NullString{String: *github, Valid: validString(pic)},
		Discord:  sql.NullString{String: *discord, Valid: validString(pic)},
	}
}

func convertDomainToUpdateDBOfficer(dOfficer domain.UpdateOfficer) dbmodels.UpdateOfficerParams {
	// -- sql null values --
	pic := dOfficer.Picture
	github := dOfficer.Github
	discord := dOfficer.Discord

	return dbmodels.UpdateOfficerParams{
		Uuid:     dOfficer.Uuid,
		FullName: *dOfficer.FullName,
		Picture:  sql.NullString{String: *pic, Valid: validString(pic)},
		Github:   sql.NullString{String: *github, Valid: validString(github)},
		Discord:  sql.NullString{String: *discord, Valid: validString(discord)},
	}
}

// ---- Announcement Converter ----
func convertDBAnnouncementToDomain(dbAnnouncement dbmodels.Announcement) domain.Announcement {
	return domain.Announcement{
		Uuid:             dbAnnouncement.Uuid,
		Visibility:       dbAnnouncement.Visibility,
		AnnounceAt:       time.Unix(dbAnnouncement.AnnounceAt, 0),
		DiscordChannelID: &dbAnnouncement.DiscordChannelID.String,
		DiscordMessageID: &dbAnnouncement.DiscordMessageID.String,
	}
}

func convertDomainToCreateDBAnnouncement(dAnnouncement domain.Announcement) dbmodels.CreateAnnouncementParams {
	// -- sql null values --
	chanID := dAnnouncement.DiscordChannelID
	msgID := dAnnouncement.DiscordMessageID

	return dbmodels.CreateAnnouncementParams{
		Uuid:             dAnnouncement.Uuid,
		Visibility:       dAnnouncement.Visibility,
		AnnounceAt:       dAnnouncement.AnnounceAt.Unix(),
		DiscordChannelID: sql.NullString{String: *chanID, Valid: validString(chanID)},
		DiscordMessageID: sql.NullString{String: *msgID, Valid: validString(msgID)},
	}
}

func convertDomainToUpdateDBAnnouncement(dAnnouncement domain.UpdateAnnouncement) dbmodels.UpdateAnnouncementParams {
	// -- sql null values --

	var vis string
	if dAnnouncement.Visibility != nil {
		vis = *dAnnouncement.Visibility
	}
	var announceAt int64
	announceAtPtr := dAnnouncement.AnnounceAt
	if announceAtPtr != nil {
		announceAt = announceAtPtr.Unix()
	}
	var chanID string
	if dAnnouncement.DiscordChannelID != nil {
		chanID = *dAnnouncement.DiscordChannelID
	}

	var msgID string
	if dAnnouncement.DiscordMessageID != nil {
		msgID = *dAnnouncement.DiscordMessageID
	}

	return dbmodels.UpdateAnnouncementParams{
		Uuid:             dAnnouncement.Uuid,
		Visibility:       sql.NullString{String: vis, Valid: validString(dAnnouncement.Visibility)},
		AnnounceAt:       sql.NullInt64{Int64: announceAt, Valid: validTime(announceAtPtr)},
		DiscordChannelID: sql.NullString{String: chanID, Valid: validString(dAnnouncement.DiscordChannelID)},
		DiscordMessageID: sql.NullString{String: msgID, Valid: validString(dAnnouncement.DiscordMessageID)},
	}
}

// ---- Tier Converter ----
func convertDBTierToDomain(dbTier dbmodels.Tier) domain.Tier {
	// note: &int(exp) / &(int)(exp) is illegal, so it is split into v and then &v
	v := int(dbTier.TIndex.Int64)
	return domain.Tier{
		Tier:   int(dbTier.Tier),
		Title:  &dbTier.Title.String,
		Tindex: &v,
		Team:   &dbTier.Team.String,
	}
}

func convertDomainToCreateDBTier(dTier domain.Tier) dbmodels.CreateTierParams {
	// -- sql null values --
	title := dTier.Title
	tIdx := int64(*dTier.Tindex)
	team := dTier.Team

	return dbmodels.CreateTierParams{
		Tier:   int64(dTier.Tier),
		Title:  sql.NullString{String: *title, Valid: validString(title)},
		TIndex: sql.NullInt64{Int64: tIdx, Valid: validInt64(&tIdx)},
		Team:   sql.NullString{String: *team, Valid: validString(team)},
	}
}

func convertDomainToUpdateDBTier(dTier domain.UpdateTier) dbmodels.UpdateTierParams {
	// -- sql null values --
	title := dTier.Title
	tIdx := int64(*dTier.Tindex)
	team := dTier.Team

	return dbmodels.UpdateTierParams{
		Tier:   int64(dTier.Tier),
		Title:  sql.NullString{String: *title, Valid: validString(title)},
		TIndex: sql.NullInt64{Int64: tIdx, Valid: validInt64(&tIdx)},
		Team:   sql.NullString{String: *team, Valid: validString(team)},
	}
}

// ---- Position Converter ----
func convertDBPositionToDomain(dbPosition dbmodels.Position) domain.Position {
	return domain.Position{
		Oid:      dbPosition.Oid,
		Semester: dbPosition.Semester,
		Tier:     int(dbPosition.Tier),
		FullName: dbPosition.FullName,
		Title:    &dbPosition.Team.String,
		Team:     &dbPosition.Team.String,
	}
}

func convertDomainToCreateDBPosition(dPosition domain.Position) dbmodels.CreatePositionParams {
	// -- sql null types --
	title := dPosition.Title
	team := dPosition.Team

	return dbmodels.CreatePositionParams{
		Oid:      dPosition.Oid,
		Semester: dPosition.Semester,
		Tier:     int64(dPosition.Tier),
		FullName: dPosition.FullName,
		Title:    sql.NullString{String: *title, Valid: validString(title)},
		Team:     sql.NullString{String: *team, Valid: validString(team)},
	}
}

func convertDomainToUpdateDBPosition(dPosition domain.UpdatePosition) dbmodels.UpdatePositionParams {
	// -- sql null types --
	title := dPosition.Title
	team := dPosition.Team

	return dbmodels.UpdatePositionParams{
		Oid:      dPosition.Oid,
		Semester: dPosition.Semester,
		Tier:     int64(dPosition.Tier),
		FullName: dPosition.FullName,
		Title:    sql.NullString{String: *title, Valid: validString(title)},
		Team:     sql.NullString{String: *team, Valid: validString(team)},
	}
}

func convertDomainToDeleteDBPosition(dPosition domain.Position) dbmodels.DeletePositionParams {
	return dbmodels.DeletePositionParams{
		Oid:      dPosition.Oid,
		Semester: dPosition.Semester,
		Tier:     int64(dPosition.Tier),
	}
}

// ---- Functions to check validity ----
func int64ToNullInt64(i *int64) sql.NullInt64 {
	var val int64
	var valid bool
	if i != nil {
		val = *i
	}

	return sql.NullInt64{Int64: val, Valid: valid}
}

func validString(s *string) sql.NullString {
	var val string
	var valid bool
	if s != nil {
		val = *s
	}

	return sql.NullString{String: val, Valid: valid}
}

func validBool(b *bool) sql.NullBool {
	var val bool
	var valid bool
	if b != nil {
		val = *b
	}

	return sql.NullBool{Bool: val, Valid: valid}
}

func validTime(t *time.Time) sql.NullInt64 {
	var val int64
	var valid bool
	if t != nil {
		deref := *t
		val = deref.Unix()
	}

	return sql.NullInt64{Int64: val, Valid: valid}
}
