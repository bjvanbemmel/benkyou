-- +goose Up
-- +goose StatementBegin
ALTER TABLE requirements
ALTER COLUMN estimate SET NOT NULL,
ALTER COLUMN estimate SET DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE requirements
ALTER COLUMN estimate DROP NOT NULL,
ALTER COLUMN estimate DROP DEFAULT;
-- +goose StatementEnd
