-- Language: sqlite

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
    officer_id VARCHAR(4) NOT NULL,
    semester VARCHAR(3) NOT NULL,
    tier INTEGER NOT NULL,
    full_name VARCHAR(30) NOT NULL,
    title VARCHAR(40),
    team VARCHAR(20),
    PRIMARY KEY (officer_id, semester, tier),
    CONSTRAINT fk_officer FOREIGN KEY (officer_id) REFERENCES officer (uuid),
    CONSTRAINT fk_tier FOREIGN KEY (tier) REFERENCES tier (tier)
);
