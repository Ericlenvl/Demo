package initials

import (
	"fmt"
	"hxj/internal/model"
	"os"
	"xorm.io/xorm"
)

var DB *Orm

type Orm struct {
	*xorm.Engine
}

func Init() {
	database := AppCfg.Mysql
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" +
		database.Port + ")/" + database.DBName + "?charset=" + database.CharSet
	fmt.Println(conn)
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	orm := new(Orm)
	orm.Engine = engine
	DB = orm
}

func Migration() error {
	return DB.Sync2(new(model.Test))
}