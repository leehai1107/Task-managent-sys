package request

import "github.com/leehai1107/Task-managent-sys/utils"

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdateReq struct {
	Username string     `json:"username"`
	Password string     `json:"password"`
	Dob      utils.Date `json:"dob"`
	Sex      string     `json:"sex"`
	Address  string     `json:"address"`
	Phone    string     `json:"phone"`
}

func (UserUpdateReq) TableName() string { return "users" }
