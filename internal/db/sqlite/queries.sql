-- name: CreateResource :exec
INSERT INTO resources (uuid, title, content_md, image_url, resource_type, created_at, updated_at,deleted_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: CreateEvent :exec
INSERT INTO event (uuid, location, start_at, end_at, is_all_day, host, visibility) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: CreatePerson :exec
INSERT INTO person (uuid, name, preferred_pronoun) VALUES (?, ?, ?);

-- name: CreateResourceGroupMapping :exec
INSERT INTO resource_id_group_id_mapping (resource_uuid, group_uuid, type, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?);

-- name: CreateGroupResourceMapping :exec
INSERT INTO group_id_resource_list_mapping ( group_uuid, resource_uuid, index_in_list, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?);

-- name: CreateAnnouncement :exec
INSERT INTO announcements (uuid, event_groups_group_uuid, approved_by_list_uuid, visibility, announce_at, discord_channel_id, discord_message_id) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: DeleteResource :exec
DELETE FROM resources WHERE id = ?;

