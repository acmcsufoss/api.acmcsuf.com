package models

import "time"

type Resource struct {
	UUID         string
	Title        string
	ContentMD    string
	ImageURL     string
	ResourceType string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
