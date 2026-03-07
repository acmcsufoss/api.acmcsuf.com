-- name: CreateWorkshop :exec
INSERT INTO
workshop (
    uuid,
    title,
    team,
    semester,
    start_at,
    link
)
VALUES
(?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteWorkshop :exec
DELETE FROM workshop
WHERE uuid = ?;
