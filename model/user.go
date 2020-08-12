package model

import (
	"ark/db"
)

type User struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Age        int    `json:"age"`
	Sex        int    `json:"sex"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Pwd        string `json:"pwd"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func AllUser() (users []User, err error) {
	err = db.DBE.Find(&users)
	return users, err
}

func (user *User) Save() error {
	var err error
	if user.Id != 0 {
		// id非空判断是否已存在
		var existUser User
		_, err = db.DBE.ID(user.Id).Get(&existUser)
		if existUser.Id != 0 {
			// 存在记录 按id更新
			_, err = db.DBE.ID(user.Id).Update(user)
		} else {
			_, err = db.DBE.Insert(user)
		}
	} else {
		// 新增 直接保存
		_, err = db.DBE.Insert(user)
	}
	return err
}

func (user *User) Delete() error {
	_, err := db.DBE.ID(user.Id).Delete(user)
	return err
}

func (user *User) One() error {
	_, err := db.DBE.ID(user.Id).Get(user)
	return err
}
