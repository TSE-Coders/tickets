-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA tickets;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA tickets CASCADE;
-- +goose StatementEnd
