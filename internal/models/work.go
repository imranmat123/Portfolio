package models

import "time"

type Work struct {
	WorkID    int64     `json:"work_id" db:"work_id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	About     string    `json:"about" db:"about"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
}
