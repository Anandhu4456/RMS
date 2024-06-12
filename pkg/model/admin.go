package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name         string  `json:"name" gorm:"validate:required"`
	Email        string  `json:"email" gorm:"unique"`
	UserType     bool    `json:"user_type"`
	PasswordHash string  `json:"password_hash"`
	Profile      Profile `json:"profile"`
	JobsPosted   []Job   `json:"jobs_posted"`
	Users        []User  `json:"user"`
}

type AdminLogin struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type AdminToken struct{
	Admin AdminLogin
	Token 	string
}


type UserDetailsAtAdmin struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}