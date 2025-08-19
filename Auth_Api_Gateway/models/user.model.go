package models

type User struct {
	Id        int64
	Username  string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
	Verified bool
}

type Verification struct{
	Id int64
	UserId int64
	Token string
	ExpiresAt string
	CreatedAt string
}