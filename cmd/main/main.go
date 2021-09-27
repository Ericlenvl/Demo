package main

import (
	"flag"
	"fmt"
	"hxj/initials"
	"hxj/router"
	"os"
	"os/signal"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/google/wire"
)

var name string
var age int

func init() {
	flag.StringVar(&name,"name", "t", "set name")
	flag.StringVar(&initials.ConfPath,"path", "", "set path")
	flag.IntVar(&age,"age", 1, "set name")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", "demo")
		flag.PrintDefaults()
	}
}

func main(){
	flag.Parse() //解析命令行参数
	initials.InitCfg()
	initials.Init()
	err := initials.Migration()
	if err != nil {
		fmt.Println(err)
	}
	go router.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
