package models

import "database/sql"

type Announcement struct {
	Uuid                 string         `json:"uuid"`
	EventGroupsGroupUuid sql.NullString `json:"event_groups_group_uuid"`
	ApprovedByListUuid   sql.NullString `json:"approved_by_list_uuid"`
	Visibility           string         `json:"visibility"`
	AnnounceAt           int64          `json:"announce_at"`
	DiscordChannelID     sql.NullString `json:"discord_channel_id"`
	DiscordMessageID     sql.NullString `json:"discord_message_id"`
}

type Event struct {
	Uuid       string      `json:"uuid"`
	Location   string      `json:"location"`
	StartAt    interface{} `json:"start_at"`
	EndAt      interface{} `json:"end_at"`
	IsAllDay   bool        `json:"is_all_day"`
	Host       string      `json:"host"`
	Visibility string      `json:"visibility"`
}

type GroupIDResourceListMapping struct {
	GroupUuid    sql.NullString `json:"group_uuid"`
	ResourceUuid string         `json:"resource_uuid"`
	IndexInList  int64          `json:"index_in_list"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
	DeletedAt    sql.NullTime   `json:"deleted_at"`
}

type Person struct {
	Uuid             sql.NullString `json:"uuid"`
	Name             sql.NullString `json:"name"`
	PreferredPronoun sql.NullString `json:"preferred_pronoun"`
}

type Resource struct {
	Uuid         string         `json:"uuid"`
	Title        string         `json:"title"`
	ContentMd    string         `json:"content_md"`
	ImageUrl     sql.NullString `json:"image_url"`
	ResourceType string         `json:"resource_type"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
	DeletedAt    sql.NullTime   `json:"deleted_at"`
}

type ResourceIDGroupIDMapping struct {
	ResourceUuid sql.NullString `json:"resource_uuid"`
	GroupUuid    string         `json:"group_uuid"`
	Type         sql.NullString `json:"type"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
	DeletedAt    sql.NullTime   `json:"deleted_at"`
}
