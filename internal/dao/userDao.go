package dao

import (
	"hxj/initials"
	"hxj/internal/model"
	"sync"
	"xorm.io/xorm"
)

type user struct {
	db *xorm.Engine
}

var (
	tOnce sync.Once
	t     *user
)

func NewUser() *user {
	tOnce.Do(func() {
		t = &user{
			db: initials.NewDBPool(),
		}
	})
	return t
}

func (t *user) GetById (id string) (r model.User, err error) {
	r = model.User{Id:id}
	t.db.Get(&r)
	if err != nil{
		return r, err
	}
	return r, nil
}
