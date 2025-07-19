package db

import (
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Create() error
}

type UserRepositoryImp struct {
	db *sql.DB // DB instance given by SQL
}

func NewRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImp{
		db:_db,
	}
}

func (u *UserRepositoryImp) Create() error {
	rows,err:=u.db.Query("Select * from Users")
	if err!=nil{
		fmt.Println("cannot querry")
	}
	defer rows.Close()

for rows.Next() {
    var id int
    var name string
    if err := rows.Scan(&id, &name); err != nil {
        fmt.Println("error scanning row:", err)
        continue
    }
    fmt.Printf("User: ID=%d, Name=%s\n", id, name)
}

	fmt.Println("Creating user in UserRepo")
	return nil
}
