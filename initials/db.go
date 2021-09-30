package initials

import (
	"fmt"
	"hxj/internal/model"
	"os"
	"sync"
	"xorm.io/xorm"
)

var (
	dbOnce sync.Once
	dbPool *xorm.Engine = nil
)

func NewDBPool() *xorm.Engine {
	dbOnce.Do(func() {
		database := AppCfg.Mysql
		conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" +
			database.Port + ")/" + database.DBName + "?charset=" + database.CharSet
		fmt.Println(conn)
		dbPool, err := xorm.NewEngine("mysql", conn)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dbPool.SetMaxOpenConns(10)
	})
	return dbPool
}


func Migration() error {
	return dbPool.Sync2(new(model.User))
}