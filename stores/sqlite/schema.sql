-- Language: sqlite

-- Create the 'meta' table.
CREATE TABLE meta (
    v SMALLINT NOT NULL
);

-- Insert data into 'meta' table.
INSERT INTO meta (v)
VALUES (1);

-- Create the 'resource_lists' table.
CREATE TABLE IF NOT EXISTS resource_lists (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL,
    UNIQUE (id)
);

-- Create the 'resources' table.
CREATE TABLE IF NOT EXISTS resources (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    content_md TEXT NOT NULL,
    image_url TEXT,
    resource_type TEXT NOT NULL,
    resource_list_id INTEGER REFERENCES resource_lists(id),
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL,
    UNIQUE (id, resource_list_id)
);

-- Create the 'resource_references' table.
CREATE TABLE IF NOT EXISTS resource_references (
    id INTEGER PRIMARY KEY,
    resource_id TEXT NOT NULL REFERENCES resources(id) ON DELETE CASCADE,
    resource_list_id INTEGER NOT NULL REFERENCES resource_lists(id) ON DELETE CASCADE,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL,
    UNIQUE (resource_id, resource_list_id)
);

-- Create the 'events' table which is a table of event resources.
CREATE TABLE IF NOT EXISTS events (
    id TEXT PRIMARY KEY REFERENCES resources(id) ON DELETE CASCADE,
    location TEXT NOT NULL,
    start_at NUMBER NOT NULL, -- Start time in UTC milliseconds.
    duration_ms NUMBER NOT NULL,
    is_all_day BOOLEAN NOT NULL,
    host TEXT NOT NULL, -- Accepts team ID or plain text.
    visibility TEXT NOT NULL, -- Accepts 'public' or 'private'.
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL,
    UNIQUE (id)
)