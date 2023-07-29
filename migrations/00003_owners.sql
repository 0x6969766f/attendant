-- +goose Up
-- +goose StatementBegin
CREATE TABLE owners (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE owners;
-- +goose StatementEnd
