package model

import "ark/db"

type Role struct {
	Id     int64  `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Enable int    `json:"enable"`
	Menu   string `json:"menu" xorm:"varchar"`
	Button string `json:"button" xorm:"varchar"`
	Other  string `json:"other" xorm:"varchar"`
}

func AllRole() (objs []Role, err error) {
	err = db.DBE.Find(&objs)
	return objs, err
}

func (obj *Role) Save() error {
	var err error
	if obj.Id > 0 {
		// id非空判断是否已存在
		existObj :=Role{Id:obj.Id}
		has,_:=existObj.Has()
		if has {
			// 存在记录 按id更新
			_, err = db.DBE.ID(existObj.Id).Update(obj)
		} else {
			_, err = db.DBE.Insert(obj)
		}
	} else {
		// 新增 直接保存
		_, err = db.DBE.Insert(obj)
	}
	_,_=obj.Has()
	return err
}

func (obj *Role) Delete() error {
	_, err := db.DBE.ID(obj.Id).Delete(obj)
	return err
}

func (obj *Role) Has() (has bool, err error) {
	has, err = db.DBE.Get(obj)
	return has, err
}

func (obj *Role) Find() (Objs []Role, err error) {
	err = db.DBE.Find(&Objs, obj)
	return Objs, err
}

func QueryRole(v string) (objs []Role, err error) {
	value := "%" + v + "%"
	err = db.DBE.Where("name like ? or code like ?", value, value).Find(&objs)
	return objs, err
}
