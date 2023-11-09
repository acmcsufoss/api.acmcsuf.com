package api

import (
	"context"
	"errors"
	"io"
	"time"

	// TODO: Remove stores domain from API package.
	"github.com/acmcsufoss/api.acmcsuf.com/stores/sqlite"
	"github.com/google/uuid"
)

// ErrNotFound is returned if anything is not found.
var ErrNotFound = errors.New("not found")

// ErrEventIDConflict is returned if an event with the same ID already exists.
var ErrEventIDConflict = errors.New("event ID conflict")

// Visibility is a visibility level.
type Visibility string

const (
	// VisibilityPublic represents a public visibility level.
	VisibilityPublic Visibility = "public"

	// VisibilityPrivate represents a private visibility level.
	VisibilityPrivate Visibility = "private"
)

// ResourceType is a resource type.
type ResourceType string

const (
	// ResourceTypeEvent represents an event resource.
	ResourceTypeEvent ResourceType = "event"
)

// Resource is a base resource struct.
type Resource struct {
	Title        string `json:"title"`
	ContentMd    string `json:"content_md"`
	ImageURL     string `json:"image_url"`
	ResourceType string `json:"resource_type"`
}

// CreateEventRequest is the input for creating a new event.
type CreateEventRequest struct {
	Resource
	Location   string     `json:"location"`
	StartAt    time.Time  `json:"start_at"`
	DurationMs int64      `json:"duration_ms"`
	IsAllDay   bool       `json:"is_all_day"`
	Host       string     `json:"host"`
	Visibility Visibility `json:"visibility"`
}

// NewCreateEventRequest makes a new CreateEventRequest.
func NewCreateEventRequest(
	title, contentMd, imageURL, location string,
	startAt time.Time, durationMs int64, isAllDay bool,
	host string, visibility Visibility,
) CreateEventRequest {
	return CreateEventRequest{
		Resource: Resource{Title: title,
			ContentMd:    contentMd,
			ImageURL:     imageURL,
			ResourceType: string(ResourceTypeEvent)},
		Location:   location,
		StartAt:    startAt,
		DurationMs: durationMs,
		IsAllDay:   isAllDay,
		Host:       host,
		Visibility: visibility,
	}
}

// ContainsContext can be embedded by any interface to have an overrideable
// context.
type ContainsContext interface {
	WithContext(context.Context) ContainsContext
}

// Store describes a Store instance. It combines all smaller stores.
type Store interface {
	io.Closer
	ContainsContext

	// CreateEvent creates a new event resource.
	CreateEvent(event CreateEventRequest) (*sqlite.Event, error)

	// Event returns an event resource.
	Event(id string) (*sqlite.GetEventRow, error)

	// DeleteResource deletes a resource.
	DeleteResource(id string) error
}

// NewID generates a new resource ID.
func NewID() string {
	return uuid.New().String()
}

// Now generates a timestamp in milliseconds.
func Now() int64 {
	return time.Now().Unix() * 1000
}
