package entity

import (
	"github.com/leehai1107/Task-managent-sys/constant"
	"github.com/leehai1107/Task-managent-sys/utils"
)

type Task struct {
	TaskID      uint                `json:"task_id" gorm:"primaryKey"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	StartDate   utils.DateTime      `json:"start_date"`
	EndDate     utils.DateTime      `json:"end_date"`
	Status      constant.TaskStatus `json:"status"`
	UserID      uint                `json:"user_id"`
}

func (Task) TableName() string { return "tasks" }
