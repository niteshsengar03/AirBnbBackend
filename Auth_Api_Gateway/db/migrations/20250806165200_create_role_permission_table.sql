-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS role_permissions(
    id SERIAL PRIMARY KEY,
    role_id BIGINT UNSIGNED  NOT NULL,
    permission_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);

-- Assign all existing permissions to role with ID 1 (admins have all permission)
-- INSERT INTO role_permissions (role_id,permission_id) SELECT 1, id from permissions;

-- INSERT INTO role_permissions(role_id,permission_id)
-- SELECT 2, id FROM permissions WHERE name In('user:read'); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role_permissions;
-- +goose StatementEnd
