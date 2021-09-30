package service

import (
	"hxj/interfaces"
	"xorm.io/xorm"
)

var (
	dbUser interfaces.DBUser
	dbPool *xorm.Engine
)

func SetDBPool(i *xorm.Engine) {
	dbPool = i
}