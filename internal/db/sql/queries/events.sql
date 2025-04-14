-- name: CreateEvent :exec
INSERT INTO
event (
    uuid,
    location,
    start_at,
    end_at,
    is_all_day,
    host
    -- the following doens't exist in schema
    -- visibility
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
