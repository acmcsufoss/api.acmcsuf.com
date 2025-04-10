-- name: GetBoard :one
SELECT
    fullName,
    picture,
    discord
FROM
    officers
WHERE
    uuid = ?;

-- name: GetPositions :one
SELECT
    tiers.title,
    tiers.team,
    positions.semester
FROM
    officers
INNER JOIN positions
    ON officers.uuid = positions.oid
INNER JOIN tiers
    ON positions.tier = tiers.tier
WHERE officers.fullName = ?