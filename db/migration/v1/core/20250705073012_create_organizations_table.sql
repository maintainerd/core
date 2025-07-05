-- +goose Up

-- CREATE TABLE
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS organizations (
    organization_id       SERIAL PRIMARY KEY,
    organization_uuid     UUID UNIQUE NOT NULL,
    name                  VARCHAR(255) NOT NULL,
    description           TEXT,
    is_active             BOOLEAN DEFAULT TRUE,
    created_at            TIMESTAMPTZ DEFAULT now(),
    updated_at            TIMESTAMPTZ
);
-- +goose StatementEnd

-- ADD INDEXES
-- +goose StatementBegin
CREATE INDEX idx_organizations_organization_uuid ON organizations(organization_uuid);
CREATE INDEX idx_organizations_name ON organizations(name);
CREATE INDEX idx_organizations_is_active ON organizations(is_active);
-- +goose StatementEnd

-- +goose Down

-- DROP INDEXES
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_organizations_organization_uuid;
DROP INDEX IF EXISTS idx_organizations_name;
DROP INDEX IF EXISTS idx_organizations_is_active;
-- +goose StatementEnd

-- DROP TABLE
-- +goose StatementBegin
DROP TABLE IF EXISTS organizations;
-- +goose StatementEnd
