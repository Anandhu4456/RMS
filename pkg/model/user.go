package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name            string  `json:"name" gorm:"validate:required"`
	Email           string  `json:"email" gorm:"unique"`
	Address         string  `json:"address"`
	UserType        string  `json:"user_type"`
	PasswordHash    string  `json:"password_hash"`
	ProfileHeadline string  `json:"profile_headline"`
	Profile         Profile `json:"profile"`
	JobsPosted      []Job   `json:"jobs_posted"`
}

type UserLogin struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type UserSignup struct {
	Email               string `json:"email"`
	PasswordHash        string `json:"password_hash"`
	ConfirmPasswordHash string `json:"confirm_password_hash"`
}

type UserToken struct {
	User  UserSignup
	Token string
}
