package models

import (
	"time"
)

// User describes the internal structure of an user
type User struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	PasswordHash string    `json:"-"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// UserRegister describes the structure needed for register a new user
type UserRegister struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required" min:"8"`
	Name     string `json:"name"  binding:"required"`
}

// UserLogin describes the structure needed for a user to log in
type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"-" binding:"required"`
}

// UserResponse describes the user structure that can be returned (not pwd related info)
type UserResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (user *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
}
