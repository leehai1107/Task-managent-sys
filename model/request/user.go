package request

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
