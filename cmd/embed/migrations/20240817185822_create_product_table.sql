-- +goose Up
-- +goose StatementBegin
CREATE TABLE product (
  "name" VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product;
-- +goose StatementEnd
