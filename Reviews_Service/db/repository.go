package db

import "database/sql"

type ReviewRepository interface{
	Create() error
	GetAll() error
	GetById() error
	DeleteById() error

}
// ReviewRepositoryImplmentation
type ReviewRepositoryImp struct{
	db *sql.DB
}

func NewRepository(_db *sql.DB) ReviewRepository{
	return &ReviewRepositoryImp{
		db : _db,
	}
}

func (r *ReviewRepositoryImp) Create() error{
	return nil
}

func (r *ReviewRepositoryImp) GetAll() error{
	return nil
}

func (r *ReviewRepositoryImp) GetById() error{
	return nil
}

func (r *ReviewRepositoryImp) DeleteById() error{
	return nil
}

