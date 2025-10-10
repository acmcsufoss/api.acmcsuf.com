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

-- name: UpdateTier :exec
UPDATE tier
SET
    title = COALESCE(:title, title),
    t_index = COALESCE(:t_index, t_index),
    team = COALESCE(:team, team)
WHERE
    tier = :tier;

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

-- name: DeleteTier :exec
DELETE FROM tier
WHERE tier = ?;

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

-- name: UpdatePosition
UPDATE position
SET
    full_name = COALESCE(:full_name, full_name),
    title = COALESCE(:title, title),
    team = COALESCE(:team, team)
WHERE
    oid = :oid
    AND semester = :semester
    AND tier = :tier;

-- name: DeletePosition :exec
DELETE FROM position
WHERE
    oid = ?
    AND semester = ?
    AND tier = ?;
