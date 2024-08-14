package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `joson:"created_at"`
}

type Register struct {
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password 	string    `json:"password"`
}