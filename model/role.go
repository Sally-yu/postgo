package model

type Role struct {
	Id string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Admin bool `json:"admin"`
}