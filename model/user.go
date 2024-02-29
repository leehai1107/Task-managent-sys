package model

import "time"

type User struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	DoB      time.Time `json:"dob"`
	Sex      string    `json:"sex"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
	Tasks    []Task    `json:"tasks"`
}
