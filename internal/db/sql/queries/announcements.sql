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
