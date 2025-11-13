-- Language: sqlite

-- Note: Using datatypes the sqlc does not know how to interpret (such as INT instead of INTEGER)
-- will result in sqlc generating interfaces, which is undesired. I recommend skimming the docs
-- if you are unsure about your sqlc: 
-- https://docs.sqlc.dev/en/stable/reference/datatypes.html

-- Create the 'events' table which is a table of event resources.
CREATE TABLE IF NOT EXISTS event (
    uuid TEXT PRIMARY KEY,
    location TEXT NOT NULL,
    start_at INTEGER NOT NULL, -- Start time in UTC milliseconds.
    end_at INTEGER NOT NULL,
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
    discord_message_id TEXT UNIQUE -- Discord message ID. If present, the announcement has been posted.
    --UNIQUE (id)
);

CREATE TABLE IF NOT EXISTS officer (
    uuid VARCHAR(4) PRIMARY KEY,
    full_name VARCHAR(30) NOT NULL,
    picture VARCHAR(37),
    github VARCHAR(64),
    discord VARCHAR(32)
);

CREATE TABLE IF NOT EXISTS tier (
    tier INT PRIMARY KEY,
    title VARCHAR(40),
    t_index INT,
    team VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS position (
    oid VARCHAR(4) NOT NULL,
    semester VARCHAR(3) NOT NULL,
    tier INTEGER NOT NULL,
    full_name VARCHAR(30) NOT NULL,
    title VARCHAR(40),
    team VARCHAR(20),
    PRIMARY KEY (oid, semester, tier),
    CONSTRAINT fk_officer FOREIGN KEY (oid) REFERENCES officer (uuid),
    CONSTRAINT fk_tier FOREIGN KEY (tier) REFERENCES tier (tier)
);

-- TODO: Create a table for access tokens for the API.
