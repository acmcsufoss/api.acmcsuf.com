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

// ---- Functions to check validity ----
func intToNullInt64(i *int) sql.NullInt64 {
	var val int64
	var valid bool
	if i != nil {
		deref := *i
		val = int64(deref)
	}

	return sql.NullInt64{Int64: val, Valid: valid}
}

func stringToNullString(s *string) sql.NullString {
	var val string
	var valid bool
	if s != nil {
		val = *s
	}

	return sql.NullString{String: val, Valid: valid}
}

func boolToNullBool(b *bool) sql.NullBool {
	var val bool
	var valid bool
	if b != nil {
		val = *b
	}

	return sql.NullBool{Bool: val, Valid: valid}
}

func timeToNullInt64(t *time.Time) sql.NullInt64 {
	var val int64
	var valid bool
	if t != nil {
		deref := *t
		val = deref.Unix()
	}

	return sql.NullInt64{Int64: val, Valid: valid}
}
