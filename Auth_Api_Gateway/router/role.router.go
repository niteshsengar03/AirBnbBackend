package router

import (
	"Auth_Api_Gateway/controller"

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
	r.Get("/roles",rr.RoleController.GetAllRoles)
}
