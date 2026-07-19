-- name: CreateTask :one
INSERT INTO tasks (
    project_id,
    creator_id,
    assignee_id,
    title,
    description,
    status,
    priority,
    due_date
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING
    id,
    project_id,
    creator_id,
    assignee_id,
    title,
    description,
    status,
    priority,
    due_date,
    created_at,
    updated_at;


-- name: GetTaskByID :one
SELECT
    id,
    project_id,
    creator_id,
    assignee_id,
    title,
    description,
    status,
    priority,
    due_date,
    created_at,
    updated_at
FROM tasks
WHERE id = $1
LIMIT 1;


-- name: ListTasksByProject :many
SELECT
    id,
    project_id,
    creator_id,
    assignee_id,
    title,
    description,
    status,
    priority,
    due_date,
    created_at,
    updated_at
FROM tasks
WHERE project_id = $1
ORDER BY created_at DESC;


-- name: UpdateTask :one
UPDATE tasks
SET
    assignee_id = $2,
    title = $3,
    description = $4,
    status = $5,
    priority = $6,
    due_date = $7,
    updated_at = now()
WHERE id = $1
RETURNING
    id,
    project_id,
    creator_id,
    assignee_id,
    title,
    description,
    status,
    priority,
    due_date,
    created_at,
    updated_at;


-- name: DeleteTask :exec
DELETE
FROM tasks
WHERE id = $1;


-- name: ListTasksByAssignee :many
SELECT
    id,
    project_id,
    creator_id,
    assignee_id,
    title,
    description,
    status,
    priority,
    due_date,
    created_at,
    updated_at
FROM tasks
WHERE assignee_id = $1
ORDER BY created_at DESC;


-- name: ListTasksByStatus :many
SELECT
    id,
    project_id,
    creator_id,
    assignee_id,
    title,
    description,
    status,
    priority,
    due_date,
    created_at,
    updated_at
FROM tasks
WHERE project_id = $1
AND status = $2
ORDER BY created_at DESC;