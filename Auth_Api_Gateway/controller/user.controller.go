package controller

import (
	"Auth_Api_Gateway/dto"
	"Auth_Api_Gateway/service"
	"Auth_Api_Gateway/utils"
	"fmt"
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

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	uc.UserService.CreateUser("nik", "ndik@gmail.com", "12d34")
	w.Write([]byte("User created sucessfully"))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login User called in controller")

	var payload dto.LoginUserRequestDTO

	Jsonerr := utils.ReadJsonBody(r, &payload)
	if Jsonerr != nil {
		w.Write([]byte("Something went wrong"))
		return
	}
	// two ways to check error
	utils.Validator.Struct(payload)

	if ValidationErr := utils.Validator.Struct(payload); ValidationErr != nil {
		w.Write([]byte("Invalid input type"))
		fmt.Println("Validation Error: ", ValidationErr)
		return
	}

	token, err := uc.UserService.LoginUser(&payload)
	if err != nil {
		w.Write([]byte("Something went wrong"))
		return
	}
	response := map[string]any{
		"token":    "User logged in successfully",
		"data":     token,
		"successs": true,
		"error":    nil,
	}
	utils.WriteJsonResponse(w, 200, response)

}
