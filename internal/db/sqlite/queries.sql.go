// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package sqlite

import (
	"context"
	"database/sql"
)

const createAnnouncement = `-- name: CreateAnnouncement :exec
INSERT INTO
    announcement (
        uuid,
        event_groups_group_uuid,
        approved_by_list_uuid,
        visibility,
        announce_at,
        discord_channel_id,
        discord_message_id
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?)
`

type CreateAnnouncementParams struct {
	Uuid                 string         `json:"uuid"`
	EventGroupsGroupUuid sql.NullString `json:"event_groups_group_uuid"`
	ApprovedByListUuid   sql.NullString `json:"approved_by_list_uuid"`
	Visibility           string         `json:"visibility"`
	AnnounceAt           int64          `json:"announce_at"`
	DiscordChannelID     sql.NullString `json:"discord_channel_id"`
	DiscordMessageID     sql.NullString `json:"discord_message_id"`
}

func (q *Queries) CreateAnnouncement(ctx context.Context, arg CreateAnnouncementParams) error {
	_, err := q.db.ExecContext(ctx, createAnnouncement,
		arg.Uuid,
		arg.EventGroupsGroupUuid,
		arg.ApprovedByListUuid,
		arg.Visibility,
		arg.AnnounceAt,
		arg.DiscordChannelID,
		arg.DiscordMessageID,
	)
	return err
}

const createEvent = `-- name: CreateEvent :exec
INSERT INTO
    event (
        uuid,
        location,
        start_at,
        end_at,
        is_all_day,
        host,
        visibility
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?)
`

type CreateEventParams struct {
	Uuid       string      `json:"uuid"`
	Location   string      `json:"location"`
	StartAt    interface{} `json:"start_at"`
	EndAt      interface{} `json:"end_at"`
	IsAllDay   bool        `json:"is_all_day"`
	Host       string      `json:"host"`
	Visibility string      `json:"visibility"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) error {
	_, err := q.db.ExecContext(ctx, createEvent,
		arg.Uuid,
		arg.Location,
		arg.StartAt,
		arg.EndAt,
		arg.IsAllDay,
		arg.Host,
		arg.Visibility,
	)
	return err
}

const createPerson = `-- name: CreatePerson :exec
INSERT INTO
    person (uuid, name, preferred_pronoun)
VALUES
    (?, ?, ?)
`

type CreatePersonParams struct {
	Uuid             sql.NullString `json:"uuid"`
	Name             sql.NullString `json:"name"`
	PreferredPronoun sql.NullString `json:"preferred_pronoun"`
}

func (q *Queries) CreatePerson(ctx context.Context, arg CreatePersonParams) error {
	_, err := q.db.ExecContext(ctx, createPerson, arg.Uuid, arg.Name, arg.PreferredPronoun)
	return err
}

const getAnnouncement = `-- name: GetAnnouncement :exec
SELECT
    uuid,
    event_groups_group_uuid,
    approved_by_list_uuid,
    visibility,
    announce_at,
    discord_channel_id,
    discord_message_id
from
    announcement
where
    uuid = ?
`

func (q *Queries) GetAnnouncement(ctx context.Context, uuid string) error {
	_, err := q.db.ExecContext(ctx, getAnnouncement, uuid)
	return err
}

const getEvent = `-- name: GetEvent :exec
SELECT
    uuid,
    location,
    start_at,
    end_at,
    is_all_day,
    host,
    visibility
from
    event
where
    uuid = ?
`

func (q *Queries) GetEvent(ctx context.Context, uuid string) error {
	_, err := q.db.ExecContext(ctx, getEvent, uuid)
	return err
}

const getPerson = `-- name: GetPerson :exec
SELECT
    uuid,
    name,
    preferred_pronoun
from
    person
where
    uuid = ?
`

func (q *Queries) GetPerson(ctx context.Context, uuid sql.NullString) error {
	_, err := q.db.ExecContext(ctx, getPerson, uuid)
	return err
}