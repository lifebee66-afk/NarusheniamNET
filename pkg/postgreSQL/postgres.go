package postgresql

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func SQLinit() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(50) UNIQUE NOT NULL CHECK (LENGTH(login) >= 3),
    password VARCHAR(60) NOT NULL CHECK (LENGTH(password) >= 6),
    full_name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(15) NOT NULL CHECK (LENGTH(phone_number) >= 11),
    email VARCHAR(255) UNIQUE NOT NULL CHECK (email LIKE '%@%.%'),
);`
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=5286 dbname=usersdb sslmode=disable")
	if err != nil {
		return err
	}
	db.ExecContext(context.Background(), query)
	return err
}
