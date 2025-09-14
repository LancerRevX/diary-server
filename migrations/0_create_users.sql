DROP TABLE IF EXISTS users;

CREATE TABLE users (
    login TEXT PRIMARY KEY,
    password_hash TEXT NOT NULL
);

INSERT INTO
    users
VALUES
    ('nikita', MD5('m5bg8'));

SELECT
    *
FROM
    users;