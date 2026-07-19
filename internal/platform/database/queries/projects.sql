-- name: CreateProject :one
INSERT INTO projects (
    name,
    description,
    owner_id,
    visibility
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;


-- name: GetProjectByID :one
SELECT *
FROM projects
WHERE id = $1
LIMIT 1;


-- name: ListProjectsByUser :many
SELECT p.*
FROM projects p
JOIN project_members pm
    ON pm.project_id = p.id
WHERE pm.user_id = $1
ORDER BY p.created_at DESC;


-- name: UpdateProject :one
UPDATE projects
SET
    name = $2,
    description = $3,
    visibility = $4,
    updated_at = now()
WHERE id = $1
RETURNING *;


-- name: DeleteProject :exec
DELETE
FROM projects
WHERE id = $1;


-- name: AddProjectMember :exec
INSERT INTO project_members (
    project_id,
    user_id,
    role
)
VALUES (
    $1,
    $2,
    $3
)
ON CONFLICT (project_id, user_id)
DO NOTHING;


-- name: RemoveProjectMember :exec
DELETE
FROM project_members
WHERE project_id = $1
AND user_id = $2;


-- name: GetProjectRole :one
SELECT role
FROM project_members
WHERE project_id = $1
AND user_id = $2
LIMIT 1;


-- name: ProjectExistsForUser :one
SELECT EXISTS(
    SELECT 1
    FROM project_members
    WHERE project_id = $1
    AND user_id = $2
);