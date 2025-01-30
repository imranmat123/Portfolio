package models

import "time"

type User struct {
	UserID         int64  `json:"user_id" db:"user_id"`
	FirstName      string `json:"first_name" db:"first_name"`
	LastName       string `json:"last_name" db:"last_name"`
	Email          string `json:"email" db:"email"`
	Password       string `json:"password" db:"password"`
	ProfilePicture string `json:"profile_picture_url" db:"profile_picture_url"`
	AboutMe        string `json:"about_me" db:"about_me"`
}

type GetUser struct {
	UserID         int64     `json:"user_id" db:"user_id"`
	FirstName      string    `json:"first_name" db:"first_name"`
	LastName       string    `json:"last_name" db:"last_name"`
	Email          string    `json:"email" db:"email"`
	Password       string    `json:"password" db:"password"`
	ProfilePicture string    `json:"profile_picture_url" db:"profile_picture_url"`
	AboutMe        string    `json:"about_me" db:"about_me"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	IsActive       bool      `json:"is_active" db:"is_active"`
}

//about_me TEXT,  -- User's description or bio
//created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Timestamp of account creation
//updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Timestamp of last update
//is_active BOOLEAN DEFAULT TRUE  -- Boolean for account status
