-- +goose Up
-- +goose StatementBegin
CREATE TABLE region (
  "name" VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE region;
-- +goose StatementEnd
