-- +goose Up
-- +goose StatementBegin
ALTER TABLE books ADD COLUMN publisher_id INTEGER;

ALTER TABLE books
    ADD CONSTRAINT fk_books_publisher
    FOREIGN KEY (publisher_id) REFERENCES publishers(id) ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE books DROP CONSTRAINT IF EXISTS fk_books_publisher;
ALTER TABLE books DROP COLUMN IF EXISTS publisher_id;
-- +goose StatementEnd

