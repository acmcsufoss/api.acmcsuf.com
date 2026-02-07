package repository

import (
	"database/sql"
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/domain"
)

// ---- Event Helper ----
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

func convertDomaintoCreateDBEvent(dEvent *domain.Event) *dbmodels.CreateEventParams {
	return &dbmodels.CreateEventParams{
		Uuid:     dEvent.Uuid,
		Location: dEvent.Location,
		StartAt:  dEvent.StartAt.Unix(),
		EndAt:    dEvent.EndAt.Unix(),
		IsAllDay: dEvent.IsAllDay,
		Host:     dEvent.Host,
	}
}

func convertDomaintoUpdateDBEvent(dEvent *domain.Event) *dbmodels.UpdateEventParams {
	return &dbmodels.UpdateEventParams{
		Uuid:     dEvent.Uuid,
		Location: sql.NullString{String: dEvent.Location, Valid: true},
		StartAt:  sql.NullInt64{Int64: dEvent.StartAt.Unix(), Valid: true},
		EndAt:    sql.NullInt64{Int64: dEvent.EndAt.Unix(), Valid: true},
		IsAllDay: sql.NullBool{Bool: dEvent.IsAllDay, Valid: true},
		Host:     sql.NullString{String: dEvent.Host, Valid: true},
	}
}

// ---- Officer Helper ----
func convertDBOfficerToDomain(dbOfficer *dbmodels.Officer) *domain.Officer {
	return &domain.Officer{
		Uuid:     dbOfficer.Uuid,
		FullName: dbOfficer.FullName,
		Picture:  dbOfficer.Picture.String,
		Github:   dbOfficer.Github.String,
		Discord:  dbOfficer.Discord.String,
	}
}

func convertDomaintoCreateDBOfficer(dOfficer *domain.Officer) *dbmodels.CreateOfficerParams {
	return &dbmodels.CreateOfficerParams{
		Uuid:     dOfficer.Uuid,
		FullName: dOfficer.FullName,
		Picture:  sql.NullString{String: dOfficer.Picture, Valid: true},
		Github:   sql.NullString{String: dOfficer.Github, Valid: true},
		Discord:  sql.NullString{String: dOfficer.Discord, Valid: true},
	}
}

func convertDomaintoUpdateDBOfficer(dOfficer *domain.Officer) *dbmodels.UpdateOfficerParams {
	return &dbmodels.UpdateOfficerParams{
		Uuid:     dOfficer.Uuid,
		FullName: dOfficer.FullName,
		Picture:  sql.NullString{String: dOfficer.Picture, Valid: true},
		Github:   sql.NullString{String: dOfficer.Github, Valid: true},
		Discord:  sql.NullString{String: dOfficer.Discord, Valid: true},
	}
}
