package service

import (
	"Auth_Api_Gateway/config"
	db "Auth_Api_Gateway/db/repositories"
	"Auth_Api_Gateway/dto"
	"Auth_Api_Gateway/models"
	"Auth_Api_Gateway/utils"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById() error
	CreateUser(string, string, string) error
	LoginUser(*dto.LoginUserRequestDTO) (string, error)
	GetUserByEmail(string)(*models.User,error)
}

type UserServiceImp struct {
	UserRepository db.UserRepository
}

// constructor
func NewUserService(_userRepository db.UserRepository) UserService {
	// here we are depending on other to give object we are not creating inside here
	// this is dependeny injection
	return &UserServiceImp{
		UserRepository: _userRepository,
	}
}

func (u *UserServiceImp) GetUserById() error {
	fmt.Println("Fetching user in UserService")
	u.UserRepository.GetById()
	return nil
}

func (u *UserServiceImp) GetUserByEmail(email string)(*models.User,error){
	user,err:=u.UserRepository.GetByEmail(email)
	if err!=nil{
		if errors.Is(err,sql.ErrNoRows){
			fmt.Printf("user with email %s not found",email)
			return nil, fmt.Errorf("user with email %s not found",email)
		}
		return nil,err
	}
	return user,nil
}

func (u *UserServiceImp) CreateUser(username string, email string, password string) error {
	_,err :=  u.GetUserByEmail(email)
	if err == nil {
		return fmt.Errorf("user with email %s already exists",email)
	}else if !strings.Contains(err.Error(),"not found") {
		return err
	}
	// proceed email is unique

	fmt.Println("Creating user in UserService")
	HassPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	errRepo:=u.UserRepository.Create(username, email, HassPassword)
	if errRepo != nil {
		return fmt.Errorf("failed to create user in repository: %v", err)
	}
	fmt.Println("User created successfully in service")
	return nil
}

func (u *UserServiceImp) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {
	fmt.Println("LoginUser called in Service layer")

	// var userss *dto.LoginUserRequestDTO
	// email := "nik@gmail.com"
	// pswd := "1234"
	// userss = &dto.LoginUserRequestDTO{
	// 	Email:    email,
	// 	Password: pswd,
	// }
	// err := utils.Validator.Struct(userss)
	// if err != nil {
	// 	// Validation failed, handle the error
	// 	return "",err
	// }

	secret := []byte(config.GetString("JWT_SECRET_KEY", "hello"))
	user, err := u.UserRepository.GetByEmail(payload.Email)
	if user == nil {
		fmt.Println("No user with given eamil found")
		return "", err
	} else if err != nil {
		fmt.Println("Error in fetching the user")
		return "", err
	}
	if utils.CheckpasswordHash(payload.Password, user.Password) {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": payload.Email,
			"id":    user.Id,
			"iss":   "auth",
			"sub":   user.Id,
			"exp":   time.Now().Add(time.Hour).Unix(),
			"iat":   time.Now().Unix(),
		})
		token, err := claims.SignedString(secret)
		if err != nil {
			fmt.Println("Error in creating the token ")
			return "", err
		}
		return token, nil
	} else {
		fmt.Println("Incorrect password")
	}

	return "", nil
}
