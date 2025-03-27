-- name: GetBoard :one
SELECT
    id,
    name,
    branch,
    github,
    discord,
    year,
    bio
FROM
    board_member
WHERE
    id = ?;
