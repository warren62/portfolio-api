-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts
(
    id            UUID      PRIMARY KEY,
    name          TEXT      NOT NULL,
    email         TEXT      NOT NULL,
    message       TEXT      NOT NULL,
    description   TEXT      NOT NULL,
    subject       SMALLINT  NOT NULL,
    created_date  TIMESTAMP NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts;

-- +goose StatementEnd
