package models

import "time"

// Job represents jobs for applicants
type Job struct {
	ID          int
	Title       string
	Description string
	Company     string
	Salary      string
	// CreatedAt is time at which this entity was created.
	CreatedAt time.Time

	// UpdatedAt is time at which this entity was last updated.
	UpdatedAt time.Time
}
