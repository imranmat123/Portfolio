package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	d, err := time.Parse("2006-01-02", strings.Trim(string(b), "\""))
	if err != nil {
		return fmt.Errorf("failed to parse date: %w", err)
	}
	t.Time = d
	return nil
}

func (t CustomTime) Value() (driver.Value, error) {
	return t.Time, nil
}
