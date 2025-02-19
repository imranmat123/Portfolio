package models

type Education struct {
	EducationID int64      `json:"education_id" db:"education_id"`
	UserID      int64      `json:"user_id" db:"user_id"`
	Name        string     `json:"name" db:"name"`
	About       string     `json:"about" db:"about"`
	StartDate   CustomTime `json:"start_date" db:"start_date"`
	EndDate     CustomTime `json:"end_date" db:"end_date" `
}
