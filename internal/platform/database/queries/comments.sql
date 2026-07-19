-- name: CreateComment :one
INSERT INTO comments (
    task_id,
    author_id,
    body
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;


-- name: GetComment :one
SELECT *
FROM comments
WHERE id = $1
LIMIT 1;


-- name: ListTaskComments :many
SELECT *
FROM comments
WHERE task_id = $1
ORDER BY created_at ASC;


-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;