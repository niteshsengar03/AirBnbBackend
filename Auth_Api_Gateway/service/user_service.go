package service

type UserServie interface{
	createUser() error
}

type UserServiceImp struct{

}

func NewUserService ()UserServie{
	return &UserServiceImp{}
}

func(u *UserServiceImp) createUser() error{
	return nil
}