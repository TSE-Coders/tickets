-- +goose Up
-- +goose StatementBegin
CREATE TABLE production.ticket (
  "id" BIGSERIAL UNIQUE NOT NULL,
  "product" VARCHAR(255) NOT NULL,
  "office" VARCHAR(255) NOT NULL,
  "difficulty" SMALLINT NOT NULL,
  "timestamp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE production.ticket;
-- +goose StatementEnd
