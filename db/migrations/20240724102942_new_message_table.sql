-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS messages(
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL ,
    is_processed BOOLEAN DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE messages;
-- +goose StatementEnd
