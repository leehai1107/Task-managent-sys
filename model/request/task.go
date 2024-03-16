package request

import (
	"github.com/leehai1107/Task-managent-sys/constant"
	"github.com/leehai1107/Task-managent-sys/utils"
)

type Task struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	StartDate   utils.DateTime      `json:"start_date"`
	EndDate     utils.DateTime      `json:"end_date"`
	Status      constant.TaskStatus `json:"status"`
	UserID      uint                `json:"user_id"`
}
