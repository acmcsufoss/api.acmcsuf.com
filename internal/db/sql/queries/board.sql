-- name: CreateOfficer :one
INSERT INTO
officer (
    uuid,
    full_name,
    picture,
    github,
    discord
)
VALUES
(?, ?, ?, ?, ?)
RETURNING *;

-- name: CreateTier :one
INSERT INTO
tier (
    tier,
    title,
    t_index,
    team
)
VALUES
(?, ?, ?, ?)
RETURNING *;

-- name: CreatePosition :one
INSERT INTO
position (
    oid,
    semester,
    tier
)
VALUES
(?, ?, ?)
RETURNING *;

-- name: GetOfficer :one
SELECT
    full_name,
    picture,
    github,
    discord
FROM
    officer
WHERE
    uuid = ?;

-- name: GetTier :one
SELECT
    tier,
    title,
    t_index,
    team
FROM
    tier
WHERE
    tier = ?;

-- name: GetPosition :one
SELECT
    oid,
    semester,
    tier,
    full_name,
    title,
    team
FROM
    position
WHERE
    full_name = ?;

-- NOTE: Had to declare above table as :one, may need to change later to :many

-- name: UpdateOfficer :exec
UPDATE officer
SET
    full_name = COALESCE(:full_name, full_name),
    picture = COALESCE(:picture, picture),
    github = COALESCE(:github, github),
    discord = COALESCE(:discord, discord)
WHERE
    uuid = :uuid;

-- name: UpdateTier :exec 
UPDATE tier
SET
    title = COALESCE(sqlc.narg('title'), title),
    t_index = COALESCE(sqlc.narg('t_index'), t_index),
    team = COALESCE(sqlc.narg('team'), team)
WHERE
    tier = sqlc.arg('tier');

-- name: UpdatePosition :exec 
UPDATE position
SET
    full_name = COALESCE(sqlc.narg('full_name'), full_name),
    title = COALESCE(sqlc.narg('title'), title),
    team = COALESCE(sqlc.narg('team'), team)
WHERE
    oid = sqlc.arg('oid')
AND
    semester = sqlc.arg('semester')
AND
    tier = sqlc.arg('tier');

-- name: DeleteOfficer :exec
DELETE FROM officer
WHERE uuid = ?;

-- name: DeleteTier :exec
DELETE FROM tier
WHERE tier = ?;

-- name: DeletePosition :exec
DELETE FROM position
WHERE oid = ?
AND semester = ?
AND tier = ?;
