package utils

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Date time.Time

func (d *Date) Scan(value interface{}) error {
	if value == nil {
		*d = Date(time.Time{})
		return nil
	}
	*d = Date(value.(time.Time))
	return nil
}

func (d Date) Value() (driver.Value, error) {
	return time.Time(d), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format("2006-01-02"))
}
