-- +goose Up
-- +goose StatementBegin
CREATE TABLE production.seed (
  "seeded" bool
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE production.seed;
-- +goose StatementEnd
