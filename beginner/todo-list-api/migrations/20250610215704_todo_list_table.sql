-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS todo_list(
    task_id SERIAL PRIMARY KEY,
    task TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todo_list;
-- +goose StatementEnd
