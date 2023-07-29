-- +goose Up
-- +goose StatementBegin
CREATE TABLE games (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  link TEXT,
  image TEXT,
  owner_id INT REFERENCES owners (id),
  console_id INT REFERENCES consoles (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE games;
-- +goose StatementEnd
