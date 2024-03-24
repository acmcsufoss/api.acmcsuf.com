-- name: CreateResource :exec
INSERT INTO
    resource (
        uuid,
        title,
        content_md,
        image_url,
        resource_type,
        created_at,
        updated_at,
        deleted_at
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?);

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

-- name: CreateResourceGroupMapping :exec
INSERT INTO
    resource_id_group_id_mapping (
        resource_uuid,
        group_uuid,
        type,
        created_at,
        updated_at,
        deleted_at
    )
VALUES
    (?, ?, ?, ?, ?, ?);

-- name: CreateGroupResourceMapping :exec
INSERT INTO
    group_id_resource_list_mapping (
        group_uuid,
        resource_uuid,
        index_in_list,
        created_at,
        updated_at,
        deleted_at
    )
VALUES
    (?, ?, ?, ?, ?, ?);

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

-- name: DeleteResource :exec
DELETE FROM resource
WHERE
    id = ?;

-- name: GetResource :exec
SELECT
    uuid,
    title,
    content_md,
    image_url,
    resource_type,
    created_at,
    updated_at,
    deleted_at
from
    resource
where
    uuid = ?;

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

-- name: GetResourceGroupMapping :exec
SELECT
    resource_uuid,
    group_uuid,
    type,
    created_at,
    updated_at,
    deleted_at
from
    resource_id_group_id_mapping
where
    resource_uuid = ?;

-- name: GetGroupResourceMapping :exec
SELECT
    group_uuid,
    resource_uuid,
    index_in_list,
    created_at,
    updated_at,
    deleted_at
from
    group_id_resource_list_mapping
where
    group_uuid = ?;

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
    announcements
where
    uuid = ?;