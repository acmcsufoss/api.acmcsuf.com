package stores

import (
	"context"
	"database/sql"

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

func (s sqliteStore) CreateResource(r api.Resource) (*string, *int64, error) {
	resourceID, now := api.NewID(), api.Now()
	if err := s.q.CreateResource(s.ctx, sqlite.CreateResourceParams{
		ID:           resourceID,
		Title:        r.Title,
		ContentMd:    r.ContentMd,
		ImageUrl:     sql.NullString{String: r.ImageURL, Valid: r.ImageURL != ""},
		ResourceType: r.ResourceType,
		CreatedAt:    now,
		UpdatedAt:    now,
	}); err != nil {
		if sqlite.IsConstraintFailed(err) {
			return nil, nil, api.ErrEventIDConflict
		}

		return nil, nil, sqliteErr(err)
	}

	return &resourceID, &now, nil
}

func (s sqliteStore) CreateEvent(r api.CreateEventRequest) (*sqlite.Event, error) {
	resourceID, now, err := s.CreateResource(r.Resource)
	if err != nil {
		return nil, err
	}

	event := sqlite.Event{
		ID:         *resourceID,
		Location:   r.Location,
		StartAt:    r.StartAt,
		DurationMs: r.DurationMs,
		IsAllDay:   r.IsAllDay,
		Host:       r.Host,
		Visibility: string(r.Visibility),
		CreatedAt:  *now,
		UpdatedAt:  *now,
	}
	if err := s.q.CreateEvent(s.ctx, sqlite.CreateEventParams(event)); err != nil {
		if sqlite.IsConstraintFailed(err) {
			return nil, api.ErrEventIDConflict
		}

		return nil, sqliteErr(err)
	}

	return &event, nil
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
