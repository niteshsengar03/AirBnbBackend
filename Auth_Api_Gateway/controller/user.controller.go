package controller

import (
	"Auth_Api_Gateway/dto"
	"Auth_Api_Gateway/service"
	"Auth_Api_Gateway/utils"
	"net/http"
	"strings"
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
	var payload dto.SignupUserRequestDTO
	Jsonerr := utils.ReadJsonBody(r, &payload)
	if Jsonerr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Error in reading the payload", Jsonerr)
		return
	}
	utils.Validator.Struct(payload)

	if ValidationErr := utils.Validator.Struct(payload); ValidationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input type", ValidationErr)
		return
	}

	err := uc.UserService.CreateUser(payload.Username, payload.Email, payload.Password)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User already exists", err)
			return
		}
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User created successfully", nil)	
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
