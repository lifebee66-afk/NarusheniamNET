package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lifebee66-afk/NarusheniamNET/dto"
)

type RegisterUser struct {
	DB *sql.DB
}

func NewRegisterUser(DB *sql.DB) *RegisterUser {
	return &RegisterUser{DB: DB}
}

func (r *RegisterUser) RegisterUser(ctx context.Context, user *dto.RegisterUserDTO) error {
	query := `INSERT INTO users (login, password, full_name, phone_number, email) VALUES ($1, $2, $3, $4, $5)`
	res, err := r.DB.ExecContext(ctx, query, &user.Login, &user.Password, &user.Full_name, &user.Phone_number, &user.Email)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("failed register users, try again")
	}
	return nil
}
