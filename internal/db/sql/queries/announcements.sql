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

-- name: UpdateAnnouncement :exec
UPDATE announcement
SET
    visibility = COALESCE(sqlc.narg('visibility'), visibility),
    announce_at = COALESCE(sqlc.narg('announce_at'), announce_at),
    discord_channel_id = COALESCE(sqlc.narg('discord_channel_id'), discord_channel_id),
    discord_message_id = COALESCE(sqlc.narg('discord_message_id'), discord_message_id)

WHERE
    uuid = sqlc.arg('uuid');

-- name: GetAnnouncements :many
SELECT
    uuid,
    visibility,
    announce_at,
    discord_channel_id,
    discord_message_id
FROM
    announcement;

-- name: DeleteAnnouncement :exec
DELETE FROM announcement
where uuid = ?;
