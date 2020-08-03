package model

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Age int `json:"age"`
	Sex int `json:"sex"`
	Phone int `json:"phone"`
	Email string `json:"email"`
	Pwd string `json:"pwd"`
	CreateTime int64 `json:"create_time"`
	LastLoginTime int64 `json:"last_login_time"`
}