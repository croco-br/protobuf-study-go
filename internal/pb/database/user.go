package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	db    *sql.DB
	ID    string
	Name  string
	Email *string
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (c *User) Create(name string, email string) (User, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO users (id, name, email) VALUES ($1, $2, $3)",
		id, name, email)
	if err != nil {
		return User{}, err
	}
	return User{ID: id, Name: name, Email: &email}, nil
}

func (c *User) FindAll() ([]User, error) {
	rows, err := c.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var id, name string
		var email sql.NullString
		if err := rows.Scan(&id, &name, &email); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		// Check if email is not null before accessing its value.
		var userEmail string
		if email.Valid {
			userEmail = email.String
		}

		users = append(users, User{ID: id, Name: name, Email: &userEmail})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return users, nil
}
