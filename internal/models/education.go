package models

import "time"

type Education struct {
	EducationID int64     `json:"education_id" db:"education_id"`
	UserID      int64     `json:"user_id" id:"user_id"`
	Name        string    `json:"name" db:"name"`
	About       string    `json:"about" db:"about"`
	StartDate   time.Time `json:"start_date" db:"start_date"`
	EndDate     time.Time `json:"end_date" db:"end_date"`
}
