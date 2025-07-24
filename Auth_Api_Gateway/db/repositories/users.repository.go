package db

import (
	"Auth_Api_Gateway/models"
	"database/sql"
	"fmt"

	// "github.com/ydb-platform/ydb-go-sdk/v3/query"
	// "github.com/ydb-platform/ydb-go-sdk/v3/query"
)

type UserRepository interface {
	GetById() error
	Create(string,string,string) error
	GetAll() error
	GetByEmail(string) (*models.User,error)
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

func (u *UserRepositoryImp) GetByEmail(email string) (*models.User,error){
	query := "SELECT id,username,password FROM users WHERE email=?"
	row:= u.db.QueryRow(query,email)
	user := &models.User{}
	err := row.Scan(&user.Id,&user.Username,&user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// fmt.Println("No user found with the given Email")
			return nil,err
		} else {
			return nil,fmt.Errorf("error scanning user: %w", err)
		}
	}
	// fmt.Println("User fetched succesfully", *user)
	return user,nil
}

func (u *UserRepositoryImp) Create(username string,email string,hassPassword string) error {

	// prepare querry
	query := "INSERT INTO users(username,email,password)VALUES(?,?,?)"
	
	result, err := u.db.Exec(query, username, email, hassPassword)

	if err != nil {
		fmt.Println("Canot insert data")
		return fmt.Errorf("failed to insert user: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("user inserted, but failed to fetch rows affected: %w", err)
	}


	if rows == 0 {
		return fmt.Errorf("user not created: no rows affected")
	}

	fmt.Println("User created successfully. Rows affected:", rows)
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