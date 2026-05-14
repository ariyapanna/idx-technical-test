DROP INDEX IF EXISTS idx_todos_title;
DROP INDEX IF EXISTS idx_todos_completed;
DROP INDEX IF EXISTS idx_todos_category_id;
DROP INDEX IF EXISTS idx_todos_priority;
DROP INDEX IF EXISTS idx_todos_created_at;
DROP TABLE IF EXISTS todos;
DROP TYPE IF EXISTS todo_priority;