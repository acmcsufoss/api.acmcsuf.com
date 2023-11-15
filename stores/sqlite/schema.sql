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
    id TEXT PRIMARY KEY,
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
    resource_list_id TEXT REFERENCES resource_lists(id), -- Related resources are added to this list.
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL,
    UNIQUE (id, resource_list_id)
);

-- Create the 'resource_references' table.
CREATE TABLE IF NOT EXISTS resource_references (
    id TEXT PRIMARY KEY,
    resource_id TEXT NOT NULL REFERENCES resources(id) ON DELETE CASCADE,
    resource_list_id TEXT NOT NULL REFERENCES resource_lists(id) ON DELETE CASCADE,
    index_in_list INTEGER NOT NULL, -- The index of the resource in the list.
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
    UNIQUE (id)
)

-- Create the 'announcements' table which is a table of announcement resources.
CREATE TABLE IF NOT EXISTS announcements (
    id TEXT PRIMARY KEY REFERENCES resources(id) ON DELETE CASCADE,
    event_list_id TEXT REFERENCES resource_lists(id) ON DELETE CASCADE,
    approved_by_list_id TEXT REFERENCES resource_lists(id) ON DELETE CASCADE,
    visibility TEXT NOT NULL, -- Accepts 'public' or 'private'.
    announce_at INTEGER NOT NULL, -- UTC milliseconds.
    discord_channel_id TEXT, -- Discord channel ID. If present, the announcement has been posted.
    discord_message_id TEXT, -- Discord message ID. If present, the announcement has been posted.
    UNIQUE (id)
)

-- Create the 'person' table
CREATE TABLE IF NOT EXISTS person (
    id TEXT PRIMARY KEY REFERENCES resources(id) ON DELETE CASCADE,
    display_name TEXT NOT NULL,
    full_name TEXT NOT NULL,
    picture_url TEXT, -- URL of their profile picture, if applicable
    team_id INTEGER REFERENCES team(id),
    date_joined INTEGER NOT NULL,
    UNIQUE (id)
);

-- Create the 'team' table
CREATE TABLE IF NOT EXISTS team (
    id INTEGER PRIMARY KEY REFERENCES resources(id) ON DELETE CASCADE,         
    name TEXT NOT NULL, -- Name of the team
    date_created INTEGER NOT NULL, -- UTC milliseconds.
    UNIQUE (id)
);

-- Create the 'blog_post' table
-- Reference used:
-- https://github.com/EthanThatOneKid/acmcsuf.com/blob/main/src/lib/public/blog/types.ts
CREATE TABLE IF NOT EXISTS blog_post (
    id INTEGER PRIMARY KEY REFERENCES resources(resource_id) ON DELETE CASCADE,     
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    author TEXT REFERENCES person(name),
    author_id INTEGER REFERENCES person(person_id),
    author_display_name TEXT REFERENCES person(full_name),
    date_created INTEGER NOT NULL, -- UTC milliseconds.
    last_edited INTEGER, -- UTC milliseconds.
    tags TEXT, -- Array of tags for the blog post separated by a delimiter
    html TEXT NOT NULL, -- Blog post content in HTML form 
    UNIQUE (id, author_id)
);
