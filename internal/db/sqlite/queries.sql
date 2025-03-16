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
(?, ?, ?, ?, ?, ?);

-- name: CreatePerson :exec
INSERT INTO
person (uuid, name, preferred_pronoun)
VALUES
(?, ?, ?);

-- name: CreateAnnouncement :exec
INSERT INTO
announcement (
    uuid,
    visibility,
    announce_at,
    discord_channel_id,
    discord_message_id
)
VALUES
(?, ?, ?, ?, ?);

-- name: GetEvent :exec
SELECT
    uuid,
    location,
    start_at,
    end_at,
    is_all_day,
    host
    -- the following does not exist in schema
    -- visibility
FROM
    event
WHERE
    uuid = ?;

-- name: GetPerson :exec
SELECT
    uuid,
    name,
    preferred_pronoun
FROM
    person
WHERE
    uuid = ?;

-- name: GetAnnouncement :exec
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


-- name: GetBoard :exec
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
    id = ?