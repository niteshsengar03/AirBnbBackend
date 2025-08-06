package db

import (
	"Auth_Api_Gateway/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetById(id string) (*models.User,error)
	Create(username string,email string,hassPassword string) error
	GetAll() ([]*models.User, error)
	GetByEmail(email string) (*models.User, error)
	DeleteById(id int64) error
}

type UserRepositoryImp struct {
	db *sql.DB // DB instance given by SQL
}

// cannot return *UserRepository becuase it's a inteface not struct
func NewRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImp{
		db: _db,
	}
}

func (u *UserRepositoryImp) GetById(id string) (*models.User,error) {

	fmt.Println("Fetching user in UserRepo")
	// question based sytnax helps in avoiding sql injection

	// prepare the querry
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?"

	// exccute the querry
	row := u.db.QueryRow(query, id)

	// process the result
	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil,err
		} else {
			fmt.Print("Error scanning user:", err)
			return nil,err
		}
	}

	// print user details
	fmt.Println("User fetched succesfully", *user)
	return user,nil
}

func (u *UserRepositoryImp) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id,username,password FROM users WHERE email=?"
	row := u.db.QueryRow(query, email)
	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// fmt.Println("No user found with the given Email")
			return nil, err
		} else {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
	}
	// fmt.Println("User fetched succesfully", *user)
	return user, nil
}

func (u *UserRepositoryImp) Create(username string, email string, hassPassword string) error {

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

func (u *UserRepositoryImp) GetAll() ([]*models.User, error) {
	// prepare querry
	query := "Select id, username, email, created_at, updated_at from users"
	rows, err := u.db.Query(query)
	if err != nil {
		fmt.Println("Error fetching the user: ", err)
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			fmt.Print("can not scan", err)
			return nil, err
		}

		users = append(users, user)
	}
	// for _, user := range users {
	// 	fmt.Println(*user)
	// }
	return users, nil
}

func (u *UserRepositoryImp) DeleteById(id int64) error {
	query := "DELETE FROM users WHERE id=?"
	result, err := u.db.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return err
	}
	rowsAffected, rowErr := result.RowsAffected()
	if rowErr != nil {
		fmt.Println("Error getting rows affected:", rowErr)
		return rowErr
	}
	if rowsAffected == 0 {
		fmt.Println("No rows where affected, user not deleted")
		return nil
	}
	fmt.Println("User deleted successfully, rows affected: ", rowsAffected)
	return nil
}
