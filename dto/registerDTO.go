package dto

import _ "github.com/go-playground/validator/v10"

type RegisterUserDTO struct {
	Login        string `json:"login" validate:"required,min=3,max=50,alphanum"`
	Password     string `json:"password" validate:"required,min=6,max=60"`
	Full_name    string `json:"full_name" validate:"required,min=2,max=100"`
	Phone_number string `json:"phone_number" validate:"required,min=11,max=15,numeric"`
	Email        string `json:"email" validate:"required,email,max=255"`
}
