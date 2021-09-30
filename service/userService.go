package service

import (
	"hxj/interfaces"
	"sync"
	"xorm.io/xorm"
)

var (
	uOnce sync.Once
	u     *userService
)

type userService struct {
	db   interfaces.DBUser
	pool *xorm.Engine
}

func NewUserService() *userService {
	uOnce.Do(func() {
		u = &userService{
			db:  dbUser,
			pool:dbPool,
		}
	})
	return u
}

func (u *userService)GetUserInfo(id string) (userInfos []interfaces.UserInfo,err error) {
	rows,err := u.db.GetById(id)
	if err != nil {
		return nil, err
	}
	userInfos = append(userInfos, interfaces.UserInfo{
		UserId: rows.Id,
	})
	return
}
