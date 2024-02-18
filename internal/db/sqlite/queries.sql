-- name: CreateResource :exec
INSERT INTO resources (id, title, content_md, image_url, resource_type, resource_list_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: CreateResourceList :exec
INSERT INTO resource_lists (title, created_at, updated_at) VALUES (?, ?, ?);

-- name: CreateResourceReference :exec
INSERT INTO resource_references (resource_id, resource_list_id, created_at, updated_at) VALUES (?, ?, ?, ?);

-- name: GetResourceList :many
SELECT rr.resource_id, rr.resource_list_id, rr.created_at, rr.updated_at
FROM resource_references rr
JOIN resources r ON rr.resource_id = r.id
JOIN resource_lists rl ON rr.resource_list_id = rl.id
WHERE rl.id = ?
ORDER BY rr.index_in_list ASC;

-- name: AddResource :exec
INSERT INTO resource_references (resource_id, resource_list_id, index_in_list, created_at, updated_at) VALUES (?, ?, ?, ?, ?);

-- name: DeleteResource :exec
DELETE FROM resources WHERE id = ?;

-- name: CreateEvent :exec
INSERT INTO events (id, location, start_at, duration_ms, is_all_day, host, visibility) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetEvent :one
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
WHERE r.id = ?;

-- name: CreateAnnouncement :exec
INSERT INTO announcements (id, event_list_id, approved_by_list_id, visibility, announce_at, discord_channel_id, discord_message_id) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetAnnouncement :one
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
WHERE r.id = ?;
