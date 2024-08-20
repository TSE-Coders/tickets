-- +goose Up
-- +goose StatementBegin
CREATE TABLE production.ticket (
  "id" BIGSERIAL UNIQUE NOT NULL,
  "name" VARCHAR(255) UNIQUE NOT NULL,
  "timestamp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE production.ticket;
-- +goose StatementEnd
