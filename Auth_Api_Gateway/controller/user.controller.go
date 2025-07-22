package controller

import (
	"Auth_Api_Gateway/service"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

// constructor
func NewUserController(_userService service.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	uc.UserService.GetUserById()
	w.Write([]byte("User by ID called sucessfully"))
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request){
	uc.UserService.CreateUser("nik","nik@gmail.com","1234")
	w.Write([]byte("User created sucessfully"))
}


func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request){
	uc.UserService.LoginUser();
	w.Write([]byte("User loged in sucesfuly"))
}

