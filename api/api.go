package api

import (
	"errors"
	"time"
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

	// ResourceTypeAnnouncement represents an announcement resource.
	ResourceTypeAnnouncement ResourceType = "announcement"
)

// Resource is a base resource struct.
type Resource struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	ContentMd    string `json:"content_md"`
	ImageURL     string `json:"image_url"`
	ResourceType string `json:"resource_type"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

// TODO: Consider sum type for resource list.
// Reference:
// https://zackoverflow.dev/writing/hacking-go-to-give-it-sumtypes/

// ResourceList is a list of resources.
type ResourceList []interface{}

// AddResourceRequest is the input for adding a resource to a resource list.
type AddResourceRequest struct {
	ResourceID     string `json:"resource_id"`
	ResourceListID string `json:"resource_list_id"`
	Index          int64  `json:"index"`
}

// CreateEventRequest is the input for creating a new event.
type CreateEventRequest struct {
	Resource
	Location   string     `json:"location"`
	StartAt    time.Time  `json:"start_at"`
	DurationMs uint64     `json:"duration_ms"`
	IsAllDay   bool       `json:"is_all_day"`
	Host       string     `json:"host"`
	Visibility Visibility `json:"visibility"`
}

// Event is an event resource.
type Event struct {
	Resource

	Location   string     `json:"location"`
	StartAt    time.Time  `json:"start_at"`
	DurationMs uint64     `json:"duration_ms"`
	IsAllDay   bool       `json:"is_all_day"`
	Host       string     `json:"host"`
	Visibility Visibility `json:"visibility"`
}

// CreateAnnouncementRequest is the input for creating a new announcement.
type CreateAnnouncementRequest struct {
	Resource

	EventListID      string     `json:"event_list_id"`
	ApprovedByListID string     `json:"approved_by_list_id"`
	Visibility       Visibility `json:"visibility"`
	AnnounceAt       time.Time  `json:"announce_at"`
	DiscordChannelID string     `json:"discord_channel_id"`
	DiscordMessageID string     `json:"discord_message_id"`
}

// Announcement is an announcement resource.
type Announcement struct {
	Resource

	EventListID      string     `json:"event_list_id"`
	ApprovedByListID string     `json:"approved_by_list_id"`
	Visibility       Visibility `json:"visibility"`
	AnnounceAt       time.Time  `json:"announce_at"`
	DiscordChannelID string     `json:"discord_channel_id"`
	DiscordMessageID string     `json:"discord_message_id"`
}

// NewCreateEventRequest makes a new CreateEventRequest.
func NewCreateEventRequest(
	title, contentMd, imageURL, location string,
	startAt time.Time, durationMs uint64, isAllDay bool,
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
