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

-- name: CreatePerson :one
INSERT INTO
person (uuid, name, preferred_pronoun)
VALUES
(?, ?, ?)
RETURNING *;

-- name: CreateAnnouncement :one
INSERT INTO
announcement (
    uuid,
    visibility,
    announce_at,
    discord_channel_id,
    discord_message_id
)
VALUES
(?, ?, ?, ?, ?)
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

-- name: GetPerson :one
SELECT
    uuid,
    name,
    preferred_pronoun
FROM
    person
WHERE
    uuid = ?;

-- name: GetAnnouncement :one
SELECT
    uuid,
    visibility,
    announce_at,
    discord_channel_id,
    discord_message_id
FROM
    announcement
WHERE
    uuid = ?;

-- name: GetBoard :one
SELECT
    id,
    name,
    branch,
    github,
    discord,
    year,
    bio
FROM
    board_member
WHERE
    id = ?;
