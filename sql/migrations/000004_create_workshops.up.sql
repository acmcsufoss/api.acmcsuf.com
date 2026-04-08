-- Language: sqlite

CREATE TABLE IF NOT EXISTS workshop (
    uuid INT PRIMARY KEY,
    title VARCHAR(100),
    team VARCHAR(20),
    semester CHAR(3),
    start_at DATE,
    link TEXT
);
