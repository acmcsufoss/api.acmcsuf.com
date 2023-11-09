package stores

import (
	"context"
	"database/sql"
	"time"

	"github.com/acmcsufoss/api.acmcsuf.com/api"
	"github.com/acmcsufoss/api.acmcsuf.com/stores/sqlite"
	"github.com/pkg/errors"
)

type sqliteStore struct {
	q   *sqlite.Queries
	db  *sql.DB
	ctx context.Context
}

func (s sqliteStore) Close() error {
	return s.db.Close()
}

func (s sqliteStore) WithContext(ctx context.Context) api.ContainsContext {
	s.ctx = ctx
	return s
}

// NewSQLite creates a new SQLite store.
func NewSQLite(ctx context.Context, uri string) (StoreCloser, error) {
	db, err := sql.Open("sqlite", uri)
	if err != nil {
		return nil, errors.Wrap(err, "sql/sqlite3")
	}

	if err := sqlite.Migrate(ctx, db); err != nil {
		return nil, errors.Wrap(err, "cannot migrate sqlite db")
	}

	return sqliteStore{
		q:   sqlite.New(db),
		db:  db,
		ctx: ctx,
	}, nil
}

func (s sqliteStore) CreateResource(r api.Resource) (*api.Resource, error) {
	resourceID, now := api.NewID(), api.Now()
	resource := &api.Resource{
		ID:           resourceID,
		Title:        r.Title,
		ContentMd:    r.ContentMd,
		ImageURL:     r.ImageURL,
		ResourceType: r.ResourceType,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	if err := s.q.CreateResource(s.ctx, sqlite.CreateResourceParams{
		ID:           resource.ID,
		Title:        resource.Title,
		ContentMd:    resource.ContentMd,
		ImageUrl:     sql.NullString{String: resource.ImageURL, Valid: resource.ImageURL != ""},
		ResourceType: resource.ResourceType,
		CreatedAt:    now,
		UpdatedAt:    now,
	}); err != nil {
		if sqlite.IsConstraintFailed(err) {
			return nil, api.ErrEventIDConflict
		}

		return nil, sqliteErr(err)
	}

	return resource, nil
}

func (s sqliteStore) CreateEvent(r api.CreateEventRequest) (*api.Event, error) {
	resource, err := s.CreateResource(r.Resource)
	if err != nil {
		return nil, err
	}

	event := sqlite.Event{
		ID:         resource.ID,
		Location:   r.Location,
		StartAt:    r.StartAt,
		DurationMs: r.DurationMs,
		IsAllDay:   r.IsAllDay,
		Host:       r.Host,
		Visibility: string(r.Visibility),
	}
	if err := s.q.CreateEvent(s.ctx, sqlite.CreateEventParams(event)); err != nil {
		if sqlite.IsConstraintFailed(err) {
			return nil, api.ErrEventIDConflict
		}

		return nil, sqliteErr(err)
	}

	return &api.Event{
		Resource:   *resource,
		Location:   r.Location,
		StartAt:    r.StartAt,
		DurationMs: r.DurationMs,
		IsAllDay:   r.IsAllDay,
		Host:       r.Host,
		Visibility: r.Visibility,
	}, nil
}

func (s sqliteStore) Event(id string) (*api.Event, error) {
	eventRow, err := s.q.GetEvent(s.ctx, id)
	if err != nil {
		return nil, sqliteErr(err)
	}

	var startAt time.Time
	switch v := eventRow.StartAt.(type) {
	default:
		return nil, errors.New("unknown type")
	case int64:
		startAt = time.Unix(v/1000, 0)
	}

	var durationMs uint64
	switch v := eventRow.DurationMs.(type) {
	default:
		return nil, errors.New("unknown type")
	case int64:
		durationMs = uint64(v)
	}

	return &api.Event{
		Resource: api.Resource{
			ID:           eventRow.ID,
			Title:        eventRow.Title,
			ContentMd:    eventRow.ContentMd,
			ImageURL:     eventRow.ImageUrl.String,
			ResourceType: string(api.ResourceTypeEvent),
			CreatedAt:    eventRow.CreatedAt,
			UpdatedAt:    eventRow.UpdatedAt,
		},
		Location:   eventRow.Location,
		StartAt:    startAt,
		DurationMs: durationMs,
		IsAllDay:   eventRow.IsAllDay,
		Host:       eventRow.Host,
		Visibility: api.Visibility(eventRow.Visibility),
	}, nil
}

func (s sqliteStore) CreateAnnouncement(r api.CreateAnnouncementRequest) (*api.Announcement, error) {
	resource, err := s.CreateResource(r.Resource)
	if err != nil {
		return nil, err
	}

	announcement := sqlite.Announcement{
		ID:               resource.ID,
		EventListID:      sql.NullString{String: r.EventListID, Valid: r.EventListID != ""},
		ApprovedByListID: sql.NullString{String: r.ApprovedByListID, Valid: r.ApprovedByListID != ""},
		Visibility:       string(r.Visibility),
		AnnounceAt:       r.AnnounceAt.Unix() * 1000,
		DiscordChannelID: sql.NullString{String: r.DiscordChannelID, Valid: r.DiscordChannelID != ""},
		DiscordMessageID: sql.NullString{String: r.DiscordMessageID, Valid: r.DiscordMessageID != ""},
	}
	if err := s.q.CreateAnnouncement(s.ctx, sqlite.CreateAnnouncementParams(announcement)); err != nil {
		if sqlite.IsConstraintFailed(err) {
			return nil, api.ErrEventIDConflict
		}

		return nil, sqliteErr(err)
	}

	return &api.Announcement{
		Resource:         *resource,
		EventListID:      r.EventListID,
		ApprovedByListID: r.ApprovedByListID,
		Visibility:       r.Visibility,
		AnnounceAt:       r.AnnounceAt,
		DiscordChannelID: r.DiscordChannelID,
		DiscordMessageID: r.DiscordMessageID,
	}, nil
}

func (s sqliteStore) Announcement(id string) (*api.Announcement, error) {
	announcementRow, err := s.q.GetAnnouncement(s.ctx, id)
	if err != nil {
		return nil, sqliteErr(err)
	}

	return &api.Announcement{
		Resource: api.Resource{
			ID:           announcementRow.ID,
			Title:        announcementRow.Title,
			ContentMd:    announcementRow.ContentMd,
			ImageURL:     announcementRow.ImageUrl.String,
			ResourceType: string(api.ResourceTypeAnnouncement),
			CreatedAt:    announcementRow.CreatedAt,
			UpdatedAt:    announcementRow.UpdatedAt,
		},
		EventListID:      announcementRow.EventListID.String,
		ApprovedByListID: announcementRow.ApprovedByListID.String,
		Visibility:       api.Visibility(announcementRow.Visibility),
		AnnounceAt:       time.Unix(announcementRow.AnnounceAt/1000, 0),
		DiscordChannelID: announcementRow.DiscordChannelID.String,
		DiscordMessageID: announcementRow.DiscordMessageID.String,
	}, nil
}

func (s sqliteStore) ResourceList(id string) (*api.ResourceList, error) {
	resourceRows, err := s.q.GetResourceList(s.ctx, id)
	if err != nil {
		return nil, sqliteErr(err)
	}

	resourceList := make(api.ResourceList, len(resourceRows))
	for i, resourceRow := range resourceRows {
		resourceList[i] = resourceRow
	}

	return &resourceList, nil
}

func (s sqliteStore) DeleteResource(id string) error {
	if err := s.q.DeleteResource(s.ctx, id); err != nil {
		return sqliteErr(err)
	}

	return nil
}

func sqliteErr(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return api.ErrNotFound
	}
	return err
}
