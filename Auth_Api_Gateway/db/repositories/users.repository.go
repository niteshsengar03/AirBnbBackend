package db

import (
	"Auth_Api_Gateway/models"
	"database/sql"
	"fmt"
	"time"
)

type UserRepository interface {
	GetById(id string) (*models.User, error)
	Create(username string, email string, hassPassword string) (int64, error)
	GetAll() ([]*models.User, error)
	GetByEmail(email string) (*models.User, error)
	DeleteById(id int64) error
	CreateVerification(userId int64, token string, expireAt time.Time) error
	GetVerificationByToken(token string) (*models.Verification, error)
	MarkUserVerified(userId int64) error
	DeleteVerificationToken(token string) error
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

func (u *UserRepositoryImp) GetById(id string) (*models.User, error) {

	fmt.Println("Fetching user in UserRepo")
	// question based sytnax helps in avoiding sql injection

	// prepare the querry
	query := "SELECT id, username, email, created_at, updated_at,verified FROM users WHERE id = ?"

	// exccute the querry
	row := u.db.QueryRow(query, id)

	// process the result
	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt,&user.Verified)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil, err
		} else {
			fmt.Print("Error scanning user:", err)
			return nil, err
		}
	}

	// print user details
	fmt.Println("User fetched succesfully", *user)
	return user, nil
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

func (u *UserRepositoryImp) Create(username string, email string, hassPassword string) (int64, error) {
	query := "INSERT INTO users(username,email,password) VALUES (?, ?, ?)"

	result, err := u.db.Exec(query, username, email, hassPassword)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user: %w", err)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to fetch last inserted user ID: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return userId, fmt.Errorf("user inserted, but failed to fetch rows affected: %w", err)
	}

	if rows == 0 {
		return userId, fmt.Errorf("user not created: no rows affected")
	}

	fmt.Println("User created successfully. Rows affected:", rows)
	return userId, nil
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

func (u *UserRepositoryImp) CreateVerification(userId int64, token string, expireAt time.Time) error {
	query := "INSERT INTO verification_tokens(user_id,token,expires_at)VALUES(?,?,?)"
	result, err := u.db.Exec(query, userId, token, expireAt)

	if err != nil {
		fmt.Println("Canot insert data")
		return fmt.Errorf("failed to insert verificationtoken: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("verification inserted, but failed to fetch rows affected: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("verification not created: no rows affected")
	}
	return nil
}

func (u *UserRepositoryImp) GetVerificationByToken(token string) (*models.Verification, error) {
	query := "SELECT id, user_id, token, expires_at, created_at FROM verification_tokens WHERE token = ?"
	row := u.db.QueryRow(query, token)
	veri := &models.Verification{}
	err := row.Scan(&veri.Id, &veri.UserId, &veri.Token, &veri.ExpiresAt, &veri.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			//No Verification found with the given ID
			return nil, nil
		} else {
			//Error scanning Verification
			return nil, err
		}
	}
	return veri, nil
}

func (u *UserRepositoryImp) MarkUserVerified(userId int64) error {
	query := "UPDATE users SET verified = TRUE WHERE id = ?"
	result, err := u.db.Exec(query, userId)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", userId)
	}

	return nil
}

func (u *UserRepositoryImp) DeleteVerificationToken(token string) error {
	query := "DELETE FROM verification_tokens WHERE token = ?"
	result, err := u.db.Exec(query, token)
	if err != nil {
		return fmt.Errorf("failed to delete verification token: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("verification token deletion attempted, but failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no verification token found for deletion")
	}

	return nil
}
