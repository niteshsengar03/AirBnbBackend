package router

import (
	"Auth_Api_Gateway/controller"
	"Auth_Api_Gateway/middlewares"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	RoleController *controller.RoleController
}

func NewRoleRouter(rc *controller.RoleController) Router {
	return &RoleRouter{
		RoleController: rc,
	}
}

func (rr *RoleRouter) Register(r chi.Router) {
	r.Get("/roles/{id}", rr.RoleController.GetRoleById)
	r.Get("/roles", rr.RoleController.GetAllRoles)
	r.With(middlewares.CreateRoleRequestValidator).Post("/roles", rr.RoleController.CreateRole)
	r.With(middlewares.UpdateRoleRequestValidator).Put("/roles/{id}", rr.RoleController.UpdateRole)
	r.Delete("/roles/{id}", rr.RoleController.DeleteRole)

	r.Get("/roles/{id}/permissions", rr.RoleController.GetRolePermissions)
	r.With(middlewares.AssignPermissionRequestValidator).Post("/roles/{id}/permissions", rr.RoleController.AssignPermissionToRole)
	r.With(middlewares.RemovePermissionRequestValidator).Delete("/roles/{id}/permissions", rr.RoleController.RemovePermissionFromRole)
	r.Get("/role-permissions", rr.RoleController.GetAllRolePermissions)

	r.With(middlewares.JWTAuthMiddleware,middlewares.RequireAllRoles("admin")).Post("/roles/{userId}/assign/{roleId}", rr.RoleController.AssignRoleToUser)
} 
