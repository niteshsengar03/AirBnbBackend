package db

import (
	"fmt"
)

type UserRepository interface {
	Create() error
}

type UserRepositoryImp struct {
	// db *sql.DB // DB instance given by SQL
}

func NewRepository() UserRepository {
	return &UserRepositoryImp{
		// db:db,
	}
}

func (u *UserRepositoryImp) Create() error {
	fmt.Println("Creating user in UserRepo")
	return nil
}
