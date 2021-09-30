package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hxj/initials"
	"hxj/router"
	"os"
	"os/signal"
	//"github.com/google/wire"
)

func init() {
	flag.StringVar(&initials.ConfPath,"path", "", "set path")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", "demo")
		flag.PrintDefaults()
	}
}

func main(){
	flag.Parse() //解析命令行参数
	initials.InitCfg()
	//initials.InitDb()
	err := initials.Migration()
	if err != nil {
		fmt.Println(err)
	}
	server := &router.Testobj{
		Rhandler: router.NewRestHandler(),
	}
	server.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
