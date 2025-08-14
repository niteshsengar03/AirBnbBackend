package service

import (
	db "Auth_Api_Gateway/db/repositories"
	"Auth_Api_Gateway/models"
)

type RoleService interface {
	GetRoleById(id int64) (*models.Roles, error)
	GetRoleByName(name string) (*models.Roles, error)
	GetAllRoles() ([]*models.Roles, error)
	CreateRole(name string, description string) (*models.Roles, error)
	DeleteRoleById(id int64) error
	UpdateRole(id int64, name string, description string) (*models.Roles, error)
	GetRolePermissions(roleId int64) ([]*models.RolePermission, error)
	AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
	RemovePermissionFromRole(roleId int64, permissionId int64) error
	GetAllRolePermissions() ([]*models.RolePermission, error)
	AssignRoleToUser(userId int64, roleId int64) error
}

type RoleServiceImpl struct {
	RoleRepository           db.RoleRepository
	RolePermissionRepository db.RolePermissionRepository
	UserRoleRepository db.UserRoleRepository
}

func NewRoleService(roleRepo db.RoleRepository, rolePermRepo db.RolePermissionRepository, userRoleRepo db.UserRoleRepository) RoleService {
	return &RoleServiceImpl{
		RoleRepository:           roleRepo,
		RolePermissionRepository: rolePermRepo,
		UserRoleRepository :userRoleRepo,
	}
}

func (r *RoleServiceImpl) GetRoleById(id int64) (*models.Roles, error) {
	return r.RoleRepository.GetRoleById(id)
}

func (r *RoleServiceImpl) GetRoleByName(name string) (*models.Roles, error) {
	return r.RoleRepository.GetRoleByName(name)
}

func (r *RoleServiceImpl) GetAllRoles() ([]*models.Roles, error) {
	return r.RoleRepository.GetAllRoles()
}

func (r *RoleServiceImpl) CreateRole(name string, description string) (*models.Roles, error) {
	return r.RoleRepository.CreateRole(name, description)
}

func (r *RoleServiceImpl) DeleteRoleById(id int64) error {
	return r.RoleRepository.DeleteRoleById(id)
}

func (r *RoleServiceImpl) UpdateRole(id int64, name string, description string) (*models.Roles, error) {
	return r.RoleRepository.UpdateRole(id, name, description)
}

func (r *RoleServiceImpl) GetRolePermissions(roleId int64) ([]*models.RolePermission, error) {
	return r.RolePermissionRepository.GetRolePermissionByRoleId(roleId)
}
func (r *RoleServiceImpl) AddPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	return r.RolePermissionRepository.AddPermissionToRole(roleId, permissionId)
}
func (r *RoleServiceImpl) RemovePermissionFromRole(roleId int64, permissionId int64) error {
	return r.RolePermissionRepository.RemovePermissionFromRole(roleId, permissionId)
}

func (r *RoleServiceImpl) GetAllRolePermissions() ([]*models.RolePermission, error) {
	return r.RolePermissionRepository.GetAllRolePermissions()
}

func (r *RoleServiceImpl) AssignRoleToUser(userId int64, roleId int64) error {
	return r.UserRoleRepository.AssignRoleToUser(userId,roleId)
}