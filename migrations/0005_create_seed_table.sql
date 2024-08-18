-- +goose Up
-- +goose StatementBegin
CREATE TABLE tickets.seed (
  "seeded" bool
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tickets.seed;
-- +goose StatementEnd
