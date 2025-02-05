-- name: CreateEvent :exec
INSERT INTO
    event (
        uuid,
        location,
        start_at,
        end_at,
        is_all_day,
        host,
        visibility
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?);

-- name: CreatePerson :exec
INSERT INTO
    person (uuid, name, preferred_pronoun)
VALUES
    (?, ?, ?);

-- name: CreateAnnouncement :exec
INSERT INTO
    announcement (
        uuid,
        event_groups_group_uuid,
        approved_by_list_uuid,
        visibility,
        announce_at,
        discord_channel_id,
        discord_message_id
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?);

-- name: GetEvent :exec
SELECT
    uuid,
    location,
    start_at,
    end_at,
    is_all_day,
    host,
    visibility
from
    event
where
    uuid = ?;

-- name: GetPerson :exec
SELECT
    uuid,
    name,
    preferred_pronoun
from
    person
where
    uuid = ?;

-- name: GetAnnouncement :exec
SELECT
    uuid,
    event_groups_group_uuid,
    approved_by_list_uuid,
    visibility,
    announce_at,
    discord_channel_id,
    discord_message_id
from
    announcement
where
    uuid = ?;