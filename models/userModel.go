package models

import (
	"BackEnd_21BCE5685/db"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create a new user in the database
func (u *User) Create() error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	err := db.DB.QueryRow(query, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	return nil
}

// Find a user by email
func GetUserByEmail(email string) (*User, error) {
	user := &User{}
	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		log.Println("Error finding user:", err)
		return nil, err
	}
	return user, nil
}
