package router

import (
	"Auth_Api_Gateway/controller"
	"Auth_Api_Gateway/middlewares"
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
	r.Get("/users",u.UserController.GetAllUser)
	r.Get("/profile", u.UserController.GetUserById)
	r.With(middlewares.UserCreateRequestValidator).Post("/signup", u.UserController.CreateUser)
	r.With(middlewares.UserLoginRequestValidator).Post("/signin", u.UserController.LoginUser)
}
