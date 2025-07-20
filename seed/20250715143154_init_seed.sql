-- +goose Up
-- +goose StatementBegin

-- Seed publishers
INSERT INTO publishers (name, address, phone) VALUES
  ('Penguin Books', '80 Strand, London, UK', '123-456-7890'),
  ('O''Reilly Media', '1005 Gravenstein Hwy N, Sebastopol, CA', '987-654-3210');

-- Seed authors
INSERT INTO authors (first_name, last_name, bio) VALUES
  ('George', 'Orwell', 'Author of 1984 and Animal Farm.'),
  ('Douglas', 'Crockford', 'Author of JavaScript: The Good Parts.');

-- Seed books
INSERT INTO books (title, price, published_date, stock, publisher_id) VALUES
  ('1984', 15.99, '1949-06-08', 50, 1),
  ('JavaScript: The Good Parts', 25.00, '2008-05-15', 30, 2);

-- Seed users
INSERT INTO users (first_name, last_name, email, role) VALUES
  ('Alice', 'Johnson', 'alice@example.com', 'customer'),
  ('Bob', 'Smith', 'bob@example.com', 'admin');

-- Seed orders
INSERT INTO orders (order_date, total_amount, user_id) VALUES
  ('2025-07-10', 40.99, 1);

-- Seed book_author relationships
INSERT INTO book_author (book_id, author_id) VALUES
  (1, 1), -- 1984 by George Orwell
  (2, 2); -- JavaScript book by Crockford

-- Seed order_items
INSERT INTO order_items (order_id, book_id, quantity) VALUES
  (1, 1, 1),
  (1, 2, 1);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Reverse order for deletion (child → parent)
DELETE FROM order_items;
DELETE FROM book_author;
DELETE FROM orders;
DELETE FROM users;
DELETE FROM books;
DELETE FROM authors;
DELETE FROM publishers;

-- +goose StatementEnd

