-- name: CreateOfficer :exec
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

-- name: CreateTier :exec
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

-- name: CreatePosition :exec
INSERT INTO
position (
    oid,
    semester,
    tier,
    full_name,
    title,
    team
)
VALUES
(?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetOfficer :one
SELECT
    uuid,
    full_name,
    picture,
    github,
    discord
FROM
    officer
WHERE
    uuid = ?;

-- name: UpdateOfficer :exec
UPDATE officer
SET
    full_name = COALESCE(sqlc.narg('full_name'), full_name),
    picture = COALESCE(sqlc.narg('picture'), picture),
    picture = COALESCE(sqlc.narg('picture'), picture),
    github = COALESCE(sqlc.narg('github'), github),
    discord = COALESCE(sqlc.narg('discord'), discord)
WHERE
    uuid = sqlc.arg('uuid');

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
