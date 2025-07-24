package controller

import (
	"Auth_Api_Gateway/dto"
	"Auth_Api_Gateway/service"
	"Auth_Api_Gateway/utils"
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

	var payload dto.LoginUserRequestDTO

	Jsonerr := utils.ReadJsonBody(r, &payload)
	if Jsonerr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Error in reading the payload", Jsonerr)
		return
	}
	// two ways to check error
	utils.Validator.Struct(payload)

	if ValidationErr := utils.Validator.Struct(payload); ValidationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input type", ValidationErr)
		return
	}

	token, err := uc.UserService.LoginUser(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to login User", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successful", token)

}
