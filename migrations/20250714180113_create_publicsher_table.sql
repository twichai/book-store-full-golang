-- +goose Up
-- +goose StatementBegin
CREATE TABLE publishers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    phone TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE publishers;
-- +goose StatementEnd

