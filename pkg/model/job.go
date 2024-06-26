package model

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	PostedOn          time.Time `json:"posted_on"`
	TotalApplications int       `json:"total_applicants"`
	CompanyName       string    `json:"company_name"`
	PostedByID        uint      `json:"posted_by_id"`
	PostedBy          Admin      `json:"posted_by"`
}
