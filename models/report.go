package models

import "time"

// EmploymentReport represents an employment report with an ID and Content.
type EmploymentReport struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}
