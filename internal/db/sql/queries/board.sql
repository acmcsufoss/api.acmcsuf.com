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

-- name: GetPosition :one
SELECT
    oid,
    semester,
    tier,
    full_name,
    title,
    team
FROM
    positions
WHERE 
    full_name = ?;
