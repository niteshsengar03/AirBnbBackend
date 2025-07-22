package db

import (
	"Auth_Api_Gateway/models"
	"database/sql"
	"fmt"
	// "github.com/ydb-platform/ydb-go-sdk/v3/query"
)

type UserRepository interface {
	GetById() error
	Create(string) error
	GetAll() error
	// DeleteById() error
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
	row := u.db.QueryRow(query, 3)

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

func (u *UserRepositoryImp) Create(hass string) error {

	// prepare querry
	query := "INSERT INTO users(username,email,password)VALUES(?,?,?)"
	
	result, err := u.db.Exec(query, "amit", "nik@gmail.com", hass)

	if err != nil {
		fmt.Println("Canot insert data")
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Cannot insert data")
	}
	fmt.Println("no.rows affeted ", rows)
	return nil
}

func (u *UserRepositoryImp) GetAll() error {
	// prepare querry
	query := "Select * from users"
	rows, err := u.db.Query(query)
	if err != nil {
		return err
	}

	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			fmt.Print("can not scan", err)
			return err
		}

		users = append(users, user)
	}
	for _, user := range users {
		fmt.Println(*user)
	}
	return nil
}