-- +goose Up

CREATE TABLE comments (

    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),

    task_id uuid NOT NULL
        REFERENCES tasks(id)
        ON DELETE CASCADE,

    author_id uuid NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    body text NOT NULL,

    created_at timestamptz NOT NULL DEFAULT now(),

    updated_at timestamptz NOT NULL DEFAULT now()

);

CREATE INDEX idx_comments_task
ON comments(task_id);

CREATE INDEX idx_comments_author
ON comments(author_id);

-- +goose Down

DROP TABLE comments;