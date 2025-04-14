package models

import "time"

type User struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	FullName   string    `json:"full_name"`
	Bio        string    `json:"bio"`
	AvatarURL  string    `json:"avatar_url"`
	Website    string    `json:"website"`
	IsVerified bool      `json:"is_verified"`
	IsPrivate  bool      `json:"is_private"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
