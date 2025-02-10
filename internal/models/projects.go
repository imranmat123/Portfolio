package models

import "time"

type Projects struct {
	ProjectID   int64      `json:"project_id" db:"project_id"`
	UserID      int64      `json:"user_id" db:"user_id"`
	ProjectName *string    `json:"project_name" db:"project_name"`
	About       *string    `json:"about" db:"about"`
	CreatedAt   *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
	ProjectURL  *string    `json:"project_url" db:"project_url"`
}
