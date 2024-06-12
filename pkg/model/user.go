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
