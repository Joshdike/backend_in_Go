-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts(
    post_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS comments(
    comment_id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL,
    name VARCHAR(100) NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE
);
CREATE OR REPLACE FUNCTION update_modified_column() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER update_posts_modtime
BEFORE UPDATE ON posts
FOR EACH ROW EXECUTE FUNCTION update_modified_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_posts_modtime ON posts;
DROP FUNCTION IF EXISTS update_modified_column();
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd
