package model

import "ark/db"

type Company struct {
	Id     int64  `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Parent int64  `json:"parent"`
}

func AllCompany() (objs []Company, err error) {
	err = db.DBE.Find(&objs)
	return objs, err
}

func (obj *Company) Save() error {
	var err error
	if obj.Id > 0 {
		existObj := Company{Id: obj.Id}
		has, _ := existObj.Has()
		if has {
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

func (obj *Company) Delete() error {
	_, err := db.DBE.ID(obj.Id).Delete(obj)
	return err
}

func (obj *Company) Has() (has bool, err error) {
	has, err = db.DBE.Get(obj)
	return has, err
}

func (obj *Company) Find() (Objs []Company, err error) {
	err = db.DBE.Find(&Objs, obj)
	return Objs, err
}

func (obj *Company) FindChildren() (objs []Company, err error) {
	err = db.DBE.Where("parent = ?", obj.Id).Find(&objs)
	return objs, err
}

func QueryCompany(v string) (objs []Company, err error) {
	value := "%" + v + "%"
	err = db.DBE.Where("name like ? or code like ?", value, value).Find(&objs)
	return objs, err
}
