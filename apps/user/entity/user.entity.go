package entity

type AddUserRequestDTO struct {
	Email    string `validate:"required,email,min=4,max=64"`
	Password string `validate:"required,alphanum,min=8"`
	Username string `validate:"required,alpha,min=2"`
}
