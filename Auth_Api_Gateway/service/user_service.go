package service

import (
	db "Auth_Api_Gateway/db/repositories"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface{
	GetUserById() error
	HashPassword(string) string
	CreateUser(string,string,string) error
}

type UserServiceImp struct{
	UserRepository db.UserRepository
}

// constructor to create newObject
func NewUserService (_userRepository db.UserRepository)UserService{
	// here we are depending on other to give object we are not creating inside here 
	// this is dependeny injection
	return &UserServiceImp{
		UserRepository: _userRepository,
	}
}

func(u *UserServiceImp) GetUserById() error{
	fmt.Println("Fetching user in UserService")
	// u.UserRepository.Create()
	// u.UserRepository.GetById()
	// u.UserRepository.GetAll()
	u.CreateUser("Alic","alice@gmail.com","Bob")
	return nil
}

func(u *UserServiceImp) CreateUser(username string,email string,password string)error{
	HassPassword := u.HashPassword(password)
	u.UserRepository.Create(username,email,HassPassword)
	return nil
}

func (u *UserServiceImp) HashPassword(password string)string{
	pass := []byte(password)
	hass,err := bcrypt.GenerateFromPassword(pass,bcrypt.DefaultCost)
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(hass))
	return string(hass)
}