package entity

import "time"

type Task struct {
	TaskID      uint      `json:"task_id gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      string    `json:"status"`
}
