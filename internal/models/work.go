package models

type Work struct {
	WorkID    int64      `json:"work_id" db:"work_id"`
	UserID    int64      `json:"user_id" db:"user_id"`
	Title     string     `json:"title" db:"title"`
	About     string     `json:"about" db:"about"`
	StartDate CustomTime `json:"start_date" db:"start_date"`
	EndDate   CustomTime `json:"end_date" db:"end_date"`
}
