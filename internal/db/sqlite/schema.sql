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
    discord_message_id TEXT UNIQUE -- Discord message ID. If present, the announcement has been posted.
    --UNIQUE (id)
);

CREATE TABLE IF NOT EXISTS boardMember (
    id CHAR(4) PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    branch VARCHAR(20) NOT NULL,
    github VARCHAR(39),
    discord VARCHAR(32),
    year INT,
    bio TEXT
);

CREATE TABLE IF NOT EXISTS branch (
    name VARCHAR(20) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS memberOf (
    bmid CHAR(4) NOT NULL,
    bname VARCHAR(20) NOT NULL,
    PRIMARY KEY (bmid, bname),

    CONSTRAINT fk_board FOREIGN KEY (bmid) REFERENCES boardMember(id),
    CONSTRAINT fk_branch FOREIGN KEY (bname) REFERENCES branch(name)
);

-- TODO: Create a table for access tokens for the API.
