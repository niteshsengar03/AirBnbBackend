package router

import (
	"Auth_Api_Gateway/controller"
	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	UserController *controller.UserController
}

func NewUserRouter(_userController *controller.UserController) Router {
	return &UserRouter{
		UserController: _userController,
	}
}

func (u *UserRouter) Register(r chi.Router) {
	r.Get("/profile", u.UserController.GetUserById)
	r.Get("/signup",u.UserController.CreateUser)
	r.Get("/signin",u.UserController.LoginUser)
}
