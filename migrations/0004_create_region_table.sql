-- +goose Up
-- +goose StatementBegin
CREATE TABLE tickets.region (
  "name" VARCHAR(255) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tickets.region;
-- +goose StatementEnd
