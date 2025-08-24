package models

import "time"

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
	IsAdmin      bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
