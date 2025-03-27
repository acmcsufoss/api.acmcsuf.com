-- name: CreatePerson :one
INSERT INTO
person (uuid, name, preferred_pronoun)
VALUES
(?, ?, ?)
RETURNING *;

-- name: GetPerson :one
SELECT
    uuid,
    name,
    preferred_pronoun
FROM
    person
WHERE
    uuid = ?;
