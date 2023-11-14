package api

import (
	"context"
	"io"
	"time"

	"github.com/google/uuid"
)

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
	CreateEvent(r CreateEventRequest) (*Event, error)

	// Event returns an event resource.
	Event(id string) (*Event, error)

	// UpdateEvent updates an event resource.
	// UpdateEvent(r UpdateEventRequest) (*Event, error)

	// CreateAnnouncement creates a new announcement resource.
	CreateAnnouncement(r CreateAnnouncementRequest) (*Announcement, error)

	// Announcement returns an announcement resource.
	Announcement(id string) (*Announcement, error)

	// ResourceList returns a resource list.
	ResourceList(id string) (*ResourceList, error)

	// AddResource adds a resource to a resource list.
	AddResource(r AddResourceRequest) error

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
