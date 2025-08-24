package models

import (
	"testing"
)

func TestUser(*testing.T) *User {
	return &User{
		Username:     "someusername",
		Email:        "email",
		PasswordHash: "passwordhash",
		IsAdmin:      false,
	}
}
