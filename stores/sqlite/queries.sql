-- name: CreateAnnouncement :exec
INSERT INTO
	announcements (id, is_draft, scheduled_at, created_at, updated_at)
VALUES
	($1, $2, $3, $4);

-- name: CreateResourceGroup :exec
INSERT INTO
    resource_groups (id, content, attachments, discord_message_id, created_at, updated_at, index_in_announcement)

-- name: ListResourceGroupsByAnnouncementID :many
SELECT
    id, content, attachments, discord_message_id, created_at, updated_at, index_in_announcement