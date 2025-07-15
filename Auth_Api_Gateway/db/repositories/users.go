package db

import "database/sql"

type UserRepository interface{
	createUser() error
}

type UserRepositoryImp struct{
	db *sql.DB // DB instance given by SQL
}

func(u *UserRepositoryImp) createUser()error{
	return nil
}