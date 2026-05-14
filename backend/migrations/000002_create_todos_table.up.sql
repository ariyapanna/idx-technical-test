CREATE TYPE todo_priority AS ENUM (
    'high',
    'medium',
    'low'
);

CREATE TABLE IF NOT EXISTS todos (
    id BIGSERIAL PRIMARY KEY,
    category_id BIGINT REFERENCES categories(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    priority todo_priority NOT NULL,
    due_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_todos_title
ON todos(title);

CREATE INDEX IF NOT EXISTS idx_todos_completed
ON todos(completed);

CREATE INDEX IF NOT EXISTS idx_todos_category_id
ON todos(category_id);

CREATE INDEX IF NOT EXISTS idx_todos_priority
ON todos(priority);

CREATE INDEX IF NOT EXISTS idx_todos_created_at
ON todos(created_at DESC);