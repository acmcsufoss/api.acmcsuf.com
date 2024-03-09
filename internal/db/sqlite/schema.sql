-- Language: sqlite

-- Create the 'resource_mappings' table.
CREATE TABLE IF NOT EXISTS resource_id_group_id_mapping (
    resource_uuid TEXT REFERENCES resource(uuid),
    group_uuid TEXT NOT NULL REFERENCES group_resource_list_mapping(uuid),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
);

CREATE TABLE IF NOT EXISTS group_id_resource_list_mapping (
    group_uuid TEXT,
    resource_uuid TEXT NOT NULL REFERENCES resource(uuid),
    index_in_list INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
);

-- Create the 'resources' table.
CREATE TABLE IF NOT EXISTS resource (
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
CREATE TABLE IF NOT EXISTS event (
    uuid TEXT PRIMARY KEY REFERENCES resource(uuid),
    location TEXT NOT NULL,
    start_at NUMBER NOT NULL, -- Start time in UTC milliseconds.
    end_at NUMBER NOT NULL,
    is_all_day BOOLEAN NOT NULL,
    host TEXT NOT NULL, -- Accepts team ID or plain text.
    visibility TEXT NOT NULL, -- Accepts 'public' or 'private'.
)

-- Create the 'announcements' table which is a table of announcement resources.
CREATE TABLE IF NOT EXISTS announcement (
    uuid TEXT PRIMARY KEY REFERENCES resource(uuid),
    -- no it is not resource list by uuid because one uuid may have more than one list.
    -- in this case, we have 2. One has people, one has events.
    event_groups_group_uuid TEXT REFERENCES resource_group_mapping(resource_uuid),  --[1,2,3] , [4,5,6] , [6,7,8]
    approved_by_list_uuid TEXT REFERENCES resource_person_mapping(uuid), -- [1:user_1,1:user_2,1:user_3]
    visibility TEXT NOT NULL, -- Accepts 'public' or 'private'.
    announce_at INTEGER NOT NULL, -- UTC milliseconds.
    discord_channel_id TEXT, -- Discord channel ID. If present, the announcement has been posted.
    discord_message_id TEXT, -- Discord message ID. If present, the announcement has been posted.
    UNIQUE (id)
)

-- Create the 'persons' table which is a table of person resources.
CREATE TABLE IF NOT EXISTS person (
    uuid TEXT REFERENCES resource(uuid),
    person_uuid TEXT REFERENCES event(uuid),
)
