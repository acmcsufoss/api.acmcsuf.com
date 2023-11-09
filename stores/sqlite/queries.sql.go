// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: queries.sql

package sqlite

import (
	"context"
	"database/sql"
)

const createAnnouncement = `-- name: CreateAnnouncement :exec
INSERT INTO announcements (id, event_list_id, approved_by_list_id, visibility, announce_at, discord_channel_id, discord_message_id) VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateAnnouncementParams struct {
	ID               string         `json:"id"`
	EventListID      sql.NullString `json:"event_list_id"`
	ApprovedByListID sql.NullString `json:"approved_by_list_id"`
	Visibility       string         `json:"visibility"`
	AnnounceAt       int64          `json:"announce_at"`
	DiscordChannelID sql.NullString `json:"discord_channel_id"`
	DiscordMessageID sql.NullString `json:"discord_message_id"`
}

func (q *Queries) CreateAnnouncement(ctx context.Context, arg CreateAnnouncementParams) error {
	_, err := q.db.ExecContext(ctx, createAnnouncement,
		arg.ID,
		arg.EventListID,
		arg.ApprovedByListID,
		arg.Visibility,
		arg.AnnounceAt,
		arg.DiscordChannelID,
		arg.DiscordMessageID,
	)
	return err
}

const createEvent = `-- name: CreateEvent :exec
INSERT INTO events (id, location, start_at, duration_ms, is_all_day, host, visibility) VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateEventParams struct {
	ID         string      `json:"id"`
	Location   string      `json:"location"`
	StartAt    interface{} `json:"start_at"`
	DurationMs interface{} `json:"duration_ms"`
	IsAllDay   bool        `json:"is_all_day"`
	Host       string      `json:"host"`
	Visibility string      `json:"visibility"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) error {
	_, err := q.db.ExecContext(ctx, createEvent,
		arg.ID,
		arg.Location,
		arg.StartAt,
		arg.DurationMs,
		arg.IsAllDay,
		arg.Host,
		arg.Visibility,
	)
	return err
}

const createResource = `-- name: CreateResource :exec
INSERT INTO resources (id, title, content_md, image_url, resource_type, resource_list_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateResourceParams struct {
	ID             string         `json:"id"`
	Title          string         `json:"title"`
	ContentMd      string         `json:"content_md"`
	ImageUrl       sql.NullString `json:"image_url"`
	ResourceType   string         `json:"resource_type"`
	ResourceListID sql.NullString `json:"resource_list_id"`
	CreatedAt      int64          `json:"created_at"`
	UpdatedAt      int64          `json:"updated_at"`
}

func (q *Queries) CreateResource(ctx context.Context, arg CreateResourceParams) error {
	_, err := q.db.ExecContext(ctx, createResource,
		arg.ID,
		arg.Title,
		arg.ContentMd,
		arg.ImageUrl,
		arg.ResourceType,
		arg.ResourceListID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const createResourceList = `-- name: CreateResourceList :exec
INSERT INTO resource_lists (title, created_at, updated_at) VALUES (?, ?, ?)
`

type CreateResourceListParams struct {
	Title     string `json:"title"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (q *Queries) CreateResourceList(ctx context.Context, arg CreateResourceListParams) error {
	_, err := q.db.ExecContext(ctx, createResourceList, arg.Title, arg.CreatedAt, arg.UpdatedAt)
	return err
}

const createResourceReference = `-- name: CreateResourceReference :exec
INSERT INTO resource_references (resource_id, resource_list_id, created_at, updated_at) VALUES (?, ?, ?, ?)
`

type CreateResourceReferenceParams struct {
	ResourceID     string `json:"resource_id"`
	ResourceListID string `json:"resource_list_id"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}

func (q *Queries) CreateResourceReference(ctx context.Context, arg CreateResourceReferenceParams) error {
	_, err := q.db.ExecContext(ctx, createResourceReference,
		arg.ResourceID,
		arg.ResourceListID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteResource = `-- name: DeleteResource :exec
DELETE FROM resources WHERE id = ?
`

func (q *Queries) DeleteResource(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteResource, id)
	return err
}

const getAnnouncement = `-- name: GetAnnouncement :one
SELECT
  r.id,
  r.title,
  r.content_md,
  r.image_url,
  r.resource_type,
  r.resource_list_id,
  r.created_at,
  r.updated_at,
  a.event_list_id,
  a.approved_by_list_id,
  a.visibility,
  a.announce_at,
  a.discord_channel_id,
  a.discord_message_id
FROM resources r
INNER JOIN announcements a ON r.id = a.id
WHERE r.id = ?
`

type GetAnnouncementRow struct {
	ID               string         `json:"id"`
	Title            string         `json:"title"`
	ContentMd        string         `json:"content_md"`
	ImageUrl         sql.NullString `json:"image_url"`
	ResourceType     string         `json:"resource_type"`
	ResourceListID   sql.NullString `json:"resource_list_id"`
	CreatedAt        int64          `json:"created_at"`
	UpdatedAt        int64          `json:"updated_at"`
	EventListID      sql.NullString `json:"event_list_id"`
	ApprovedByListID sql.NullString `json:"approved_by_list_id"`
	Visibility       string         `json:"visibility"`
	AnnounceAt       int64          `json:"announce_at"`
	DiscordChannelID sql.NullString `json:"discord_channel_id"`
	DiscordMessageID sql.NullString `json:"discord_message_id"`
}

func (q *Queries) GetAnnouncement(ctx context.Context, id string) (GetAnnouncementRow, error) {
	row := q.db.QueryRowContext(ctx, getAnnouncement, id)
	var i GetAnnouncementRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ContentMd,
		&i.ImageUrl,
		&i.ResourceType,
		&i.ResourceListID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EventListID,
		&i.ApprovedByListID,
		&i.Visibility,
		&i.AnnounceAt,
		&i.DiscordChannelID,
		&i.DiscordMessageID,
	)
	return i, err
}

const getEvent = `-- name: GetEvent :one
SELECT
  r.id,
  r.title,
  r.content_md,
  r.image_url,
  r.resource_type,
  r.resource_list_id,
  r.created_at,
  r.updated_at,
  e.location,
  e.start_at,
  e.duration_ms,
  e.is_all_day,
  e.host,
  e.visibility
FROM resources r
INNER JOIN events e ON r.id = e.id
WHERE r.id = ?
`

type GetEventRow struct {
	ID             string         `json:"id"`
	Title          string         `json:"title"`
	ContentMd      string         `json:"content_md"`
	ImageUrl       sql.NullString `json:"image_url"`
	ResourceType   string         `json:"resource_type"`
	ResourceListID sql.NullString `json:"resource_list_id"`
	CreatedAt      int64          `json:"created_at"`
	UpdatedAt      int64          `json:"updated_at"`
	Location       string         `json:"location"`
	StartAt        interface{}    `json:"start_at"`
	DurationMs     interface{}    `json:"duration_ms"`
	IsAllDay       bool           `json:"is_all_day"`
	Host           string         `json:"host"`
	Visibility     string         `json:"visibility"`
}

func (q *Queries) GetEvent(ctx context.Context, id string) (GetEventRow, error) {
	row := q.db.QueryRowContext(ctx, getEvent, id)
	var i GetEventRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ContentMd,
		&i.ImageUrl,
		&i.ResourceType,
		&i.ResourceListID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Location,
		&i.StartAt,
		&i.DurationMs,
		&i.IsAllDay,
		&i.Host,
		&i.Visibility,
	)
	return i, err
}

const getResourceList = `-- name: GetResourceList :many
SELECT rr.id, rr.resource_id, rr.resource_list_id, rr.created_at, rr.updated_at
FROM resource_references rr
JOIN resources r ON rr.resource_id = r.id
JOIN resource_lists rl ON rr.resource_list_id = rl.id
WHERE rl.id = ?
ORDER BY rr.index_in_list ASC
`

type GetResourceListRow struct {
	ID             string `json:"id"`
	ResourceID     string `json:"resource_id"`
	ResourceListID string `json:"resource_list_id"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}

func (q *Queries) GetResourceList(ctx context.Context, id string) ([]GetResourceListRow, error) {
	rows, err := q.db.QueryContext(ctx, getResourceList, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetResourceListRow
	for rows.Next() {
		var i GetResourceListRow
		if err := rows.Scan(
			&i.ID,
			&i.ResourceID,
			&i.ResourceListID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
