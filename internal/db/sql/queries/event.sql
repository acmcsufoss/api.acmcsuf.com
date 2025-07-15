-- name: CreateEvent :exec
INSERT INTO
event (
    uuid,
    location,
    start_at,
    end_at,
    is_all_day,
    host
)
VALUES
(?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetEvent :one
SELECT
    uuid,
    location,
    start_at,
    end_at,
    is_all_day,
    host
FROM
    event
WHERE
    uuid = ?;

-- name: UpdateEvent :exec
UPDATE event
SET
    location = COALESCE(sqlc.narg('location'), location),
    start_at = COALESCE(sqlc.narg('start_at'), start_at),
    end_at = COALESCE(sqlc.narg('end_at'), end_at),
    is_all_day = COALESCE(sqlc.narg('is_all_day'), is_all_day),
    host = COALESCE(sqlc.narg('host'), host)
WHERE
    uuid = sqlc.arg('uuid');

-- name: GetEvents :many
SELECT
    uuid,
    location,
    start_at,
    end_at,
    is_all_day,
    host
FROM
    event;

-- name: DeleteEvent :exec
DELETE FROM event
WHERE uuid = ?;
