-- Language: sqlite

-- Create the 'announcement' table which is a table of announcement resources.
CREATE TABLE IF NOT EXISTS announcement (
    uuid TEXT PRIMARY KEY,
    visibility TEXT NOT NULL, -- Accepts 'public' or 'private'.
    announce_at INTEGER NOT NULL, -- UTC milliseconds.
    discord_channel_id TEXT, -- Discord channel ID.
    discord_message_id TEXT UNIQUE -- Discord message ID. If present, the announcement has been posted.
    --UNIQUE (id)
);
