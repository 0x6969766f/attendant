-- +goose Up
-- +goose StatementBegin
CREATE TABLE consoles (
  id SERIAL PRIMARY KEY,
  brand TEXT NOT NULL,
  model TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE consoles;
-- +goose StatementEnd
