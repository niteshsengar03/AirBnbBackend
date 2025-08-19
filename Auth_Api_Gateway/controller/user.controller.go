package controller

import (
	"Auth_Api_Gateway/dto"
	"Auth_Api_Gateway/service"
	"Auth_Api_Gateway/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
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
	fmt.Println("User id from context or querry:", userId)

	if userId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User id is required", fmt.Errorf("missing user id"))
		return
	}
	user, err := uc.UserService.GetUserById(userId)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Falied to fetch user", fmt.Errorf("user with ID %s not found", userId))
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "user feteched successfully", user)
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

func (uc *UserController) GetVerificationByToken(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	userId := chi.URLParam(r, "userId")
	if token == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Token is required", fmt.Errorf("missing token"))
		return
	}
	if userId == "" {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "userId is required", fmt.Errorf("missing token"))
		return
	}

	veri, err := uc.UserService.GetVerificationByToken(token)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Falied to fetch", err)
		return
	}
	if err == nil && veri == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "Verification token not found", fmt.Errorf("token not found"))
		return
	}
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid userId", err)
		return
	}
	// Ensure token belongs to the same user
	if userIdInt != veri.UserId {
		utils.WriteJsonErrorResponse(w, http.StatusUnauthorized, "User is not Authorised to verify", fmt.Errorf("unAuthorised access"))
		return
	}

	layout := "2006-01-02 15:04:05"

	expiryTime, err := time.Parse(layout, veri.ExpiresAt)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Invalid token expiry format", err)
		return
	}
	// Check if token is already expired
	if expiryTime.Before(time.Now()) {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Toeken is expired", fmt.Errorf("token expired"))
		return
	}
	// Mark user as verified
	markErr := uc.UserService.MarkUserVerified(userIdInt)
	if markErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "cannot mark the user", markErr)
		return
	}
	// Delete the verification token after successful verification
	deleteErr := uc.UserService.DeleteVerificationToken(token)
	if deleteErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "cannot delete the token", deleteErr)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "user verifed successfully", veri)
}
