package db

// import(
// 	"database/sql"
// 	"github.com/go-sql-driver/mysql"
// )

// Storage Struct take out of the responsibitly from
// service layer to make objects of all repository

type Storage struct {
	userRepository UserRepository
}

func NewStorage() *Storage {
	return &Storage{
		userRepository: &UserRepositoryImp{},
	}
}
