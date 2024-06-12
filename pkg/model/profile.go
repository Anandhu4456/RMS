package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID            uint   `json:"user_id"`
	ResumeFileAddress string `json:"resume_file_address"`
	Skills            string `json:"skills" gorm:"not null"`
	Education         string `json:"education" gorm:"not null"`
	Experience        string `json:"experience"`
	Name              string `json:"name" gorm:"validate:required"`
	Email             string `json:"email" gorm:"unique;not null"`
	Phone             string `json:"phone" gorm:"unique;not null"`
}
