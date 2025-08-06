-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS permissions(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE,
    description TEXT,
    resource VARCHAR(100) NOT NULL,
    action  VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)
-- seder data
-- INSERT INTO permissions (name, description, resource, action)
-- VALUES 
--     ('user:read', 'Permission to view user details', 'user', 'read'),
--     ('user:write', 'Permission to edit user details', 'user', 'write'),
--     ('user:delete', 'Permission to delete a user', 'user', 'delete'),
--     ('role:read', 'Permission to view role details', 'role', 'read'),
--     ('role:write', 'Permission to edit role details', 'role', 'write'),
--     ('role:delete', 'Permission to delete a role', 'role', 'delete'),
--     ('role:manage', 'Permission to manage roles', 'role', 'manage'),
--     ('permission:read', 'Permission to view permission details', 'permission', 'read'),
--     ('permission:write', 'Permission to edit permission details', 'permission', 'write'),
--     ('permission:delete', 'Permission to delete a permission', 'permission', 'delete'),
--     ('permission:manage', 'Permission to manage permissions', 'permission', 'manage');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
