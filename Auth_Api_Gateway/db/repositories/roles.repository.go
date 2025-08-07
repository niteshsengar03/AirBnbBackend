package db

import (
	"Auth_Api_Gateway/models"
	"database/sql"
)

type RoleRepository interface {
	GetRoleById(id int64) (*models.Roles, error)
	GetRoleByName(name string) (*models.Roles, error)
	GetAllRoles() ([]*models.Roles, error)
	CreateRole(name string, description string) (*models.Roles, error)
	DeleteRoleById(id int64) error
	UpdateRole(id int64, name string, description string) (*models.Roles, error)
}

type RoleRepositoryImpl struct {
	db *sql.DB
}

func NewRoleRepository(_db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: _db,
	}
}

func (r *RoleRepositoryImpl) GetRoleById(id int64) (*models.Roles, error) {
	query := "SELECT id,name,description,created_at,updated_at FROM roles WHERE id = ?"
	row := r.db.QueryRow(query, id)
	role := &models.Roles{}
	err := row.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepositoryImpl) GetRoleByName(name string) (*models.Roles, error) {
	query := "SELECT id,name,description,created_at,updated_at FROM roles WHERE name = ?"
	row := r.db.QueryRow(query, name)
	role := &models.Roles{}
	err := row.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepositoryImpl) GetAllRoles() ([]*models.Roles, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM roles"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Roles
	for rows.Next() {
		role := &models.Roles{}
		err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepositoryImpl) CreateRole(name string, description string) (*models.Roles, error) {
	query := "INSERT INTO roles (name, description, created_at, updated_at) VALUES (?, ?, NOW(), NOW())"
	result, err := r.db.Exec(query, name, description)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.Roles{
		Id:          id,
		Name:        name,
		Description: description,
		CreatedAt:   "", // Will be set by the database
		UpdatedAt:   "", // Will be set by the database
	}, nil

}

func (r *RoleRepositoryImpl) DeleteRoleById(id int64) error {
	query := "DELETE FROM roles WHERE id = ?"
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

func (r *RoleRepositoryImpl) UpdateRole(id int64, name string, description string) (*models.Roles, error) {
	query := "UPDATE roles SET name = ?, description = ?, updated_at = NOW() WHERE id = ?"
	result, err := r.db.Exec(query, name, description, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return r.GetRoleById(id)
}
