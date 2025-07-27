package controller

import (
	"Auth_Api_Gateway/dto"
	"Auth_Api_Gateway/service"
	"Auth_Api_Gateway/utils"
	"fmt"
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
	fmt.Println("Fetching user by ID in UserController")
	userId := r.URL.Query().Get("id")
	if userId == "" {
		userId = r.Context().Value("userId").(string)
	}
	fmt.Println("User id from context or querry:",userId)
	
	if userId == ""{
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User id is required", fmt.Errorf("missing user id"))
		return
	}
	user,err:=uc.UserService.GetUserById(userId)
	if err!=nil{
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"Falied to fetch user",fmt.Errorf("user with ID %s not found", userId))
		return
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"user feteched successfully",user)
}

func (uc *UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := uc.UserService.GetAllUser()
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error in getting all the users", err)
	} else {
		utils.WriteJsonSuccessResponse(w, http.StatusOK, "Fetched all users succesfully", users)
	}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// adding .(*dto.LoginUserRequestDTO) becuase earlier it was returning any, so to typecast it for desirable type
	payload := r.Context().Value("signupPaylaod").(*dto.SignupUserRequestDTO)

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

	payload := r.Context().Value("loginPayload").(*dto.LoginUserRequestDTO)
	token, err := uc.UserService.LoginUser(payload)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect password") {
			utils.WriteJsonErrorResponse(w, http.StatusUnauthorized, "Incorrect Password", err)
			return
		}
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to login User", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successful", token)

}
