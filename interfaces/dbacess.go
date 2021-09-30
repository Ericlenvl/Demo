package interfaces

import "hxj/internal/model"

type DBUser interface {
	GetById (id string) (r model.User, err error)
}


