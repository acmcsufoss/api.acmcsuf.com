-- name: CreateResourceList :exec
INSERT INTO resource_lists (title, created_at, updated_at) VALUES (?, ?, ?);

-- name: CreateResource :exec
INSERT INTO resources (id, title, content_md, image_url, resource_type, resource_list_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: DeleteResource :exec
DELETE FROM resources WHERE id = ?;

-- name: CreateResourceReference :exec
INSERT INTO resource_references (resource_id, resource_list_id, created_at, updated_at) VALUES (?, ?, ?, ?);

-- name: CreateEvent :exec
INSERT INTO events (id, location, start_at, duration_ms, is_all_day, host, visibility, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

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