-- Language: sqlite
-- Create the 'resource_mapping' table.
CREATE TABLE
    IF NOT EXISTS resource_id_group_id_mapping (
        resource_uuid TEXT REFERENCES resource (uuid),
        group_uuid TEXT NOT NULL REFERENCES group_resource_list_mapping (uuid),
        type TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT NULL,
    );

CREATE TABLE
    IF NOT EXISTS group_id_resource_list_mapping (
        group_uuid TEXT,
        resource_uuid TEXT NOT NULL REFERENCES resource (uuid),
        index_in_list INTEGER NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT NULL,
    );

-- Create the 'resource' table.
CREATE TABLE
    IF NOT EXISTS resource (
        uuid TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        content_md TEXT NOT NULL,
        image_url TEXT,
        resource_type TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT NULL,
    );

-- Create the 'events' table which is a table of event resources.
CREATE TABLE
    IF NOT EXISTS event (
        uuid TEXT PRIMARY KEY REFERENCES resource (uuid),
        location TEXT NOT NULL,
        start_at NUMBER NOT NULL, -- Start time in UTC milliseconds.
        end_at NUMBER NOT NULL,
        is_all_day BOOLEAN NOT NULL,
        host TEXT NOT NULL, -- Accepts team ID or plain text.
        visibility TEXT NOT NULL, -- Accepts 'public' or 'private'.
    )
    -- Create the 'person' table which is a table of person resources.
CREATE TABLE
    IF NOT EXISTS person (
        uuid TEXT REFERENCES resource (uuid),
        name TEXT,
        preferred_pronoun TEXT
    )
    -- Create the 'announcement' table which is a table of announcement resources.
CREATE TABLE
    IF NOT EXISTS announcement (
        uuid TEXT PRIMARY KEY REFERENCES resource (uuid),
        event_groups_group_uuid TEXT REFERENCES resource_group_mapping (resource_uuid),
        approved_by_list_uuid TEXT REFERENCES group_id_resource_list_mapping (uuid),
        visibility TEXT NOT NULL, -- Accepts 'public' or 'private'.
        announce_at INTEGER NOT NULL, -- UTC milliseconds.
        discord_channel_id TEXT, -- Discord channel ID.
        discord_message_id TEXT, -- Discord message ID. If present, the announcement has been posted.
        UNIQUE (id)
    )