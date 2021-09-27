package dao

import (
	"hxj/initials"
	"hxj/internal/model"
)

var TestDaoInstance = TestDao{}

type TestDao struct {}

func (t *TestDao) GetById (id int) (r model.Test, err error) {
	r = model.Test{Id:id}
	_, err = initials.DB.Get(&r)
	if err != nil{
		return r, err
	}
	return r, nil
}
