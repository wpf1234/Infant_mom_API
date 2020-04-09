package model

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisterReq struct {
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
