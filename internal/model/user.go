package model

type User struct {
	Id    string  `xorm:"pk autoincr" json:"id"`
}
