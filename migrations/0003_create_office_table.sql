-- +goose Up
-- +goose StatementBegin
CREATE TABLE production.office (
  "name" VARCHAR(255) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE production.office;
-- +goose StatementEnd
