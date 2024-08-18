-- +goose Up
-- +goose StatementBegin
CREATE TABLE tickets.product (
  "name" VARCHAR(255) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tickets.product;
-- +goose StatementEnd
