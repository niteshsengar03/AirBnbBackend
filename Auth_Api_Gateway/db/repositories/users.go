package db

import (
	"Auth_Api_Gateway/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetById() error
}

type UserRepositoryImp struct {
	db *sql.DB // DB instance given by SQL
}

func NewRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImp{
		db: _db,
	}
}

func (u *UserRepositoryImp) GetById() error {

	fmt.Println("Fetching user in UserRepo")
	// question based sytnax helps in avoiding sql injection

	// prepare the querry
	query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = ?"

	// exccute the querry
	row := u.db.QueryRow(query, 1)

	// process the result
	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil
		} else {
			fmt.Print("Error scanning user:", err)
			return err
		}
	}

	// print user details
	fmt.Println("User fetched succesfully", *user)
	return nil
}
