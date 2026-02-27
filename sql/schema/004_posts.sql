-- +goose Up
CREATE TABLE posts(
    id           UUID PRIMARY KEY,
    feed_id      UUID NOT NULL REFERENCES feeds (id),
    title        TEXT NOT NULL,
    url          TEXT NOT NULL UNIQUE,
    description  TEXT NOT NULL,
    published_at TIMESTAMP,
    created_at   TIMESTAMP NOT NULL,
    updated_at   TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE posts;
