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
    position.semester,
    tier.title,
    tier.team
FROM
    officer
INNER JOIN position
    ON officer.uuid = position.oid
INNER JOIN tier
    ON position.tier = tier.tier
WHERE officer.full_name = ?
