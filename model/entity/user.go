package entity

type User struct {
	UserID   uint   `json:"user_id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dob      string `json:"dob"`
	Sex      string `json:"sex"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}
