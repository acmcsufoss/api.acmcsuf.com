-- Language: sqlite

-- Create the 'events' table which is a table of event resources.
CREATE TABLE IF NOT EXISTS event (
    uuid TEXT PRIMARY KEY,
    location TEXT NOT NULL,
    start_at NUMBER NOT NULL, -- Start time in UTC milliseconds.
    end_at NUMBER NOT NULL,
    is_all_day BOOLEAN NOT NULL,
    host TEXT NOT NULL -- Accepts team ID or plain text.
);

-- Create the 'person' table which is a table of person resources.
CREATE TABLE IF NOT EXISTS person (
    uuid TEXT PRIMARY KEY,
    name TEXT,
    preferred_pronoun TEXT
);

-- Create the 'announcement' table which is a table of announcement resources.
CREATE TABLE IF NOT EXISTS announcement (
    uuid TEXT PRIMARY KEY,
    visibility TEXT NOT NULL, -- Accepts 'public' or 'private'.
    announce_at INTEGER NOT NULL, -- UTC milliseconds.
    discord_channel_id TEXT, -- Discord channel ID.
    discord_message_id TEXT UNIQUE, -- Discord message ID. If present, the announcement has been posted.
    --UNIQUE (id)
);

-- TODO: Create a table for access tokens for the API.
