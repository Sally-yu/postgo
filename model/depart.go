package model

import "ark/db"

type Department struct {
	Id        int64  `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Parent    int64  `json:"parent"`
	Companyid int64  `json:"companyid"`
}

type DepartQuery struct {
	Department `xorm:"extends"`
	Company    Company `json:"company" xorm:"extends"`
}

func (DepartQuery) TableName() string {
	return "department"
}

// 全部部门记录，关联公司
func AllDepart() (objs []DepartQuery, err error) {
	err = db.DBE.
		Join("LEFT", "company", "company.id = department.companyid").
		Find(&objs)
	return objs, err
}

func (obj *Department) Save() error {
	var err error
	if obj.Id > 0 {
		existObj := Department{Id: obj.Id}
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

func (obj *Department) Delete() error {
	_, err := db.DBE.ID(obj.Id).Delete(obj)
	return err
}

func (obj *Department) Has() (has bool, err error) {
	has, err = db.DBE.Get(obj)
	return has, err
}

// 查找一条记录，关联公司
func (obj *Department) Find() (Objs []DepartQuery, err error) {
	err = db.DBE.
		Join("LEFT", "company", "department.companyid=company.id").
		Find(&Objs, obj)
	return Objs, err
}

// 查找子记录，关联公司
func (obj *Department) FindChildren() (Objs []DepartQuery, err error) {
	err = db.DBE.Alias("d").Where("d.parent = ?", obj.Id).
		Join("LEFT", "company", "d.companyid=company.id").
		Find(&Objs)
	return Objs, err
}

// 模糊查询
func QueryDepart(v string) (Objs []DepartQuery, err error) {
	value := "%" + v + "%"
	err = db.DBE.Alias("d").Where(" d.name like ? or d.code like ?", value, value).
		Join("LEFT", "company", "d.companyid=company.id").
		Find(&Objs)
	return Objs, err
}
