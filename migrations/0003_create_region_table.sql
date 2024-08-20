-- +goose Up
-- +goose StatementBegin
CREATE TABLE production.region (
  "name" VARCHAR(255) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE production.region;
-- +goose StatementEnd
