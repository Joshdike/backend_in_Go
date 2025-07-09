-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS posts(
    post_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    author_id INTEGER NOT NULL REFERENCES users(user_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS comments(
    comment_id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL REFERENCES posts(post_id) ON DELETE CASCADE,
    author_id INTEGER NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    body TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
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
CREATE TRIGGER update_users_modtime
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION update_modified_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_posts_modtime ON posts;
DROP TRIGGER IF EXISTS update_users_modtime ON users;
DROP FUNCTION IF EXISTS update_modified_column();
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
