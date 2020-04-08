package model

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
