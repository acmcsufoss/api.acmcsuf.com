-- name: CreateOfficer :exec
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
