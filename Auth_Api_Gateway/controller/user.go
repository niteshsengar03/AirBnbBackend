package controller

import (
	"Auth_Api_Gateway/service"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

// constructor
func NewUserController(_userService service.UserService) *UserController{
	return  &UserController{
		UserService: _userService,
	}
}

func (uc  *UserController) RegisterUser(w http.ResponseWriter,r *http.Request){
	uc.UserService.CreateUser()
	w.Write([]byte("User registered sucessfully"))
}
