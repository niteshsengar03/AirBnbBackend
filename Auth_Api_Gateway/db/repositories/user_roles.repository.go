package db

import (
	"Auth_Api_Gateway/models"
	"database/sql"
	"fmt"
	"strings"
)

type UserRoleRepository interface {
	GetUserRoles(userId int64) ([]*models.Roles, error)
	AssignRoleToUser(userId int64, roleId int64) error
	RemoveRoleFromUser(userId int64, roleId int64) error
	GetUserPermissions(userId int64) ([]*models.Permission, error)
	HasPermission(userId int64, permission string) (bool, error)
	HasRole(userId int64, roleName string) (bool, error)
	HasAllRoles(userId int64, roleNames []string) (bool, error)
	HasAnyRole(userId int64, roleNames []string) (bool, error)
}

type UserRoleRepositoryImp struct {
	db *sql.DB
}

func NewUserRoleRepository(_db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImp{
		db: _db,
	}
}

func (r *UserRoleRepositoryImp) GetUserRoles(userId int64) ([]*models.Roles, error) {
	query := `
		SELECT r.id, r.name, r.description, r.created_at, r.updated_at
		FROM user_roles ur
		JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ?`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Roles
	for rows.Next() {
		var role models.Roles
		if err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *UserRoleRepositoryImp) AssignRoleToUser(userId int64, roleId int64) error {
	query := `INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)`
	_, err := r.db.Exec(query, userId, roleId)
	return err
}

func (r *UserRoleRepositoryImp) RemoveRoleFromUser(userId int64, roleId int64) error {
	query := `DELETE FROM user_roles WHERE user_id = ? AND role_id = ?`
	_, err := r.db.Exec(query, userId, roleId)
	return err
}

func (r *UserRoleRepositoryImp) GetUserPermissions(userId int64) ([]*models.Permission, error) {
	query := `
		SELECT p.id, p.name, p.resource, p.action, p.created_at, p.updated_at
		FROM user_roles ur
		INNER JOIN role_permissions rp ON ur.role_id = rp.role_id
		INNER JOIN permissions p ON rp.permission_id = p.id
		WHERE ur.user_id = ?`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission
	for rows.Next() {
		var perm models.Permission
		if err := rows.Scan(&perm.Id, &perm.Name, &perm.Resource, &perm.Action, &perm.CreatedAt, &perm.UpdatedAt); err != nil {
			return nil, err
		}
		permissions = append(permissions, &perm)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *UserRoleRepositoryImp) HasPermission(userId int64, permission string) (bool, error) {
	query := `
		SELECT COUNT(*) > 0
		FROM user_roles ur
		INNER JOIN role_permissions rp ON ur.role_id = rp.role_id
		INNER JOIN permissions p ON rp.permission_id = p.id
		WHERE ur.user_id = ? AND p.name = ?`
	var exists bool
	err := r.db.QueryRow(query, userId, permission).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *UserRoleRepositoryImp) HasRole(userId int64, roleName string) (bool, error) {
	query := `
		SELECT COUNT(*) > 0
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ? AND r.name = ?`
	var exists bool
	err := r.db.QueryRow(query, userId, roleName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *UserRoleRepositoryImp) HasAllRoles(userId int64, roleNames []string) (bool, error) {
	// If no roles are specified, return true
	if len(roleNames) == 0 {
		return true, nil
	}

	query := `
        SELECT COUNT(*) = ?
        FROM user_roles ur
        INNER JOIN roles r ON ur.role_id = r.id
        WHERE ur.user_id = ? 
          AND r.name IN (?)
        GROUP BY ur.user_id
    `
	roleNameStr := strings.Join(roleNames, ",")
	row := r.db.QueryRow(query, len(roleNames), userId, roleNameStr)

	var hasAllRoles bool
	if err := row.Scan(&hasAllRoles); err != nil {
		if err == sql.ErrNoRows {
			return false, nil // No roles found for the user
		}
		return false, err // Return any other error
	}

	return hasAllRoles, nil
}

func (r *UserRoleRepositoryImp) HasAnyRole(userId int64, roleNames []string) (bool, error) {

	if len(roleNames) == 0 {
		return true, nil // If no roles are specified, return true
	}
	placeholders := strings.Repeat("?,", len(roleNames))
	placeholders = placeholders[:len(placeholders)-1]
	query := fmt.Sprintf("SELECT COUNT(*) > 0 FROM user_roles ur INNER JOIN roles r ON ur.role_id = r.id WHERE ur.user_id = ? AND r.name IN (%s)", placeholders)

	// Create args slice with userId first, then all roleNames
	args := make([]interface{}, 0, 1+len(roleNames))
	args = append(args, userId)
	for _, roleName := range roleNames {
		args = append(args, roleName)
	}

	row := r.db.QueryRow(query, args...)

	var hasAnyRole bool
	if err := row.Scan(&hasAnyRole); err != nil {
		if err == sql.ErrNoRows {
			return false, nil // No roles found for the user
		}
		return false, err // Return any other error
	}

	fmt.Println("hasAnyRole", hasAnyRole)

	return hasAnyRole, nil
}
