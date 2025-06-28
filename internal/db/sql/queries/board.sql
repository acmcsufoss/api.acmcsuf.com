-- name: CreateOfficer :exec
INSERT INTO
officers (
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
tiers (
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
positions (
    oid,
    semester,
    tier
)
VALUES
(?, ?, ?)
RETURNING *;

-- name: GetOfficer :one
SELECT
    uuid,
    full_name,
    picture,
    github,
    discord
FROM
    officers
WHERE
    uuid = ?;

-- name: GetTier :one
SELECT
    tier,
    title,
    t_index,
    team
FROM
    tiers
WHERE
    tier = ?;

-- name: GetPosition :many
SELECT
    officers.full_name,
    positions.semester,
    tiers.title,
    tiers.team
FROM
    officers
INNER JOIN positions
    ON officers.uuid = positions.oid
INNER JOIN tiers
    ON positions.tier = tiers.tier
WHERE officers.full_name = ?;
