-- +goose Up
-- +goose StatementBegin
ALTER TABLE users 
ADD COLUMN verified BOOLEAN DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN verified;
-- +goose StatementEnd
