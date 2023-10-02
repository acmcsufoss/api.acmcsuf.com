-- Language: postgresql

-- Create the 'meta' table
CREATE TABLE meta (
    v SMALLINT NOT NULL
);

-- Insert data into 'meta' table
INSERT INTO meta (v)
VALUES (1);

-- https://github.com/diamondburned/acmregister/blob/main/internal/stores/postgres/queries.sql

-- Get resources of announcement groups of announcement by announcement ID.
-- Insert events into 'events' table
-- Insert announcement groups into 'announcement_groups' table
-- Insert announcements into 'announcements' table
-- Delete announcement groups from 'announcement_groups' table
-- Delete announcements from 'announcements' table
-- Delete events from 'events' table
-- Update events in 'events' table
-- Update announcement groups in 'announcement_groups' table
-- Update announcements in 'announcements' table

-- The order of the resources in the announcement group is important
-- The order of the announcement groups in the announcement is important
-- Many resources can be in many announcement groups
-- The deletion of one announcement group should not delete the resources
-- Events are considered a type of resource.
-- Board teams are considered a type of resource.

CREATE TABLE IF NOT EXISTS announcements (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY,
    is_draft BOOLEAN NOT NULL,
    scheduled_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
);

CREATE TABLE IF NOT EXISTS resource_groups (
    id INTEGER NOT NULL UNIQUE PRIMARY KEY,
    content TEXT,
	attachments TEXT[] NOT NULL,
    discord_message_id BIGINT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS blobs (
  id INTEGER NOT NULL UNIQUE PRIMARY KEY,
  url TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
);

CREATE TABLE IF NOT EXISTS announcement_resource_groups (
	announcement_id INTEGER NOT NULL REFERENCES announcements(id) ON DELETE CASCADE,
	resource_group_id INTEGER NOT NULL REFERENCES resource_groups(id) ON DELETE CASCADE,
	order_in_announcement SMALLINT NOT NULL,
	PRIMARY KEY (announcement_id, resource_group_id)
);

-- CREATE TABLE IF NOT EXISTS  (
-- 	announcement_id INTEGER NOT NULL REFERENCES announcements(id) ON DELETE CASCADE,
-- 	announcement_group_id INTEGER NOT NULL REFERENCES announcement_groups(id) ON DELETE CASCADE,
-- 	cardinality SMALLINT NOT NULL,
-- 	PRIMARY KEY (announcement_id, announcement_group_id)
-- );

-- Create the 'events' table
CREATE TABLE events (
	id INTEGER NOT NULL UNIQUE PRIMARY KEY,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	is_draft BOOLEAN NOT NULL,
	location_type TEXT,
	location_name TEXT,
	is_all_day BOOLEAN NOT NULL,
	start_at TIMESTAMP,
	end_at TIMESTAMP,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
);

-- Create the 'board_team' table
CREATE TABLE board_team (
	id INTEGER NOT NULL UNIQUE PRIMARY KEY,
	name TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

-- Create the 'resources' table
CREATE TABLE resources (
	id INTEGER NOT NULL UNIQUE PRIMARY KEY,
	name TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

CREATE TABLE announcement_group_resources (
	announcement_group_id INTEGER NOT NULL REFERENCES announcement_groups(id) ON DELETE CASCADE,
	resource_id INTEGER NOT NULL REFERENCES resources(id) ON DELETE CASCADE,
	PRIMARY KEY (announcement_group_id, resource_id)
);

-- Create the many-to-many relationship table between announcement groups and announcements
CREATE TABLE announcement_group_announcements (
);

-- Create the relationship between events and resources as events are a type of resource
CREATE TABLE event_resources (
);

-- Future tables: forms, shortlinks, etc.

-- Sample schema

CREATE TABLE
	known_guilds (
		guild_id BIGINT PRIMARY KEY,
		channel_id BIGINT NOT NULL,
		role_id BIGINT NOT NULL,
		init_user_id BIGINT NOT NULL,
		registered_message TEXT NOT NULL
	);

CREATE TABLE
	members (
		guild_id BIGINT NOT NULL REFERENCES known_guilds(guild_id) ON DELETE CASCADE,
		user_id BIGINT NOT NULL,
		email TEXT NOT NULL,
		metadata JSONB NOT NULL,
		UNIQUE (guild_id, user_id),
		UNIQUE (guild_id, email)
	);

CREATE TABLE
	registration_submissions (
		guild_id BIGINT NOT NULL REFERENCES known_guilds(guild_id) ON DELETE CASCADE,
		user_id BIGINT NOT NULL,
		metadata JSONB NOT NULL,
		expire_at TIMESTAMP NOT NULL,
		UNIQUE(guild_id, user_id)
	);

CREATE TABLE
	pin_codes (
		guild_id BIGINT NOT NULL REFERENCES known_guilds(guild_id) ON DELETE CASCADE,
		user_id BIGINT NOT NULL,
		pin SMALLINT NOT NULL,
		UNIQUE(guild_id, user_id),
		UNIQUE(guild_id, pin),
		FOREIGN KEY (guild_id, user_id) REFERENCES registration_submissions(guild_id, user_id) ON DELETE CASCADE
	);