package model

import (
	"ark/db"
	"time"
)

type User struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Code         string    `json:"code"`
	Age          int       `json:"age"`
	Sex          int       `json:"sex"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Pwd          string    `json:"pwd"`
	Enable       int       `json:"enable"`
	Level        int       `json:"level"`
	Status       int       `json:"statue"`
	Entertime    time.Time `json:"entertime"`
	Departmentid int64     `json:"departmentid"`
	Roleid       int64     `json:"roleid"`
}

type UserQuery struct {
	User       `xorm:"extends"`
	Department DepartQuery `json:"department" xorm:"extends"`
	Role       Role        `json:"role" xorm:"extends"`
}

func (UserQuery) TableName() string {
	return "user"
}

func AllUser() (users []UserQuery, err error) {
	err = db.DBE.Alias("u").
		Join("LEFT", "department", "u.departmentid = department.id").
		Join("LEFT", "company", "department.companyid=company.id").
		Join("LEFT", "role", "u.roleid=role.id").
		Find(&users)
	return users, err
}

func (user *User) Has() (has bool, Obj UserQuery, err error) {
	Obj.Pwd=user.Pwd
	Obj.Code=user.Code
	has, err = db.DBE.Alias("u").
		Join("LEFT", "department", "u.departmentid = department.id").
		Join("LEFT", "company", "department.companyid=company.id").
		Join("LEFT", "role", "u.roleid=role.id").
		Get(&Obj)
	return has, Obj, err
}

func (user *User) Save() error {
	var err error
	if user.Id > 0 {
		// id非空判断是否已存在
		existUser := User{Id: user.Id}
		has, _, _ := existUser.Has()
		if has {
			// 存在记录 按id更新
			_, err = db.DBE.ID(existUser.Id).Update(user)
		} else {
			_, err = db.DBE.Insert(user)
		}
	} else {
		// 新增 直接保存
		_, err = db.DBE.Insert(user)
	}
	_, _, _ = user.Has()
	return err
}

func (user *User) Delete() error {
	_, err := db.DBE.ID(user.Id).Delete(user)
	return err
}

// 根据传输的user某一特征查找，比如id code name 结果为数组
func (user *User) Find() (Objs []UserQuery, err error) {
	err = db.DBE.Alias("u").
		Join("LEFT", "department", "u.departmentid = department.id ").
		Join("LEFT", "company", "department.companyid=company.id").
		Join("LEFT", "role", "u.roleid=role.id").
		Find(&Objs, user)
	return Objs, err
}

// 模糊查询，code name email
func QueryUser(v string) (Objs []UserQuery, err error) {
	value := "%" + v + "%"
	err = db.DBE.Alias("u").Where("u.name like ? or u.code like ? or u.email like ?", value, value, value).
		Join("LEFT", "department", "u.departmentid=department.id").
		Join("LEFT", "company", "department.companyid=company.id").
		Join("LEFT", "role", "u.roleid=role.id").
		Find(&Objs)
	return Objs, err
}
