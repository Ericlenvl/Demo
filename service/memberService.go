package service

import (
	"hxj/internal/dao"
	"hxj/internal/model"
)

func GetMemberById(id int) (model.Test, error) {
	return dao.TestDaoInstance.GetById(id)
}
