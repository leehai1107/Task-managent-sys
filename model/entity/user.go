package entity

import "github.com/leehai1107/Task-managent-sys/utils"

type User struct {
	UserID   uint       `json:"user_id" gorm:"primaryKey"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	Dob      utils.Date `json:"dob"`
	Sex      string     `json:"sex"`
	Address  string     `json:"address"`
	Phone    string     `json:"phone"`
}

func (User) TableName() string { return "users" }
