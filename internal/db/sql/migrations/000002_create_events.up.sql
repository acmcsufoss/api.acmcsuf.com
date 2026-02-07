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
