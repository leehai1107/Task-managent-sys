package entity

import (
	"github.com/leehai1107/Task-managent-sys/utils"
)

type Task struct {
	TaskID      uint           `json:"task_id gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	StartDate   utils.DateTime `json:"start_date"`
	EndDate     utils.DateTime `json:"end_date"`
	Status      string         `json:"status"`
}
