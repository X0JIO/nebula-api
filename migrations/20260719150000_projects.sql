-- +goose Up

CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name TEXT NOT NULL,

    description TEXT NOT NULL DEFAULT '',

    owner_id UUID NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    visibility TEXT NOT NULL DEFAULT 'private',

    created_at TIMESTAMP NOT NULL DEFAULT now(),

    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX idx_projects_owner
ON projects(owner_id);

CREATE INDEX idx_projects_visibility
ON projects(visibility);



CREATE TABLE project_members (

    project_id UUID NOT NULL
        REFERENCES projects(id)
        ON DELETE CASCADE,

    user_id UUID NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    role TEXT NOT NULL DEFAULT 'member',

    joined_at TIMESTAMP NOT NULL DEFAULT now(),

    PRIMARY KEY (
        project_id,
        user_id
    )
);

CREATE INDEX idx_project_members_user
ON project_members(user_id);

CREATE INDEX idx_project_members_project
ON project_members(project_id);

-- +goose Down

DROP TABLE IF EXISTS project_members;

DROP TABLE IF EXISTS projects;