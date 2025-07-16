package service

import (
	db "Auth_Api_Gateway/db/repositories"
	"fmt"
)

type UserService interface{
	CreateUser() error
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

func(u *UserServiceImp) CreateUser() error{
	fmt.Println("Creating user in UserService")
	u.UserRepository.Create()
	return nil
}