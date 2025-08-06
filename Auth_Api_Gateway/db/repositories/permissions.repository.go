package db

import (
	"Auth_Api_Gateway/models"
	"database/sql"
)

type PermissionRepository interface {
	GetPermissionById(id int64) (*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission, error)
	GetAllPermission() ([]*models.Permission, error)
	CreatePermission(name string, description string,resource string, action string) (*models.Permission, error)
	DeletePermissionById(id int64) error
	UpdatePermission(id int64, name string, description string) (*models.Permission, error)
}

type PermissionRepositoryImpl struct {
	db *sql.DB
}

func NewPermissionRepository(_db *sql.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db: _db,
	}
}

func (r *PermissionRepositoryImpl) GetPermissionById(id int64) (*models.Permission, error) {
	query := `SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE id = ?`
	row := r.db.QueryRow(query, id)
	perm := &models.Permission{}
	err := row.Scan(&perm.Id, &perm.Name, &perm.Description, &perm.Resource, &perm.Action, &perm.CreatedAt, &perm.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return perm, nil
}

func (r *PermissionRepositoryImpl) GetPermissionByName(name string) (*models.Permission, error) {
	query := `SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE name = ?`
	row := r.db.QueryRow(query, name)
	perm := &models.Permission{}
	err := row.Scan(&perm.Id, &perm.Name, &perm.Description, &perm.Resource, &perm.Action, &perm.CreatedAt, &perm.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return perm, nil
}

func (r *PermissionRepositoryImpl) GetAllPermission() ([]*models.Permission, error) {
	query := `SELECT id, name, description, resource, action, created_at, updated_at FROM permissions`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission
	for rows.Next() {
		perm := &models.Permission{}
		err := rows.Scan(&perm.Id, &perm.Name, &perm.Description, &perm.Resource, &perm.Action, &perm.CreatedAt, &perm.UpdatedAt)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, perm)
	}
	return permissions, nil
}

func (r *PermissionRepositoryImpl) CreatePermission(name string, description string,resource string, action string) (*models.Permission, error) {
	query := `INSERT INTO permissions (name, description,resource,action,created_at,updated_at) VALUES (?,?,?,?,NOW(),NOW())`
	result, err := r.db.Exec(query, name, description,resource,action)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r.GetPermissionById(id)
}

func (r *PermissionRepositoryImpl) DeletePermissionById(id int64) error {
	query := `DELETE FROM permissions WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *PermissionRepositoryImpl) UpdatePermission(id int64, name string, description string) (*models.Permission, error) {
	query := `UPDATE permissions SET name = ?, description = ? WHERE id = ?`
	_, err := r.db.Exec(query, name, description, id)
	if err != nil {
		return nil, err
	}
	return r.GetPermissionById(id)
}
