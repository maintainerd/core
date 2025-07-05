-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS core_config (
    core_config_id      SERIAL PRIMARY KEY,
    core_config_uuid    UUID NOT NULL UNIQUE,
    version             VARCHAR(20) NOT NULL,
    is_active           BOOLEAN DEFAULT TRUE,
    is_applied          BOOLEAN DEFAULT TRUE,
    created_at          TIMESTAMPTZ DEFAULT now(),
    updated_at          TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS core_config;
-- +goose StatementEnd
