-- +goose Up
-- +goose StatementBegin
CREATE TABLE production.product (
  "name" VARCHAR(255) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE production.product;
-- +goose StatementEnd
