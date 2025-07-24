package dto

type LoginUserRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignupUserRequestDTO struct{
	Username string `json:"username" validate:"required,min=3,max=30,alphanum"`
	Email string `json:"email" vaildate:"required"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}