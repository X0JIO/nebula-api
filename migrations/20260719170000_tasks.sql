-- +goose Up

CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    project_id UUID NOT NULL
        REFERENCES projects(id)
        ON DELETE CASCADE,

    creator_id UUID NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    assignee_id UUID
        REFERENCES users(id)
        ON DELETE SET NULL,

    title TEXT NOT NULL,

    description TEXT NOT NULL DEFAULT '',

    status TEXT NOT NULL DEFAULT 'todo'
        CHECK (
            status IN (
                'todo',
                'in_progress',
                'review',
                'done',
                'archived'
            )
        ),

    priority TEXT NOT NULL DEFAULT 'medium'
        CHECK (
            priority IN (
                'low',
                'medium',
                'high',
                'critical'
            )
        ),

    due_date TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_tasks_project
ON tasks(project_id);

CREATE INDEX idx_tasks_assignee
ON tasks(assignee_id);

CREATE INDEX idx_tasks_status
ON tasks(status);

CREATE INDEX idx_tasks_priority
ON tasks(priority);

-- +goose Down

DROP TABLE IF EXISTS tasks;