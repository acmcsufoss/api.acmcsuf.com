-- name: CreateWorkshop :exec
INSERT INTO
workshop (
    uuid,
    title,
    team,
    semester,
    start_at,
    link
)
VALUES
(?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetWorkshop :one
SELECT
    title,
    team,
    semester,
    start_at,
    link
FROM
    workshop
WHERE
    uuid = ?;

-- name: UpdateWorkshop :exec
UPDATE workshop
SET
    title = COALESCE(sqlc.narg('title'), title),
    team = COALESCE(sqlc.narg('team'), team),
    semester = COALESCE(sqlc.narg('semester'), semester),
    start_at = COALESCE(sqlc.narg('start_at'), start_at),
    link = COALESCE(sqlc.narg('link'), link)
WHERE
    uuid = sqlc.arg('uuid');

-- name: GetWorkshops :many
SELECT
    uuid,
    title,
    team,
    semester,
    start_at,
    link
FROM
    workshop;

-- name: DeleteWorkshop :exec
DELETE FROM workshop
WHERE uuid = ?;
